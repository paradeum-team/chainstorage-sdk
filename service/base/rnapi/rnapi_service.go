package rnapi

import (
	"chainstorage-sdk/base/client"
	"chainstorage-sdk/common/e/code"
	"chainstorage-sdk/entity"
	"chainstorage-sdk/service/base/afs"
	"chainstorage-sdk/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
*
获取Params
*/
func (*rnApiService) RnReadParams(url string) *entity.ResponseVO {
	httpStatus, body, err := client.RestyGet(url)
	return handleRnReadParams(url, httpStatus, body, err)

}

func handleRnReadParams(url string, httpStatus int, body []byte, err error) *entity.ResponseVO {
	respBody := afs.WebServiceGenericResp{}
	response := &entity.ResponseVO{}
	if err != nil {
		response.HttpCode = httpStatus
		response.Data = err.Error()
		response.Code = code.ScodeRnNetConnectionError
		return response
	}
	err = json.Unmarshal(body, &respBody)
	if err == nil {
		if respBody.Data != nil && respBody.Code == code.ScodeOK {
			respBodyDataMap, ok := respBody.Data.(map[string]interface{})
			if ok {
				value := respBodyDataMap["parameter_value"].(string)
				response.Data = value
				response.HttpCode = http.StatusOK
				response.Code = respBody.Code
				return response
			}
		}
		utils.LogError(fmt.Sprintf("error response data reading param from %s", url))
		respBodyDataMap, ok := respBody.Data.(map[string]interface{})
		if ok {
			value := respBodyDataMap["note"].(string)
			response.Code = respBody.Code
			response.HttpCode = httpStatus
			response.Data = value
			response.Msg = respBody.Msg
			return response
		}
		response.Code = respBody.Code
		response.HttpCode = httpStatus
		response.Data = respBody.Msg
		response.Msg = respBody.Msg
		return response

	}
	utils.LogError(fmt.Sprintf("cannot parse response reading param from rnode %s", url))
	response.Code = code.ScodeJsonError
	response.HttpCode = http.StatusInternalServerError
	return response
}
