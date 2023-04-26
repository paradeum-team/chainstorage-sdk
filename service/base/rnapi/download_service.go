package rnapi

import (
	"bytes"
	"chainstorage-sdk/common/dict"
	"chainstorage-sdk/common/e/code"
	"chainstorage-sdk/common/plogger"
	"chainstorage-sdk/conf"
	"chainstorage-sdk/entity"
	"chainstorage-sdk/service/base/afs"
	"chainstorage-sdk/utils"
	"chainstorage-sdk/vo"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty"
	"github.com/juju/ratelimit"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
*
小文件下载，不落盘
*/
func (*rnApiService) RnSimpleDownload(dgstResp *entity.RespQueryDgst) entity.ResponseVO {
	var responseVO = entity.ResponseVO{}
	sprintf := fmt.Sprintf(conf.RnDownloadUrl, dgstResp.Afid)
	url := strings.Join([]string{dgstResp.RnodeRoots[0], sprintf}, "")
	utils.DevLog("RnSimpleDownload", fmt.Sprintf("当前下载使用RFS = %v \n", dgstResp.NodeType[dict.NODE_TYPE_RFS]))
	defer utils.DevTimeTrack(time.Now(), "RnSimpleDownload")
	resp, err := retryDownload(url, nil)
	if err != nil {
		responseVO.HttpCode = http.StatusInternalServerError
		responseVO.Code = code.ScodeRnNetConnectionError
		responseVO.Msg = err.Error()
		return responseVO
	}
	responseVO = handleDownloadExecption(url, responseVO, resp)
	if responseVO.HttpCode > http.StatusPartialContent && responseVO.Code != 0 {
		return responseVO
	}
	responseVO.Data = resp.Body()
	return responseVO
}

/*
*
根据Combine下载文件
*/
func (*rnApiService) RnCombineFileDownload(paramVO *vo.DownloadParamsVO, dgstResp *entity.RespQueryDgst) entity.ResponseVO {
	downLoadLimitRate := conf.Config.DownloadLimitRate
	bucket := ratelimit.NewBucketWithRate(downLoadLimitRate*dict.BASE_FILE_SIZE, int64(downLoadLimitRate*dict.BASE_FILE_SIZE))
	bucket.Wait(int64(downLoadLimitRate * dict.BASE_FILE_SIZE))

	c := paramVO.GinContext

	fileName := paramVO.FileName
	isBigfile := paramVO.IsBigfile
	isShortApi := paramVO.IsShortApi
	isSeedfile := paramVO.IsSeedFile
	var responseVO = entity.ResponseVO{http.StatusInternalServerError, code.ERROR, "", nil}
	var responseErrorVO = entity.ResponseVO{http.StatusInternalServerError, code.ERROR, "", nil}

	sprintf := fmt.Sprintf(conf.RnDownloadUrl, dgstResp.Afid)
	url := strings.Join([]string{dgstResp.RnodeRoots[0], sprintf}, "")
	utils.DevLog("RnCombineFileDownload", fmt.Sprintf("当前下载使用RFS = %v \n", dgstResp.NodeType[dict.NODE_TYPE_RFS]))

	defer utils.DevTimeTrack(time.Now(), "downloadFile")

	//分片三次重试
	resp, err := retryDownload(url, paramVO)
	if err != nil {
		responseVO.Code = code.ScodeRnNetConnectionError
		responseVO.Data = err
		return responseVO
	}
	responseVO = handleDownloadExecption(url, responseVO, resp)
	if responseVO.HttpCode > http.StatusPartialContent && responseVO.Code != 0 {
		return responseVO
	} else {
		if isSeedfile == true {
			filepath := filepath.Join(conf.DownloadDir, dgstResp.Afid+".dat")
			out, err := os.Create(filepath)
			if err != nil {
				responseErrorVO.Data = err
				return responseErrorVO
			}
			defer out.Close()

			start := time.Now()
			//n, err := out.Write(resp.Body())
			copySize, err := io.Copy(out, ratelimit.Reader(bytes.NewReader(resp.Body()), bucket))
			utils.DevLog("RnCombineFileDownload", fmt.Sprintf("Download seedfile Copied %d bytes in %s", copySize, time.Since(start)))
			if err != nil {
				responseErrorVO.Data = err
				return responseErrorVO
			}
			return responseVO
		}
		if isShortApi == true {
			setUpHeader(fileName, resp, c)
		}
		if isBigfile != true {
			c = setResponseHeader(resp, c)
			c.Writer.WriteHeader(resp.RawResponse.StatusCode)
			start := time.Now()
			copySzie, err := io.Copy(c.Writer, ratelimit.Reader(bytes.NewReader(resp.Body()), bucket))
			utils.DevLog("RnCombineFileDownload", fmt.Sprintf("Download  Copied %d bytes in %s", copySzie, time.Since(start)))
			if err != nil {
				responseErrorVO.Data = err
				return responseErrorVO
			}
			return responseVO
			//c.Writer.Write(resp.Body())
		} else {
			start := time.Now()
			copySzie, err := io.Copy(c.Writer, ratelimit.Reader(bytes.NewReader(resp.Body()), bucket))
			utils.DevLog("RnCombineFileDownload", fmt.Sprintf("Download  Copied %d bytes in %s", copySzie, time.Since(start)))
			if err != nil {
				responseErrorVO.Data = err
				return responseErrorVO
			}
			return responseVO
			//c.Writer.Write(resp.Body())
		}
	}
	return responseVO
}

/*
*
重试下载
*/
func retryDownload(url string, paramVO *vo.DownloadParamsVO) (*resty.Response, error) {
	c := &gin.Context{}
	resp := &resty.Response{}
	var err error

	restyClietnt := resty.SetTimeout(time.Duration(conf.GatewayTimeout) * time.Second)
	if paramVO != nil {
		c = paramVO.GinContext
		if paramVO.IsSeedFile == false {
			restyClietnt.Header = c.Request.Header
		}
	}
	utils.DevLog("retryDownload", fmt.Sprintf("文件当前Range = %s", restyClietnt.Header.Get("Range")))
	for i := 1; i <= 3; i++ {
		resp, err = restyClietnt.R().Get(url)
		//utils.DevLog("retryDownload",fmt.Sprintf("第[%v]次请求RN Api下载数据,httpcode = %v",i,resp.StatusCode()))
		if err != nil || resp.StatusCode() > http.StatusPartialContent {
			plogger.NewInstance().GetLogger().Errorf("第[%v]次----------请求RN Api下载数据,httpcode = %v,err: %v", i, resp.StatusCode(), err)
			time.Sleep(time.Duration(1) * time.Second)
			continue
		} else {
			break
		}
	}
	if err == nil {
		defer resp.RawResponse.Body.Close()
	}
	restyClietnt.Header = nil
	return resp, err
}

/*
*
获取设置Response header
*/
func setResponseHeader(resp *resty.Response, c *gin.Context) *gin.Context {
	headers := resp.RawResponse.Header
	c.Writer.Header().Add("Accept-Ranges", headers.Get("Accept-Ranges"))
	c.Writer.Header().Add("Content-Length", headers.Get("Content-Length"))
	c.Writer.Header().Add("Date", headers.Get("Date"))
	c.Writer.Header().Add("Last-Modified", headers.Get("Last-Modified"))
	c.Writer.Header().Add("Content-Range", headers.Get("Content-Range"))
	return c
}

/*
*
RN下载文件异常处理
*/
func handleDownloadExecption(url string, responseVO entity.ResponseVO, resp *resty.Response) entity.ResponseVO {
	if resp.StatusCode() == http.StatusRequestTimeout {
		responseVO.HttpCode = resp.StatusCode()
		responseVO.Code = code.ScodeAFSTimeout
		responseVO.Msg = errors.New("downloadFile:rnode download timeout ").Error()
		utils.LogError("downloadFile:rnode download timeout !")
		return responseVO
	}

	if resp.StatusCode() == http.StatusRequestedRangeNotSatisfiable {
		responseVO.HttpCode = resp.StatusCode()
		responseVO.Code = code.ScodeDownloadRangeErr
		responseVO.Msg = string(resp.Body())
		utils.LogError(fmt.Sprintf("handleDownloadExecption RN下载文件异常处理：%s", string(resp.Body())))
		return responseVO
	}

	if resp.StatusCode() > http.StatusPartialContent {
		plogger.NewInstance().GetLogger().Errorf("handleDownloadExecption RN下载文件异常处理 :%s", string(resp.Body()))
		var respBody afs.WebServiceGenericResp
		err := json.Unmarshal(resp.Body(), &respBody)
		if err != nil {
			responseVO.HttpCode = resp.StatusCode()
			responseVO.Code = code.ERROR
			responseVO.Msg = err.Error()
			return responseVO
		}
		respDataMap, ok := respBody.Data.(map[string]string)
		var note = ""
		if ok {
			for k, v := range respDataMap {
				if k == "note" {
					note = v
					break
				}
			}
			responseVO.Msg = fmt.Errorf("download from %s status:%d note:%s", url, resp.StatusCode(), note).Error()
			utils.LogError(responseVO.Msg)
		} else {
			utils.LogError("解析RN返回值失败,Failed to parse RN return value")
			responseVO.Msg = fmt.Errorf("download from %s status:%d note:%s", url, resp.StatusCode(), "Resolution rn respBody failure ").Error()
		}
		responseVO.HttpCode = resp.StatusCode()
		responseVO.Code = code.ScodeRNodesInfoErr
		return responseVO
	}

	if strings.HasPrefix(resp.Header().Get("Content-Type"), "application/json") {
		// Probably getting an error message in response body, instead of file data.
		// Parse RNode resposne.
		var respBody afs.WebServiceGenericResp
		if respBody.Code != code.ScodeOK {
			respDataMap, ok := respBody.Data.(map[string]string)
			raw := ""
			if ok {
				for k, v := range respDataMap {
					if k == "raw" {
						raw = v
						break
					}
				}
				responseVO.HttpCode = http.StatusInternalServerError
				responseVO.Code = code.ScodeRnErr
				responseVO.Msg = fmt.Errorf("download from %s data.raw:%s msg:%s", url, raw, respBody.Msg).Error()
				utils.LogError(responseVO.Msg)
				return responseVO
			}
			responseVO.HttpCode = http.StatusInternalServerError
			responseVO.Code = code.ScodeRnErr
			responseVO.Msg = fmt.Errorf("download from %s msg:%s", url, respBody.Msg).Error()
			utils.LogError(responseVO.Msg)
			return responseVO
		}
	}
	responseVO.Code = 0
	responseVO.HttpCode = resp.StatusCode()
	return responseVO
}

func setUpHeader(filename string, resp *resty.Response, c *gin.Context) {
	// Only the first 512 bytes are used to sniff the content type.
	buffer := resp.Body()[:512]
	contentType := http.DetectContentType(buffer)
	if strings.HasPrefix(contentType, "image/") || strings.HasPrefix(contentType, "video/") || strings.HasPrefix(contentType, "audio/") {
		c.Header("Content-Type", contentType)
		return
	}
	// Force the content-type to specific values if requested filename has the extension .m3u8, .ts or .dat.
	// Or try to match content-type based on extension of requested filename.
	if len(filename) >= 6 && filename[len(filename)-5:] == ".m3u8" {
		c.Header("Content-Type", "application/vnd.apple.mpegurl")
		return
	} else if len(filename) >= 4 && filename[len(filename)-3:] == ".ts" {
		c.Header("Content-Type", "video/mp2t")
		return
	} else if len(filename) > 0 {
		if strings.HasSuffix(filename, ".dat") {
			//utils.DevLog2(devCounter, "serveFile", fmt.Sprintf("detected content-type %s setting to application/octet-stream on .dat", contentType))
			c.Header("Content-Type", "application/octet-stream")
			return
		} else if strings.HasSuffix(filename, ".js") {
			c.Header("Content-Type", "text/javascript")
			return
		} else if strings.HasSuffix(filename, ".css") {
			c.Header("Content-Type", "text/css")
			return
		} else {
			if contentType == "application/octet-stream" {
				if strings.Contains(filename, ".") {
					ctByExt := mime.TypeByExtension(filename[strings.LastIndex(filename, "."):])
					if ctByExt != "" {
						contentType = ctByExt
					}
				}
			}
			c.Header("Content-Type", contentType)
			return
		}
	}

}
