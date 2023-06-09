package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kataras/golog"
	"github.com/paradeum-team/chainstorage-sdk/code"
	"github.com/paradeum-team/chainstorage-sdk/consts"
	"github.com/paradeum-team/chainstorage-sdk/model"
	"github.com/ulule/deepcopier"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type Bucket struct {
	Config *Configuration
	Client *RestyClient
	logger *golog.Logger
}

// region 桶数据

// 获取桶数据列表
func (b *Bucket) GetBucketList(bucketName string, pageSize, pageIndex int) (model.BucketPageResponse, error) {
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

	if pageIndex > 0 && pageIndex <= 1000 {
		urlQuery += fmt.Sprintf("pageIndex=%d&", pageIndex)
	}

	// 请求Url
	urlQuery = strings.TrimSuffix(urlQuery, "&")

	//apiBaseAddress := conf.cssConfig.ChainStorageApiEndpoint
	apiBaseAddress := b.Config.ChainStorageApiEndpoint
	apiPath := "api/v1/buckets"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	if len(urlQuery) != 0 {
		apiUrl += "?" + urlQuery
	}

	// API调用
	httpStatus, body, err := b.Client.RestyGet(apiUrl)
	//httpStatus, body, err := client.RestyGet(apiUrl)
	if err != nil {
		//utils.LogError(fmt.Sprintf("API:GetBucketList:HttpGet, apiUrl:%s, httpStatus:%d, err:%+v\n", apiUrl, httpStatus, err))
		b.logger.Errorf(fmt.Sprintf("API:GetBucketList:HttpGet, apiUrl:%s, httpStatus:%d, err:%+v\n", apiUrl, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		b.logger.Errorf(fmt.Sprintf("API:GetBucketList:HttpGet, apiUrl:%s, httpStatus:%d, body:%s\n", apiUrl, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		b.logger.Errorf(fmt.Sprintf("API:GetBucketList:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 创建桶数据
func (b *Bucket) CreateBucket(bucketName string, storageNetworkCode, bucketPrincipleCode int) (model.BucketCreateResponse, error) {
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
	apiBaseAddress := b.Config.ChainStorageApiEndpoint
	apiPath := "api/v1/bucket"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := b.Client.RestyPost(apiUrl, params)
	if err != nil {
		b.logger.Errorf(fmt.Sprintf("API:CreateBucket:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		b.logger.Errorf(fmt.Sprintf("API:CreateBucket:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		b.logger.Errorf(fmt.Sprintf("API:CreateBucket:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 清空桶数据
func (b *Bucket) EmptyBucket(bucketId int) (model.BucketEmptyResponse, error) {
	response := model.BucketEmptyResponse{}

	// 参数设置
	if bucketId <= 0 {
		return response, code.ErrInvalidBucketId
	}

	params := map[string]interface{}{
		"id": bucketId,
	}

	// 请求Url
	apiBaseAddress := b.Config.ChainStorageApiEndpoint
	apiPath := "api/v1/bucket/status/clean"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := b.Client.RestyPost(apiUrl, params)
	if err != nil {
		b.logger.Errorf(fmt.Sprintf("API:EmptyBucket:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		b.logger.Errorf(fmt.Sprintf("API:EmptyBucket:HttpPost, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		b.logger.Errorf(fmt.Sprintf("API:EmptyBucket:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 删除桶数据
func (b *Bucket) RemoveBucket(bucketId int, autoEmptyBucketData bool) (model.BucketRemoveResponse, error) {
	response := model.BucketRemoveResponse{}

	// 参数设置
	if bucketId <= 0 {
		return response, code.ErrInvalidBucketId
	}

	// 自动清空数据
	if autoEmptyBucketData {
		bucketEmptyResponse, err := b.EmptyBucket(bucketId)
		if err != nil {
			deepcopier.Copy(&bucketEmptyResponse).To(&response)
			b.logger.Errorf(fmt.Sprintf("API:RemoveBucket:EmptyBucket, bucketId:%d, bucketEmptyResponse:%+v, err:%+v\n", bucketId, bucketEmptyResponse, err))
			return response, err
		}
	}

	// 请求Url
	apiBaseAddress := b.Config.ChainStorageApiEndpoint
	apiPath := fmt.Sprintf("api/v1/bucket/%d", bucketId)
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := b.Client.RestyDelete(apiUrl, nil)
	if err != nil {
		b.logger.Errorf(fmt.Sprintf("API:RemoveBucket:HttpDelete, apiUrl:%s, bucketId:%d, httpStatus:%d, err:%+v\n", apiUrl, bucketId, httpStatus, err))
		return response, err
	}

	if httpStatus != http.StatusOK {
		b.logger.Errorf(fmt.Sprintf("API:RemoveBucket:HttpDelete, apiUrl:%s, bucketId:%d, httpStatus:%d, body:%s\n", apiUrl, bucketId, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		b.logger.Errorf(fmt.Sprintf("API:RemoveBucket:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 根据桶名称获取桶数据
func (b *Bucket) GetBucketByName(bucketName string) (model.BucketCreateResponse, error) {
	response := model.BucketCreateResponse{}

	// 参数设置
	if err := checkBucketName(bucketName); err != nil {
		return response, err
	}

	apiBaseAddress := b.Config.ChainStorageApiEndpoint
	apiPath := fmt.Sprintf("api/v1/bucket/name/%s", url.QueryEscape(bucketName))
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := b.Client.RestyGet(apiUrl)
	if err != nil {
		b.logger.Errorf(fmt.Sprintf("API:GetBucketByName:HttpGet, apiUrl:%s, httpStatus:%d, err:%+v\n", apiUrl, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		b.logger.Errorf(fmt.Sprintf("API:GetBucketByName:HttpGet, apiUrl:%s, httpStatus:%d, body:%s\n", apiUrl, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		b.logger.Errorf(fmt.Sprintf("API:GetBucketByName:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// endregion 桶数据

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
