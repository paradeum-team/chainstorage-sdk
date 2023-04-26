package rnapi

import (
	"bytes"
	"chainstorage-sdk/common/app"
	"chainstorage-sdk/common/e/code"
	"chainstorage-sdk/common/plogger"
	"chainstorage-sdk/conf"
	"chainstorage-sdk/service/base/afs"
	"chainstorage-sdk/utils"
	"chainstorage-sdk/vo"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
*
上传文件并且删除源文件
*/
func (*rnApiService) RnCombineFileUpload(paramsVO *vo.UploadParamsVO) {
	defer os.RemoveAll(paramsVO.SourceFilePath)
	c := paramsVO.GinContext
	uploadPath := conf.RnUploadUrl

	// add by dxc: add post file to single bfs for test.
	rnodeTargetURL := strings.Join([]string{paramsVO.RnodeRoot, uploadPath}, "")
	fileb, err := ioutil.ReadFile(filepath.Join(conf.UploadDir, paramsVO.LocalFileName))
	if err != nil {
		sprintf := fmt.Sprintf("reading uploaded file failed:%v", err)
		utils.LogError(sprintf)
		app.NewResponse(c, http.StatusInternalServerError, code.ScodeSysErr, sprintf)
		return
	}
	resp, err := resty.SetTimeout(time.Duration(conf.GatewayTimeout)*time.Second).R().SetFileReader("file", paramsVO.SourceFileName, bytes.NewReader(fileb)).
		SetFormData(map[string]string{
			"days": paramsVO.Days, "asympubkey": paramsVO.Asympubkey, "sympassen": paramsVO.Sympassen, "encryptedType": paramsVO.EncryptedType}).
		Post(rnodeTargetURL)
	if err != nil {
		if perr, ok := err.(net.Error); ok && perr.Timeout() {
			sprintf := fmt.Sprintf("POST to %s new upload timed out", rnodeTargetURL)
			plogger.NewInstance().GetLogger().Errorf("POST fileUpload %s \n", sprintf)
			app.NewResponse(c, http.StatusGatewayTimeout, code.ScodeTnNetConnectionError, sprintf)
			return
		}
		sprintf := fmt.Sprintf("POST to %s new upload failed:%v", rnodeTargetURL, err)
		plogger.NewInstance().GetLogger().Error(sprintf)
		app.NewResponse(c, http.StatusInternalServerError, code.ScodeRnErr, sprintf)
		return
	}
	utils.DevLog("POST fileUpload", fmt.Sprintf("submitting file to RN %s received %s", rnodeTargetURL, resp.Body()))
	var respBody afs.PostFilev1JsonResp
	err = json.Unmarshal(resp.Body(), &respBody)
	if err != nil {
		sprintf := fmt.Sprintf("submitting file to RN %s received %s ,error is %s", rnodeTargetURL, resp.Body(), err.Error())
		plogger.NewInstance().GetLogger().Error(sprintf)
		app.NewResponse(c, http.StatusInternalServerError, code.ScodeSysErr, sprintf)
		return
	}
	if respBody.Code != 0 {
		note := respBody.Data.(map[string]interface{})["note"].(string)
		raw := respBody.Data.(map[string]interface{})["raw"].(string)
		sprintf := fmt.Sprintf("POST fileUpload error msg %s %s", raw, note)
		plogger.NewInstance().GetLogger().Error(sprintf)
		app.NewResponse(c, http.StatusInternalServerError, code.ScodeRnErr, &map[string]interface{}{"isExist": false, "msg": sprintf})
		return
	}
	// 5. 删除源文件，只留下afid计算后的文件
	//os.RemoveAll(paramsVO.SourceFilePath)
	app.NewResponse(c, http.StatusOK, code.ScodeOK, respBody.Data.(map[string]interface{}))
	return
}
