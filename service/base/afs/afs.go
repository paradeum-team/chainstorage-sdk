package afs

import (
	"chainstorage-sdk/common/dict"
	"chainstorage-sdk/common/e/code"
	"chainstorage-sdk/conf"
	"chainstorage-sdk/conf/dgst"
	"chainstorage-sdk/conf/uri"
	"chainstorage-sdk/entity"
	"chainstorage-sdk/service/base/tnapi"
	"chainstorage-sdk/utils"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

type PostFilev1JsonResp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type PostFileJson struct {
	Afid string `json:"afid"`
	Raw  string `json:"raw"`
}

type WebServiceGenericResp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const LARGE_FILE_IDENTIFICATION = "largefile"

/*
*
获取健康TN
*/
func GetActiveTrackRoot(allTrackRoots []string) string {
	//defer utils.DevTimeTrack(time.Now(), "getActiveTrackRoot")
	randIdx := rand.Intn(len(allTrackRoots))
	response := tnapi.NewTnApiServiceInstace().CheckSysHealthy(allTrackRoots[randIdx])
	if response.HttpCode == http.StatusOK {
		return allTrackRoots[randIdx]
	}
	for idx, root := range allTrackRoots {
		if idx != randIdx {
			response := tnapi.NewTnApiServiceInstace().CheckSysHealthy(allTrackRoots[idx])
			if response.HttpCode == http.StatusOK {
				return root
			}
		}
	}
	return ""
}

/*
*
随机选取tnode
*/
func GetRandomRNodeRoot(trackRoot string) (response entity.ResponseVO) {
	//完全可用的RN(可读可写)
	infos := []entity.RnInfo{}
	defer utils.DevTimeTrack(time.Now(), "getRandomRNodeRoot")
	//rnodeInfos, isTrackTimeout, err := GetRNodesInfo(trackRoot)
	rnInfosResponse := tnapi.NewTnApiServiceInstace().GetRnCluster(trackRoot)
	response.HttpCode = rnInfosResponse.HttpCode
	response.Code = rnInfosResponse.Code
	response.Msg = rnInfosResponse.Msg
	rnodeInfos := rnInfosResponse.Data

	if rnodeInfos != nil {
		for _, rnodeInfo := range rnodeInfos {
			if rnodeInfo.RNHealthy.Status == "healthy" {
				if rnodeInfo.RNHealthy.Available.Writeable == true && rnodeInfo.RNHealthy.Available.Readable == true {
					infos = append(infos, rnodeInfo)
				}
			}
		}
		if len(infos) > 0 {
			rnUrl := infos[rand.Intn(len(infos))].Api
			response.Data = rnUrl
			return response
		}
	}
	return response
}

/*
*
TN 根据dgst寻址，包括大文件寻址，返回一个随机地址
*/
func QueryDgstLocation(dgstString ...string) (httpcode int, dgstResp *entity.RespQueryDgst) {
	response := entity.ResponseVO{}
	defer utils.DevTimeTrack(time.Now(), "QueryDgstLocation")
	dgstType := dgst.GetDataTypeByDgst(dgstString[0])
	trackRoot := GetActiveTrackRoot(conf.Config.TrackRoots)

	sprintf := fmt.Sprintf(uri.TN_DEL_CACHE_BYDGST_URL, dgstString[0])
	url := strings.Join([]string{trackRoot, sprintf}, "")
	resty.SetTimeout(time.Duration(conf.GatewayTimeout) * time.Second).R().Delete(url)
	if len(dgstString) > 1 {
		if dgstString[1] == LARGE_FILE_IDENTIFICATION {
			response = tnapi.NewTnApiServiceInstace().TnQueryBigfileLocationByDgst(trackRoot, dgstString[0])
		}
	} else {
		useField := conf.Config.UseField
		response = tnapi.NewTnApiServiceInstace().TnQueryLocationByDgst(trackRoot, dgstString[0])

		if dgstType == dgst.AfsAlgoTypeGfid {
			useField = dict.NODE_TYPE_GFS
		}
		httpcode, dgstResp = handleQueryDgstResponse(dgstType, dgstString[0], useField, response)
		if dgstResp.NodeType[dict.NODE_TYPE_GFS] == true {
			response = tnapi.NewTnApiServiceInstace().TnQueryLocationByDgst(trackRoot, dgstResp.Afid)
			return httpcode, dgstResp
		} else {
			return httpcode, dgstResp
		}
	}

	return handleQueryDgstResponse(dgstType, dgstString[0], conf.Config.UseField, response)
}

/*
*
TN 根据dgst寻址返回一组地址
*/
func QueryDgstLocationByNodeType(dgstString string, nodeType string) (httpcode int, dgstResp *entity.RespQueryDgst) {
	response := entity.ResponseVO{}
	defer utils.DevTimeTrack(time.Now(), "QueryDgstLocationByNodeType")
	dgstType := dgst.GetDataTypeByDgst(dgstString)
	trackRoot := GetActiveTrackRoot(conf.Config.TrackRoots)
	response = tnapi.NewTnApiServiceInstace().TnFileQueryLocationDgst2Afid(trackRoot, dgstType, dgstString)
	return handleQueryDgstResponse(dgstType, dgstString, nodeType, response)
}

// 查询gifd 所在rnode
func GetRNodeGfid(gfid, trackRoot string) (rnodeRoots, rnids []string, sysCode int, msg string, isTimeout bool, err error) {
	defer utils.DevTimeTrack(time.Now(), "GetRNodeGfid")
	sysCode = code.ScodeOK
	isTimeout = false
	utils.DevLog("getRNodeRoots", fmt.Sprintf("querying track node at %s", trackRoot+uri.TrackLocationPathPrefixGFID+gfid))
	resp, err := resty.R().Get(trackRoot + uri.TrackLocationPathPrefixGFID + gfid)
	if err != nil {
		if perr, ok := err.(net.Error); ok && perr.Timeout() {
			utils.DevLog("getRNodeRoots", fmt.Sprintf("query track node at %s timed out", trackRoot+gfid))
			sysCode = code.ScodeTrackingErr
			isTimeout = true
			return
		}
		sysCode = code.ScodeTrackingErr
		utils.LogError(fmt.Sprintf("query track node locating afid:%s failed:%v", gfid, err))
		return
	}
	// Parse TNode response.
	var respBody WebServiceGenericResp
	err = json.Unmarshal(resp.Body(), &respBody)
	utils.DevLog("GetRNodeGfid", fmt.Sprintf("response body: %s", string(resp.Body())))
	if err != nil {
		sysCode = code.ScodeTrackingErr
		utils.LogError(fmt.Sprintf("cannot parse response from track node locating afid:%s", gfid))
		return
	}

	utils.DevLog("getRNodeRoots", fmt.Sprintf("track response msg:%s", respBody.Msg))
	sysCode = respBody.Code
	msg = respBody.Msg
	if resp.StatusCode() != 200 {
		utils.DevLog("getRNodeRoots", fmt.Sprintf("track failed for %s:%s", gfid, string(resp.Body())))
		if resp.StatusCode() == http.StatusRequestTimeout {
			isTimeout = true
			return
		}

		return
	}

	if respBody.Data == nil {
		utils.LogError(fmt.Sprintf("nil response from track node locating afid:%s", gfid))
		return
	}

	sData, ok := respBody.Data.(map[string]interface{})

	var rNodes []interface{}
	if !ok {
		utils.LogError(fmt.Sprintf("cannot parse response from track node locating afid:%s", gfid))
		return
	}

	for key, value := range sData {
		if key == "rnodes" {
			rNodes = value.([]interface{})
		}
	}

	if respBody.Code == code.ScodeOK {
		for _, rNode := range rNodes {
			rNodeMap := rNode.(map[string]interface{})
			if rnid, ok := rNodeMap["rnid"]; ok {
				rnids = append(rnids, rnid.(string))
			}
		}
		utils.DevLog("getRNodeRoots", fmt.Sprintf("successfully got %d len rnids", len(rnids)))
	}

	if len(rnids) == 0 {
		utils.DevLog("getRNodeRoots", fmt.Sprintf("query track node at %s timed out", trackRoot+gfid))
		sysCode = code.ScodeTrackingErr
		isTimeout = true
		return
	}

	//返回rid，【9236】，然后从TN映射相应的rid-ip
	if len(rnids) > 0 {
		var rnodeInfos map[string]entity.RnInfo
		//从TN获取现在的rnodeInfo
		rnInfoResponse := tnapi.NewTnApiServiceInstace().GetRnCluster(trackRoot)
		//rnodeInfos, isTimeout, err = GetRNodesInfo(trackRoot)
		sysCode = rnInfoResponse.Code
		rnodeInfos = rnInfoResponse.Data

		for _, rnid := range rnids {
			rnodeInfo, ok := rnodeInfos[rnid]
			if ok {
				thisAPI := rnodeInfo.Api
				utils.DevLog("GetRNodeGfid", fmt.Sprintf("rid->ip:%s", thisAPI))
				if thisAPI != "" && !strings.HasSuffix(thisAPI, "/") {
					thisAPI = rnodeInfo.Api + "/"
				}

				if strings.HasPrefix(thisAPI, "http") {
					rnodeRoots = append(rnodeRoots, thisAPI)
				}
			}
		}
	}

	sysCode = respBody.Code
	msg = respBody.Msg
	if respBody.Code == code.ScodeNotExist {
		utils.DevLog("getRNodeRoots", fmt.Sprintf("track failed expired/deleted file %d and return %d rnRoots", respBody.Code, len(rnodeRoots)))
		return
	} else if respBody.Code == code.ScodeNotExist {
		utils.DevLog("getRNodeRoots", fmt.Sprintf("track failed missing file %d and return %d rnRoots", respBody.Code, len(rnodeRoots)))
		return
	}
	utils.DevLog("getRNodeRoots", fmt.Sprintf("track code %d and return %d rnRoots", respBody.Code, len(rnodeRoots)))
	return
}
