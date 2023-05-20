package chainstoragesdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ipfs/go-cid"
	"github.com/kataras/golog"
	"github.com/paradeum-team/chainstorage-sdk/code"
	"github.com/paradeum-team/chainstorage-sdk/model"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type Object struct {
	Config *Configuration
	Client *RestyClient
	logger *golog.Logger
}

// region 对象数据

// 获取对象数据列表
func (o *Object) GetObjectList(bucketId int, objectItem string, pageSize, pageIndex int) (model.ObjectPageResponse, error) {
	response := model.ObjectPageResponse{}

	// 参数设置
	urlQuery := ""
	if bucketId <= 0 {
		return response, code.ErrInvalidBucketId
	}
	urlQuery += fmt.Sprintf("bucketId=%d&", bucketId)

	if len(objectItem) != 0 {
		urlQuery += fmt.Sprintf("objectItem=%s&", url.QueryEscape(objectItem))
	}

	if pageSize > 0 && pageSize <= 1000 {
		urlQuery += fmt.Sprintf("pageSize=%d&", pageSize)
	}

	if pageIndex > 0 && pageIndex <= 1000 {
		urlQuery += fmt.Sprintf("pageIndex=%d&", pageIndex)
	}

	// 请求Url
	urlQuery = strings.TrimSuffix(urlQuery, "&")
	apiBaseAddress := o.Config.ChainStorageApiBaseAddress
	apiPath := "api/v1/objects/search"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	if len(urlQuery) != 0 {
		apiUrl += "?" + urlQuery
	}

	// API调用
	httpStatus, body, err := o.Client.RestyGet(apiUrl)
	if err != nil {
		//utils.LogError(fmt.Sprintf("API:GetObjectList:HttpGet, apiUrl:%s, httpStatus:%d, err:%+v\n", apiUrl, httpStatus, err))
		o.logger.Errorf(fmt.Sprintf("API:GetObjectList:HttpGet, apiUrl:%s, httpStatus:%d, err:%+v\n", apiUrl, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		o.logger.Errorf(fmt.Sprintf("API:GetObjectList:HttpGet, apiUrl:%s, httpStatus:%d, body:%s\n", apiUrl, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		o.logger.Errorf(fmt.Sprintf("API:GetObjectList:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 删除对象数据
func (o *Object) RemoveObject(objectIds []int) (model.ObjectRemoveResponse, error) {
	response := model.ObjectRemoveResponse{}

	// 参数设置
	if len(objectIds) == 0 {
		return response, code.ErrInvalidObjectIds
	}

	params := map[string]interface{}{
		"objectIds": objectIds,
	}

	// 请求Url
	apiBaseAddress := o.Config.ChainStorageApiBaseAddress
	apiPath := "api/v1/object"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := o.Client.RestyDelete(apiUrl, params)
	if err != nil {
		o.logger.Errorf(fmt.Sprintf("API:RemoveObject:HttpDelete, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		o.logger.Errorf(fmt.Sprintf("API:RemoveObject:HttpDelete, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		o.logger.Errorf(fmt.Sprintf("API:RemoveObject:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 重命名对象数据
func (o *Object) RenameObject(objectId int, objectName string, isOverwrite bool) (model.ObjectRenameResponse, error) {
	response := model.ObjectRenameResponse{}

	// 参数设置
	if objectId <= 0 {
		return response, code.ErrInvalidObjectId
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
	apiBaseAddress := o.Config.ChainStorageApiBaseAddress
	apiPath := fmt.Sprintf("api/v1/object/name/%d", objectId)
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := o.Client.RestyPut(apiUrl, params)
	if err != nil {
		o.logger.Errorf(fmt.Sprintf("API:RenameObject:HttpPut, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		o.logger.Errorf(fmt.Sprintf("API:RenameObject:HttpPut, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		o.logger.Errorf(fmt.Sprintf("API:RenameObject:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 设置对象数据星标
func (o *Object) MarkObject(objectId int, isMarked bool) (model.ObjectMarkResponse, error) {
	response := model.ObjectMarkResponse{}

	// 参数设置
	if objectId <= 0 {
		return response, code.ErrInvalidObjectId
	}

	markObject := 0
	if isMarked {
		markObject = 1
	}

	params := map[string]interface{}{
		"isMarked": markObject,
	}

	// 请求Url
	apiBaseAddress := o.Config.ChainStorageApiBaseAddress
	apiPath := fmt.Sprintf("api/v1/object/mark/%d", objectId)
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := o.Client.RestyPut(apiUrl, params)
	if err != nil {
		o.logger.Errorf(fmt.Sprintf("API:MarkObject:HttpPut, apiUrl:%s, params:%+v, httpStatus:%d, err:%+v\n", apiUrl, params, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		o.logger.Errorf(fmt.Sprintf("API:MarkObject:HttpPut, apiUrl:%s, params:%+v, httpStatus:%d, body:%s\n", apiUrl, params, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		o.logger.Errorf(fmt.Sprintf("API:MarkObject:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 根据CID检查是否已经存在Object
func (o *Object) IsExistObjectByCid(objectCid string) (model.ObjectExistResponse, error) {
	response := model.ObjectExistResponse{}

	// 参数设置
	if len(objectCid) <= 0 {
		return response, code.ErrInvalidObjectCid
	}

	// CID检查
	_, err := cid.Decode(objectCid)
	if err != nil {
		return response, code.ErrInvalidObjectCid
	}

	urlQuery := url.QueryEscape(objectCid)

	// 请求Url
	apiBaseAddress := o.Config.ChainStorageApiBaseAddress
	apiPath := fmt.Sprintf("api/v1/object/existCid/%s", urlQuery)
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := o.Client.RestyGet(apiUrl)
	if err != nil {
		o.logger.Errorf(fmt.Sprintf("API:IsExistObjectByCid:HttpGet, apiUrl:%s, httpStatus:%d, err:%+v\n", apiUrl, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		o.logger.Errorf(fmt.Sprintf("API:IsExistObjectByCid:HttpGet, apiUrl:%s, httpStatus:%d, body:%s\n", apiUrl, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		o.logger.Errorf(fmt.Sprintf("API:IsExistObjectByCid:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// 根据CID检查是否已经存在Object
func (o *Object) GetObjectByName(bucketId int, objectName string) (model.ObjectCreateResponse, error) {
	response := model.ObjectCreateResponse{}

	// 参数设置
	urlQuery := "?"
	if bucketId <= 0 {
		return response, code.ErrInvalidBucketId
	}
	urlQuery += fmt.Sprintf("bucketId=%d&", bucketId)

	// 参数设置
	if len(objectName) <= 0 {
		return response, code.ErrInvalidObjectName
	}
	urlQuery += fmt.Sprintf("objectName=%s", url.QueryEscape(objectName))

	// 请求Url
	apiBaseAddress := o.Config.ChainStorageApiBaseAddress
	apiPath := fmt.Sprintf("api/v1/object/find/name%s", urlQuery)
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := o.Client.RestyGet(apiUrl)
	if err != nil {
		o.logger.Errorf(fmt.Sprintf("API:GetObjectByName:HttpGet, apiUrl:%s, httpStatus:%d, err:%+v\n", apiUrl, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		o.logger.Errorf(fmt.Sprintf("API:GetObjectByName:HttpGet, apiUrl:%s, httpStatus:%d, body:%s\n", apiUrl, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		o.logger.Errorf(fmt.Sprintf("API:GetObjectByName:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}

// endregion 对象数据

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
