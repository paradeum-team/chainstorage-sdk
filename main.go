package main

import (
	"chainstorage-sdk/internal/service"
	"fmt"
)

func main() {

	//apiName := "my/cool+blog&about,stuff"
	//pageSize := 10
	//pageIndex := 1
	//response, err := GetApiKeyList(apiName, pageSize, pageIndex)
	//if err != nil {
	//	fmt.Printf("error:%+v\n", err)
	//	return
	//}
	//
	//fmt.Printf("response:%+v\n", response)

	// region 桶数据

	// 获取桶数据列表
	bucketName := ""
	pageSize := 10
	pageIndex := 1
	response, err := service.GetBucketList(bucketName, pageSize, pageIndex)
	if err != nil {
		fmt.Printf("error:%+v\n", err)
		return
	}

	fmt.Printf("response:%+v\n", response)

	//// 创建桶数据
	//bucketName := "bucket3"
	//storageNetworkCode := 10001
	//bucketPrincipleCode := 10001
	//response, err := CreateBucket(bucketName, storageNetworkCode, bucketPrincipleCode)
	//if err != nil {
	//	fmt.Printf("error:%+v\n", err)
	//	return
	//}
	//
	//fmt.Printf("response:%+v\n", response)

	//// 清空桶数据
	//bucketId := 16
	//response, err := EmptyBucket(bucketId)
	//if err != nil {
	//	fmt.Printf("error:%+v\n", err)
	//	return
	//}
	//
	//fmt.Printf("response:%+v\n", response)

	//// 删除桶数据
	//bucketId := 16
	//response, err := RemoveBucket(bucketId)
	//if err != nil {
	//	fmt.Printf("error:%+v\n", err)
	//	return
	//}
	//
	//fmt.Printf("response:%+v\n", response)

	// endregion 桶数据

	//// region 对象数据
	//
	////// 获取对象数据列表
	////bucketId := 16
	////objectItem := ""
	////pageSize := 10
	////pageIndex := 1
	////response, err := GetObjectList(bucketId, objectItem, pageSize, pageIndex)
	////if err != nil {
	////	fmt.Printf("error:%+v\n", err)
	////	return
	////}
	////
	////fmt.Printf("response:%+v\n", response)
	//
	//// 删除对象数据
	//objectId := 90
	//objectIds := []int{objectId}
	//response, err := RemoveObject(objectIds)
	//if err != nil {
	//	fmt.Printf("error:%+v\n", err)
	//	return
	//}
	//
	//fmt.Printf("response:%+v\n", response)
	//
	////// 重命名对象数据
	////objectId := 90
	////objectName := "objectName3"
	////isOverwrite := true
	////response, err := RenameObject(objectId, objectName, isOverwrite)
	////if err != nil {
	////	fmt.Printf("error:%+v\n", err)
	////	return
	////}
	////
	////fmt.Printf("response:%+v\n", response)
	//
	////// 设置对象数据星标
	////objectId := 90
	////isMarked := true
	////response, err := MarkObject(objectId, isMarked)
	////if err != nil {
	////	fmt.Printf("error:%+v\n", err)
	////	return
	////}
	////
	////fmt.Printf("response:%+v\n", response)
	//
	//// endregion 对象数据
}

//// 获取API-Key数据列表
//func GetApiKeyList(apiName string, pageSize, pageIndex int) (model.ApiKeyPageResponse, error) {
//	response := model.ApiKeyPageResponse{}
//
//	// 参数设置
//	urlQuery := ""
//	if len(apiName) != 0 {
//		urlQuery += fmt.Sprintf("apiName=%s&", url.QueryEscape(apiName))
//	}
//
//	if pageSize > 0 && pageSize <= 1000 {
//		urlQuery += fmt.Sprintf("pageSize=%d&", pageSize)
//	}
//
//	if pageIndex > 0 && pageIndex <= 65535 {
//		urlQuery += fmt.Sprintf("pageIndex=%d&", pageIndex)
//	}
//
//	// 请求Url
//	urlQuery = strings.TrimSuffix(urlQuery, "&")
//	apiBaseAddress := conf.Config.LinkedStorageApiBaseAddress
//	apiPath := "api/v1/apiKeys"
//	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)
//
//	if len(urlQuery) != 0 {
//		apiUrl += "?" + urlQuery
//	}
//
//	// API调用
//	httpStatus, body, err := client.RestyGet(apiUrl)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:GetApiKeyList:HttpGet, apiUrl:%s, httpStatus:%d, err:%+v\n", apiUrl, httpStatus, err))
//
//		return response, err
//	}
//
//	if httpStatus != http.StatusOK {
//		utils.LogError(fmt.Sprintf("API:GetApiKeyList:HttpGet, apiUrl:%s, httpStatus:%d, body:%s\n", apiUrl, httpStatus, string(body)))
//
//		return response, errors.New(string(body))
//	}
//
//	// 响应数据解析
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:GetApiKeyList:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))
//
//		return response, err
//	}
//
//	return response, nil
//}

//// region 桶数据
//
//// 获取桶数据列表
//func GetBucketList(bucketName string, pageSize, pageIndex int) (model.BucketPageResponse, error) {
//	response := model.BucketPageResponse{}
//
//	// 参数设置
//	urlQuery := ""
//	if len(bucketName) != 0 {
//		urlQuery += fmt.Sprintf("bucketName=%s&", url.QueryEscape(bucketName))
//	}
//
//	if pageSize > 0 && pageSize <= 1000 {
//		urlQuery += fmt.Sprintf("pageSize=%d&", pageSize)
//	}
//
//	if pageIndex > 0 && pageIndex <= 65535 {
//		urlQuery += fmt.Sprintf("pageIndex=%d&", pageIndex)
//	}
//
//	// 请求Url
//	urlQuery = strings.TrimSuffix(urlQuery, "&")
//	apiBaseAddress := conf.Config.LinkedStorageApiBaseAddress
//	apiPath := "api/v1/buckets"
//	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)
//
//	if len(urlQuery) != 0 {
//		apiUrl += "?" + urlQuery
//	}
//
//	// API调用
//	httpStatus, body, err := client.RestyGet(apiUrl)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:GetBucketList:HttpGet, apiUrl:%s, httpStatus:%d, err:%+v\n", apiUrl, httpStatus, err))
//
//		return response, err
//	}
//
//	if httpStatus != http.StatusOK {
//		utils.LogError(fmt.Sprintf("API:GetBucketList:HttpGet, apiUrl:%s, httpStatus:%d, body:%s\n", apiUrl, httpStatus, string(body)))
//
//		return response, errors.New(string(body))
//	}
//
//	// 响应数据解析
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:GetBucketList:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))
//
//		return response, err
//	}
//
//	return response, nil
//}
//
//// 创建桶数据
//func CreateBucket(bucketName string, storageNetworkCode, bucketPrincipleCode int) (model.BucketCreateResponse, error) {
//	response := model.BucketCreateResponse{}
//
//	// 参数设置
//	//todo: bucket name check?
//	//todo: storageNetworkCode check?
//	//todo: bucketPrincipleCode check?
//
//	params := map[string]interface{}{
//		"bucketName":          bucketName,
//		"storageNetworkCode":  storageNetworkCode,
//		"bucketPrincipleCode": bucketPrincipleCode,
//	}
//
//	// 请求Url
//	apiBaseAddress := conf.Config.LinkedStorageApiBaseAddress
//	apiPath := "api/v1/bucket"
//	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)
//
//	// API调用
//	httpStatus, body, err := client.RestyPost(apiUrl, params)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:CreateBucket:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))
//
//		return response, err
//	}
//
//	if httpStatus != http.StatusOK {
//		utils.LogError(fmt.Sprintf("API:CreateBucket:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))
//
//		return response, errors.New(string(body))
//	}
//
//	// 响应数据解析
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:CreateBucket:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))
//
//		return response, err
//	}
//
//	return response, nil
//}
//
//// 清空桶数据
//func EmptyBucket(bucketId int) (model.BucketEmptyResponse, error) {
//	response := model.BucketEmptyResponse{}
//
//	// 参数设置
//	//todo: bucket id check?
//
//	params := map[string]interface{}{
//		"id": bucketId,
//	}
//
//	// 请求Url
//	apiBaseAddress := conf.Config.LinkedStorageApiBaseAddress
//	apiPath := "api/v1/bucket/status/clean"
//	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)
//
//	// API调用
//	httpStatus, body, err := client.RestyPost(apiUrl, params)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:EmptyBucket:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))
//
//		return response, err
//	}
//
//	if httpStatus != http.StatusOK {
//		utils.LogError(fmt.Sprintf("API:EmptyBucket:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))
//
//		return response, errors.New(string(body))
//	}
//
//	// 响应数据解析
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:EmptyBucket:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))
//
//		return response, err
//	}
//
//	return response, nil
//}
//
//// 删除桶数据
//func RemoveBucket(bucketId int) (model.BucketRemoveResponse, error) {
//	response := model.BucketRemoveResponse{}
//
//	// 参数设置
//	//todo: bucket id check?
//
//	// 请求Url
//	apiBaseAddress := conf.Config.LinkedStorageApiBaseAddress
//	apiPath := fmt.Sprintf("api/v1/bucket/%d", bucketId)
//	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)
//
//	// API调用
//	httpStatus, body, err := client.RestyDelete(apiUrl, nil)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:RemoveBucket:HttpDelete, apiUrl:%s, bucketId:%d, httpStatus:%d, err:%+v\n", apiUrl, bucketId, httpStatus, err))
//
//		return response, err
//	}
//
//	if httpStatus != http.StatusOK {
//		utils.LogError(fmt.Sprintf("API:RemoveBucket:HttpDelete, apiUrl:%s, bucketId:%d, httpStatus:%d, body:%s\n", apiUrl, bucketId, httpStatus, string(body)))
//
//		return response, errors.New(string(body))
//	}
//
//	// 响应数据解析
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:RemoveBucket:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))
//
//		return response, err
//	}
//
//	return response, nil
//}
//
//// endregion 桶数据
//
//// region 对象数据
//
//// 获取对象数据列表
//func GetObjectList(bucketId int, objectItem string, pageSize, pageIndex int) (model.ObjectPageResponse, error) {
//	response := model.ObjectPageResponse{}
//
//	// 参数设置
//	urlQuery := ""
//	//todo: bucket id check?
//	//if bucketId != 0 {
//	//	urlQuery += fmt.Sprintf("bucketId=%s&", bucketId)
//	//}
//	urlQuery += fmt.Sprintf("bucketId=%d&", bucketId)
//
//	if len(objectItem) != 0 {
//		urlQuery += fmt.Sprintf("objectItem=%s&", url.QueryEscape(objectItem))
//	}
//
//	if pageSize > 0 && pageSize <= 1000 {
//		urlQuery += fmt.Sprintf("pageSize=%d&", pageSize)
//	}
//
//	if pageIndex > 0 && pageIndex <= 65535 {
//		urlQuery += fmt.Sprintf("pageIndex=%d&", pageIndex)
//	}
//
//	// 请求Url
//	urlQuery = strings.TrimSuffix(urlQuery, "&")
//	apiBaseAddress := conf.Config.LinkedStorageApiBaseAddress
//	apiPath := "api/v1/objects/search"
//	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)
//
//	if len(urlQuery) != 0 {
//		apiUrl += "?" + urlQuery
//	}
//
//	// API调用
//	httpStatus, body, err := client.RestyGet(apiUrl)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:GetObjectList:HttpGet, apiUrl:%s, httpStatus:%d, err:%+v\n", apiUrl, httpStatus, err))
//
//		return response, err
//	}
//
//	if httpStatus != http.StatusOK {
//		utils.LogError(fmt.Sprintf("API:GetObjectList:HttpGet, apiUrl:%s, httpStatus:%d, body:%s\n", apiUrl, httpStatus, string(body)))
//
//		return response, errors.New(string(body))
//	}
//
//	// 响应数据解析
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:GetObjectList:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))
//
//		return response, err
//	}
//
//	return response, nil
//}
//
//// 删除对象数据
//func RemoveObject(objectIds []int) (model.ObjectRemoveResponse, error) {
//	response := model.ObjectRemoveResponse{}
//
//	// 参数设置
//	//todo: object id list check?
//	params := map[string]interface{}{
//		"objectIds": objectIds,
//	}
//
//	// 请求Url
//	apiBaseAddress := conf.Config.LinkedStorageApiBaseAddress
//	apiPath := "api/v1/object"
//	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)
//
//	// API调用
//	httpStatus, body, err := client.RestyDelete(apiUrl, params)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:RemoveObject:HttpDelete, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))
//
//		return response, err
//	}
//
//	if httpStatus != http.StatusOK {
//		utils.LogError(fmt.Sprintf("API:RemoveObject:HttpDelete, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))
//
//		return response, errors.New(string(body))
//	}
//
//	// 响应数据解析
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:RemoveObject:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))
//
//		return response, err
//	}
//
//	return response, nil
//}
//
//// 重命名对象数据
//func RenameObject(objectId int, objectName string, isOverwrite bool) (model.ObjectRenameResponse, error) {
//	response := model.ObjectRenameResponse{}
//
//	// 参数设置
//	//todo: object id check?
//	//todo: object name check?
//
//	forceOverwrite := 0
//	if isOverwrite {
//		forceOverwrite = 1
//	}
//
//	params := map[string]interface{}{
//		"objectName":  objectName,
//		"isOverwrite": forceOverwrite,
//	}
//
//	// 请求Url
//	apiBaseAddress := conf.Config.LinkedStorageApiBaseAddress
//	apiPath := fmt.Sprintf("api/v1/object/name/%d", objectId)
//	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)
//
//	// API调用
//	httpStatus, body, err := client.RestyPut(apiUrl, params)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:RenameObject:HttpPut, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))
//
//		return response, err
//	}
//
//	if httpStatus != http.StatusOK {
//		utils.LogError(fmt.Sprintf("API:RenameObject:HttpPut, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))
//
//		return response, errors.New(string(body))
//	}
//
//	// 响应数据解析
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:RenameObject:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))
//
//		return response, err
//	}
//
//	return response, nil
//}
//
//// 设置对象数据星标
//func MarkObject(objectId int, isMarked bool) (model.ObjectMarkResponse, error) {
//	response := model.ObjectMarkResponse{}
//
//	// 参数设置
//	//todo: object id check?
//
//	markObject := 0
//	if isMarked {
//		markObject = 1
//	}
//
//	params := map[string]interface{}{
//		"isMarked": markObject,
//	}
//
//	// 请求Url
//	apiBaseAddress := conf.Config.LinkedStorageApiBaseAddress
//	apiPath := fmt.Sprintf("api/v1/object/mark/%d", objectId)
//	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)
//
//	// API调用
//	httpStatus, body, err := client.RestyPut(apiUrl, params)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:MarkObject:HttpPut, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))
//
//		return response, err
//	}
//
//	if httpStatus != http.StatusOK {
//		utils.LogError(fmt.Sprintf("API:MarkObject:HttpPut, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))
//
//		return response, errors.New(string(body))
//	}
//
//	// 响应数据解析
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		utils.LogError(fmt.Sprintf("API:MarkObject:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))
//
//		return response, err
//	}
//
//	return response, nil
//}
//
//// endregion 对象数据
