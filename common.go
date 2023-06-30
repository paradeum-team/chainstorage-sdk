package sdk

//import (
//	"encoding/json"
//	"errors"
//	"fmt"
//	"github.com/kataras/golog"
//	"github.com/paradeum-team/chainstorage-sdk/code"
//	"github.com/paradeum-team/chainstorage-sdk/model"
//	"net/http"
//)
//
//type Common struct {
//	Config *Configuration
//	Client *RestyClient
//	logger *golog.Logger
//}
//
//// 按照存储类型获取Bucket容量统计
//func (c *Common) GetStorageNetworkBucketStat(StorageNetworkCode string) (model.BucketStorageTypeStatResp, error) {
//	response := model.BucketStorageTypeStatResp{}
//
//	// 参数设置
//	if len(StorageNetworkCode) <= 0 {
//		return response, code.ErrStorageNetworkCodeMustSet
//	}
//
//	// 请求Url
//	apiBaseAddress := c.Config.ChainStorageApiEndpoint
//	apiPath := fmt.Sprintf("api/v1/buckets/stat/%s", StorageNetworkCode)
//	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)
//
//	// API调用
//	httpStatus, body, err := c.Client.RestyGet(apiUrl)
//	if err != nil {
//		c.logger.Errorf(fmt.Sprintf("API:GetStorageNetworkBucketStat:HttpGet, apiUrl:%s, httpStatus:%d, err:%+v\n", apiUrl, httpStatus, err))
//
//		return response, err
//	}
//
//	if httpStatus != http.StatusOK {
//		c.logger.Errorf(fmt.Sprintf("API:GetStorageNetworkBucketStat:HttpGet, apiUrl:%s, httpStatus:%d, body:%s\n", apiUrl, httpStatus, string(body)))
//
//		return response, errors.New(string(body))
//	}
//
//	// 响应数据解析
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		c.logger.Errorf(fmt.Sprintf("API:GetStorageNetworkBucketStat:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))
//
//		return response, err
//	}
//
//	return response, nil
//}
