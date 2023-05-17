package chainstoragesdk

import (
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

func newClient() (*CssClient, error) {
	var err error
	once.Do(func() {
		initConfig()

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

func New() (*CssClient, error) {

	return newClient()
}
