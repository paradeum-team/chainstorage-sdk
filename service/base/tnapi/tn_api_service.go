package tnapi

import (
	"chainstorage-sdk/base/client"
	"chainstorage-sdk/common/e/code"
	"chainstorage-sdk/conf"
	"chainstorage-sdk/conf/uri"
	"chainstorage-sdk/entity"
	"chainstorage-sdk/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

/*
*
TN健康检查
*/
func (*tnApiService) CheckSysHealthy(trackRoot string) (response entity.RnClusterResponse) {
	url := strings.Join([]string{trackRoot, uri.TN_SYS_HEALTHY_URL}, "")
	httpStatus, body, err := client.RestyGet(url)
	return handleRnClusterResponse(httpStatus, body, err)
}

/*
*
获取RN健康列表
*/
func (*tnApiService) GetRnCluster(trackRoot string) (response entity.RnClusterResponse) {
	defer utils.DevTimeTrack(time.Now(), "GetRnCluster")
	url := strings.Join([]string{trackRoot, conf.Config.RNodesInfoPath}, "")
	httpStatus, body, err := client.RestyGet(url)
	return handleRnClusterResponse(httpStatus, body, err)
}

/*
*
根据dgst查询Afid
*/
func (*tnApiService) TnQueryLocationByDgst(trackRoot, dgst string) (response entity.ResponseVO) {
	rawApi := fmt.Sprintf(uri.TN_QUERY_DGST_URL, dgst)
	tnQueryURL := strings.Join([]string{trackRoot, rawApi}, "")
	utils.DevLog("TnQueryLocationByDgst", fmt.Sprintf("requesting %s\n", tnQueryURL))
	httpStatus, body, err := client.RestyGet(tnQueryURL)
	return handleRnLocationByDgstResponse(httpStatus, body, err)

}

/*
*
根据dgst查询afid，兼容Rawid
*/
func (*tnApiService) TnQueryBigfileLocationByDgst(trackRoot, dgst string) (response entity.ResponseVO) {
	rawApi := fmt.Sprintf(uri.TN_LOCATION_SEEDID_BY_DGST_URL, dgst)
	tnQueryURL := strings.Join([]string{trackRoot, rawApi}, "")
	utils.DevLog("TnQueryBigfileLocationByDgst", fmt.Sprintf("requesting %s\n", tnQueryURL))
	httpStatus, body, err := client.RestyGet(tnQueryURL)
	return handleRnLocationByDgstResponse(httpStatus, body, err)
}

/*
*
根据dgst查询Afid，兼容不同传参方式
*/
func (*tnApiService) TnFileQueryLocationDgst2Afid(trackRoot, dgstType, dgst string) (response entity.ResponseVO) {
	rawApi := fmt.Sprintf(uri.TN_QUERY_DGST_URL, dgst)
	tnQueryURL := strings.Join([]string{trackRoot, rawApi}, "")
	utils.DevLog("TnFileQueryLocationDgst2Afid", fmt.Sprintf("requesting %s\n", tnQueryURL))
	httpStatus, body, err := client.RestyGet(tnQueryURL)
	return handleRnLocationByDgstResponse(httpStatus, body, err)
}

func handleRnClusterResponse(httpStatus int, body []byte, err error) (response entity.RnClusterResponse) {
	response.HttpCode = httpStatus
	if err == nil {
		if httpStatus >= http.StatusInternalServerError {
			response.Code = code.ScodeTnNetConnectionError
			response.Msg = code.GetMsg(code.ScodeTnNetConnectionError)
			utils.LogError(fmt.Sprintf("[TnQueryByDgst] body %s\n", string(body)))
			return response
		}
		err := json.Unmarshal(body, &response)
		if err != nil {
			utils.LogError(fmt.Sprintf("[GetRnCluster] cannot parse response reading rnodes info:%s \n", err.Error()))
			utils.LogError(fmt.Sprintf("[GetRnCluster] body is %s\n", string(body)))
			response.Code = code.ERROR
			response.Msg = err.Error()
			return response
		}
		return response
	} else {
		utils.LogError(fmt.Sprintf("[GetRnCluster] fetching rnode info  invalid status code:%d %s\n", httpStatus, err.Error()))
		response.Code = code.ScodeTrackingErr
		response.Msg = err.Error()
		return response
	}
}

func handleRnLocationByDgstResponse(httpStatus int, body []byte, err error) (response entity.ResponseVO) {
	response.HttpCode = httpStatus
	if err == nil {
		if httpStatus >= http.StatusInternalServerError {
			response.Code = code.ScodeTnNetConnectionError
			response.Msg = code.GetMsg(code.ScodeTnNetConnectionError)
			utils.LogError(fmt.Sprintf("[TnQueryByDgst] body %s", string(body)))
			return response
		}
		err := json.Unmarshal(body, &response)
		if err != nil {
			utils.LogError(fmt.Sprintf("[TnQueryByDgst] cannot parse response reading rnodes info:%s", err.Error()))
			utils.LogError(fmt.Sprintf("[TnQueryByDgst] body is %s", string(body)))
			response.Code = code.ERROR
			response.Msg = err.Error()
			return response
		}
		// 处理TN 寻址返回状态码
		if response.Code != code.ScodeOK {
			response.HttpCode = http.StatusNotFound
		}
		return response
	} else {
		if httpStatus == http.StatusGatewayTimeout || httpStatus == http.StatusServiceUnavailable {
			response.Code = code.ScodeAFCTimeout
		} else {
			response.Code = code.ScodeTnNetConnectionError
		}
		utils.LogError(fmt.Sprintf("[TnQueryByDgst] fetching rnode info  invalid status code:%d %s\n", httpStatus, err.Error()))
		response.Msg = err.Error()
		return response
	}
}
