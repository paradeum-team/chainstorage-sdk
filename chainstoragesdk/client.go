package chainstoragesdk

import (
	"sync"
)

type SdkConfig Configuration

var once = sync.Once{}
var mClient *myClient

type myClient struct {
	Config     *Configuration
	Logger     *PldLogger
	httpClient *RestyClient
	Bucket     *Bucket
	Object     *Object
	Car        *Car
}

func newClient() (*myClient, error) {
	var err error
	once.Do(func() {
		initConfig()

		mClient = &myClient{}
		mClient.Config = &myConfig

		mClient.Logger = GetLogger(&myLogger)

		mClient.httpClient = &RestyClient{Config: mClient.Config}
		mClient.Bucket = &Bucket{Config: mClient.Config, Client: mClient.httpClient, logger: mClient.Logger.logger}
		mClient.Object = &Object{Config: mClient.Config, Client: mClient.httpClient, logger: mClient.Logger.logger}
		mClient.Car = &Car{Config: mClient.Config, Client: mClient.httpClient, logger: mClient.Logger.logger}
	})

	//mClient.Logger.logger.Error("client new.")
	return mClient, err
}

func New() (*myClient, error) {

	return newClient()
}
