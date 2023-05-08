package chainstoragesdk

type SdkConfig Configuration

type myClient struct {
	Config     *Configuration
	httpClient *RestyClient
	Bucket     *Bucket
}

func New() (*myClient, error) {
	initConfig()

	client := &myClient{}
	client.Config = &myConfig
	client.httpClient = &RestyClient{Config: client.Config}

	client.Bucket = &Bucket{Config: client.Config, Client: client.httpClient}

	return client, nil
}
