package service

import (
	client "chainstorage-sdk/base/tokenclient"
	consts "chainstorage-sdk/common/dict"
	"chainstorage-sdk/common/e/code"
	"chainstorage-sdk/conf"
	"chainstorage-sdk/model"
	"chainstorage-sdk/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// region 桶数据

// 获取桶数据列表
func GetBucketList(bucketName string, pageSize, pageIndex int) (model.BucketPageResponse, error) {
	response := model.BucketPageResponse{}

	// 参数设置
	urlQuery := ""
	if len(bucketName) != 0 {
		if err := checkBucketName(bucketName); err != nil {
			return response, err
		}

		urlQuery += fmt.Sprintf("bucketName=%s&", url.QueryEscape(bucketName))
	}

	if pageSize > 0 && pageSize <= 1000 {
		urlQuery += fmt.Sprintf("pageSize=%d&", pageSize)
	}

	if pageIndex > 0 && pageIndex <= 65535 {
		urlQuery += fmt.Sprintf("pageIndex=%d&", pageIndex)
	}

	// 请求Url
	urlQuery = strings.TrimSuffix(urlQuery, "&")
	apiBaseAddress := conf.Config.LinkedStorageApiBaseAddress
	apiPath := "api/v1/buckets"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	if len(urlQuery) != 0 {
		apiUrl += "?" + urlQuery
	}

	// API调用
	httpStatus, body, err := client.RestyGet(apiUrl)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:GetBucketList:HttpGet, apiUrl:%s, httpStatus:%d, err:%+v\n", apiUrl, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		utils.LogError(fmt.Sprintf("API:GetBucketList:HttpGet, apiUrl:%s, httpStatus:%d, body:%s\n", apiUrl, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:GetBucketList:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 创建桶数据
func CreateBucket(bucketName string, storageNetworkCode, bucketPrincipleCode int) (model.BucketCreateResponse, error) {
	response := model.BucketCreateResponse{}

	// 参数设置
	if err := checkBucketName(bucketName); err != nil {
		return response, err
	}

	if err := checkStorageNetworkCode(storageNetworkCode); err != nil {
		return response, err
	}

	if err := checkBucketPrincipleCode(bucketPrincipleCode); err != nil {
		return response, err
	}

	params := map[string]interface{}{
		"bucketName":          bucketName,
		"storageNetworkCode":  storageNetworkCode,
		"bucketPrincipleCode": bucketPrincipleCode,
	}

	// 请求Url
	apiBaseAddress := conf.Config.LinkedStorageApiBaseAddress
	apiPath := "api/v1/bucket"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := client.RestyPost(apiUrl, params)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:CreateBucket:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		utils.LogError(fmt.Sprintf("API:CreateBucket:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:CreateBucket:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 清空桶数据
func EmptyBucket(bucketId int) (model.BucketEmptyResponse, error) {
	response := model.BucketEmptyResponse{}

	// 参数设置
	if bucketId <= 0 {
		return response, errors.New("请输入正确的桶ID.")
	}

	params := map[string]interface{}{
		"id": bucketId,
	}

	// 请求Url
	apiBaseAddress := conf.Config.LinkedStorageApiBaseAddress
	apiPath := "api/v1/bucket/status/clean"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := client.RestyPost(apiUrl, params)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:EmptyBucket:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		utils.LogError(fmt.Sprintf("API:EmptyBucket:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:EmptyBucket:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 删除桶数据
func RemoveBucket(bucketId int) (model.BucketRemoveResponse, error) {
	response := model.BucketRemoveResponse{}

	// 参数设置
	if bucketId <= 0 {
		return response, errors.New("请输入正确的桶ID.")
	}

	// 请求Url
	apiBaseAddress := conf.Config.LinkedStorageApiBaseAddress
	apiPath := fmt.Sprintf("api/v1/bucket/%d", bucketId)
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := client.RestyDelete(apiUrl, nil)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:RemoveBucket:HttpDelete, apiUrl:%s, bucketId:%d, httpStatus:%d, err:%+v\n", apiUrl, bucketId, httpStatus, err))
		return response, err
	}

	if httpStatus != http.StatusOK {
		utils.LogError(fmt.Sprintf("API:RemoveBucket:HttpDelete, apiUrl:%s, bucketId:%d, httpStatus:%d, body:%s\n", apiUrl, bucketId, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:RemoveBucket:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 检查桶名称
func checkBucketName(bucketName string) error {
	if len(bucketName) < 3 || len(bucketName) > 63 {
		return code.ErrInvalidBucketName
	}

	// 桶名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试
	isMatch := regexp.MustCompile(`^[a-z0-9-]*$`).MatchString(bucketName)
	if !isMatch {
		return code.ErrInvalidBucketName
	}

	//// 存储网络编码必须设置
	//storageNetworkCode := req.StorageNetworkCode
	//storageNetworkCodeMapping := consts.StorageNetworkCodeMapping
	//_, exist := storageNetworkCodeMapping[storageNetworkCode]
	//if !exist {
	//	return code.ErrStorageNetworkMustSet
	//}
	//
	//// 桶策略编码必须设置
	//bucketPrincipleCode := req.BucketPrincipleCode
	//bucketPrincipleCodeMapping := consts.BucketPrincipleCodeMapping
	//_, exist = bucketPrincipleCodeMapping[bucketPrincipleCode]
	//if !exist {
	//	return code.ErrBucketPrincipleMustSet
	//}

	return nil
}

// 检查存储网络编码
func checkStorageNetworkCode(storageNetworkCode int) error {
	// 存储网络编码必须设置
	storageNetworkCodeMapping := consts.StorageNetworkCodeMapping
	_, exist := storageNetworkCodeMapping[storageNetworkCode]
	if !exist {
		return code.ErrStorageNetworkMustSet
	}

	return nil
}

// 检查存储网络编码
func checkBucketPrincipleCode(bucketPrincipleCode int) error {
	// 桶策略编码必须设置
	bucketPrincipleCodeMapping := consts.BucketPrincipleCodeMapping
	_, exist := bucketPrincipleCodeMapping[bucketPrincipleCode]
	if !exist {
		return code.ErrBucketPrincipleMustSet
	}

	return nil
}

//func checkBeforeBucketInsert(req dto.BucketInsertReq, s service.Bucket) error {
//	bucketName := req.BucketName
//	// 桶名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试
//	//if len([]rune(bucketName)) < 3 || len([]rune(bucketName)) > 63 {
//	//	return code.ErrInvalidBucketName
//	//}
//	if len(bucketName) < 3 || len(bucketName) > 63 {
//		return code.ErrInvalidBucketName
//	}
//
//	// 桶名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试
//	isMatch := regexp.MustCompile(`^[a-z0-9-]*$`).MatchString(bucketName)
//	if !isMatch {
//		return code.ErrInvalidBucketName
//	}
//
//	// 桶名称冲突，桶名称必须全平台唯一，请重新尝试
//	getReq := dto.BucketGetReq{BucketName: bucketName}
//	exists, err := s.IsExistSameBucketName(&getReq)
//	if err != nil {
//		return err
//	}
//
//	if exists {
//		return code.ErrBucketNameConflict
//	}
//
//	// 存储网络编码必须设置
//	storageNetworkCode := req.StorageNetworkCode
//	storageNetworkCodeMapping := consts.StorageNetworkCodeMapping
//	_, exist := storageNetworkCodeMapping[storageNetworkCode]
//	if !exist {
//		return code.ErrStorageNetworkMustSet
//	}
//
//	// 桶策略编码必须设置
//	bucketPrincipleCode := req.BucketPrincipleCode
//	bucketPrincipleCodeMapping := consts.BucketPrincipleCodeMapping
//	_, exist = bucketPrincipleCodeMapping[bucketPrincipleCode]
//	if !exist {
//		return code.ErrBucketPrincipleMustSet
//	}
//
//	// 当是基础版本的时候，一个网络类型下，只允许建一个桶
//	userQuotaService := service.UserQuota{}
//	userQuotaService.Orm = s.Orm
//	userQuotaService.Identity = s.Identity
//
//	var object models.UserQuota
//	userId := s.Identity.Id
//	err = userQuotaService.GetQuotaByUserNetType(userId, consts.StorageNetworkCodeIpfs, &object)
//	if err != nil {
//		return err
//	}
//
//	// 获取系统版本，确认是否是基础版
//	isBasicVersion := false
//	// todo: 增加套餐1常量
//	isBasicVersion = object.PackagePlanId == 1
//
//	if isBasicVersion {
//		getCondition := dto.BucketGetCondition{}
//		getCondition.StorageNetworkCode = storageNetworkCode
//		exist, err := s.IsExistSameStorageNetwork(&getCondition)
//		if err != nil {
//			return err
//		}
//
//		if exist {
//			return code.ErrOnlyCreate1BucketForSameStorageNetwork
//		}
//	}
//
//	return nil
//}

// endregion 桶数据
