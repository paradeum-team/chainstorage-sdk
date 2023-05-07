package service

import (
	client "chainstorage-sdk/base/tokenclient"
	"chainstorage-sdk/common/e/code"
	"chainstorage-sdk/conf"
	"chainstorage-sdk/model"
	"chainstorage-sdk/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ipfs/go-cid"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// region 对象数据

// 获取对象数据列表
func GetObjectList(bucketId int, objectItem string, pageSize, pageIndex int) (model.ObjectPageResponse, error) {
	response := model.ObjectPageResponse{}

	// 参数设置
	urlQuery := ""
	if bucketId <= 0 {
		return response, errors.New("请输入正确的桶ID.")
	}
	urlQuery += fmt.Sprintf("bucketId=%d&", bucketId)

	if len(objectItem) != 0 {
		urlQuery += fmt.Sprintf("objectItem=%s&", url.QueryEscape(objectItem))
	}

	if pageSize > 0 && pageSize <= 1000 {
		urlQuery += fmt.Sprintf("pageSize=%d&", pageSize)
	}

	if pageIndex > 0 && pageIndex <= 65535 {
		urlQuery += fmt.Sprintf("pageIndex=%d&", pageIndex)
	}

	// 请求Url
	urlQuery = strings.TrimSuffix(urlQuery, "&")
	apiBaseAddress := conf.Config.ChainStorageApiBaseAddress
	apiPath := "api/v1/objects/search"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	if len(urlQuery) != 0 {
		apiUrl += "?" + urlQuery
	}

	// API调用
	httpStatus, body, err := client.RestyGet(apiUrl)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:GetObjectList:HttpGet, apiUrl:%s, httpStatus:%d, err:%+v\n", apiUrl, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		utils.LogError(fmt.Sprintf("API:GetObjectList:HttpGet, apiUrl:%s, httpStatus:%d, body:%s\n", apiUrl, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:GetObjectList:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 删除对象数据
func RemoveObject(objectIds []int) (model.ObjectRemoveResponse, error) {
	response := model.ObjectRemoveResponse{}

	// 参数设置
	if len(objectIds) == 0 {
		return response, errors.New("请输入正确的对象ID列表.")
	}

	params := map[string]interface{}{
		"objectIds": objectIds,
	}

	// 请求Url
	apiBaseAddress := conf.Config.ChainStorageApiBaseAddress
	apiPath := "api/v1/object"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := client.RestyDelete(apiUrl, params)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:RemoveObject:HttpDelete, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		utils.LogError(fmt.Sprintf("API:RemoveObject:HttpDelete, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:RemoveObject:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 重命名对象数据
func RenameObject(objectId int, objectName string, isOverwrite bool) (model.ObjectRenameResponse, error) {
	response := model.ObjectRenameResponse{}

	// 参数设置
	if objectId <= 0 {
		return response, errors.New("请输入正确的对象ID.")
	}

	if err := checkObjectName(objectName); err != nil {
		return response, err
	}

	forceOverwrite := 0
	if isOverwrite {
		forceOverwrite = 1
	}

	params := map[string]interface{}{
		"objectName":  objectName,
		"isOverwrite": forceOverwrite,
	}

	// 请求Url
	apiBaseAddress := conf.Config.ChainStorageApiBaseAddress
	apiPath := fmt.Sprintf("api/v1/object/name/%d", objectId)
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := client.RestyPut(apiUrl, params)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:RenameObject:HttpPut, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		utils.LogError(fmt.Sprintf("API:RenameObject:HttpPut, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:RenameObject:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 设置对象数据星标
func MarkObject(objectId int, isMarked bool) (model.ObjectMarkResponse, error) {
	response := model.ObjectMarkResponse{}

	// 参数设置
	if objectId <= 0 {
		return response, errors.New("请输入正确的对象ID.")
	}

	markObject := 0
	if isMarked {
		markObject = 1
	}

	params := map[string]interface{}{
		"isMarked": markObject,
	}

	// 请求Url
	apiBaseAddress := conf.Config.ChainStorageApiBaseAddress
	apiPath := fmt.Sprintf("api/v1/object/mark/%d", objectId)
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := client.RestyPut(apiUrl, params)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:MarkObject:HttpPut, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		utils.LogError(fmt.Sprintf("API:MarkObject:HttpPut, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:MarkObject:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 检查对象名称
func checkObjectName(objectName string) error {
	if len(objectName) == 0 || len(objectName) > 255 {
		return code.ErrInvalidObjectName
	}

	isMatch := regexp.MustCompile("[<>:\"/\\|?*\u0000-\u001F]").MatchString(objectName)
	if isMatch {
		return code.ErrInvalidObjectName
	}

	isMatch = regexp.MustCompile(`^(con|prn|aux|nul|com\d|lpt\d)$`).MatchString(objectName)
	if isMatch {
		return code.ErrInvalidObjectName
	}

	if objectName == "." || objectName == ".." {
		return code.ErrInvalidObjectName
	}

	return nil
}

// 根据CID检查是否已经存在Object
func IsExistObjectByCid(objectCid string) (model.ObjectExistResponse, error) {
	response := model.ObjectExistResponse{}

	// 参数设置
	if len(objectCid) <= 0 {
		return response, errors.New("请输入正确的对象CID.")
	}

	// CID检查
	_, err := cid.Decode(objectCid)
	if err != nil {
		return response, errors.New("请输入正确的对象CID.")
	}

	urlQuery := url.QueryEscape(objectCid)

	// 请求Url
	apiBaseAddress := conf.Config.ChainStorageApiBaseAddress
	apiPath := fmt.Sprintf("api/v1/object/existCid/%s", urlQuery)
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := client.RestyGet(apiUrl)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:IsExistObjectByCid:HttpGet, apiUrl:%s, httpStatus:%d, err:%+v\n", apiUrl, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		utils.LogError(fmt.Sprintf("API:IsExistObjectByCid:HttpGet, apiUrl:%s, httpStatus:%d, body:%s\n", apiUrl, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		utils.LogError(fmt.Sprintf("API:IsExistObjectByCid:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// endregion 对象数据
