package chainstoragesdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/paradeum-team/chainstorage-sdk/sdk/model"
	"net/http"
	"sync"
)

type SdkConfig Configuration

var once = sync.Once{}
var mClient *CssClient

type CssClient struct {
	Config     *Configuration
	Logger     *PldLogger
	httpClient *RestyClient
	Bucket     *Bucket
	Object     *Object
	Car        *Car
}

func newClient(config *ApplicationConfig) (*CssClient, error) {
	var err error
	once.Do(func() {
		initConfig(config)

		mClient = &CssClient{}
		mClient.Config = &cssConfig

		mClient.Logger = GetLogger(&cssLoggerConfig)

		mClient.httpClient = &RestyClient{Config: mClient.Config}
		mClient.Bucket = &Bucket{Config: mClient.Config, Client: mClient.httpClient, logger: mClient.Logger.logger}
		mClient.Object = &Object{Config: mClient.Config, Client: mClient.httpClient, logger: mClient.Logger.logger}
		mClient.Car = &Car{Config: mClient.Config, Client: mClient.httpClient, logger: mClient.Logger.logger}
	})

	//mClient.Logger.logger.Error("client new.")
	return mClient, err
}

func New(config *ApplicationConfig) (*CssClient, error) {

	return newClient(config)
}

func (c *CssClient) GetIpfsVersion() (model.VersionResponse, error) {
	response := model.VersionResponse{}

	// 请求Url
	apiBaseAddress := c.Config.ChainStorageApiEndpoint
	apiPath := "ipfsVersion"
	apiUrl := fmt.Sprintf("%s%s", apiBaseAddress, apiPath)

	// API调用
	httpStatus, body, err := c.httpClient.RestyGet(apiUrl)
	if err != nil {
		c.Logger.logger.Errorf(fmt.Sprintf("API:GetObjectByName:HttpGet, apiUrl:%s, httpStatus:%d, err:%+v\n", apiUrl, httpStatus, err))

		return response, err
	}

	if httpStatus != http.StatusOK {
		c.Logger.logger.Errorf(fmt.Sprintf("API:GetObjectByName:HttpGet, apiUrl:%s, httpStatus:%d, body:%s\n", apiUrl, httpStatus, string(body)))

		return response, errors.New(string(body))
	}

	// 响应数据解析
	err = json.Unmarshal(body, &response)
	if err != nil {
		c.Logger.logger.Errorf(fmt.Sprintf("API:GetObjectByName:JsonUnmarshal, body:%s, err:%+v\n", string(body), err))

		return response, err
	}

	return response, nil
}
