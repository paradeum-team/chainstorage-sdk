package chainstoragesdk

import (
	"fmt"
	"github.com/flyleft/gprofile"
	"os"
	"strings"
)

var appConfig ApplicationConfig
var myConfig Configuration
var Logger LoggerConf

//var GatewayTimeout = 60
//var Port int
//var MaxExecTime = 300
//var IsAFSExist bool
//var DownloadDir string
//var UploadDir string
//var UploadFieldModel string
//
//const DownloadFolder = "download/"
//const UploadFolder = "upload/"

//var ContentTypeMap = make(map[string]string, 20)

type Configuration struct {
	ChainStorageApiBaseAddress string `profile:"chainStorageApiBaseAddress" profileDefault:"http://127.0.0.1:8821" json:"chainStorageApiBaseAddress"`

	CarFileGenerationPath string `profile:"carFileGenerationPath" profileDefault:"./temp/carfile" json:"carFileGenerationPath"`

	CarFileShardingThreshold int `profile:"carFileShardingThreshold" profileDefault:"1048576" json:"carFileShardingThreshold"`
}

type LoggerConf struct {
	Mode         string `prfile:"mode" profileDefault:"release" json:"mode"`
	Level        string `prfile:"level" profileDefault:"info" json:"level"`
	IsOutPutFile bool   `profile:"isOutPutFile" profileDefault:"false" json:"isOutPutFile"`
	MaxAgeDay    int64  `profile:"maxAgeDay" profileDefault:"7" json:"maxAgeDay"`
	RotationTime int64  `profile:"rotationTime" profileDefault:"1" json:"rotationTime"`
}

type ApplicationConfig struct {
	Server Configuration `profile:"server"`
	Logger LoggerConf    `profile:"logger"`
}

func initConfig() {
	//rand.Seed(time.Now().UnixNano())
	config, err := gprofile.Profile(&ApplicationConfig{}, "./chainstorage-sdk.yaml", true)
	if err != nil {
		fmt.Errorf("Profile execute error", err)
	}
	appConfig = *config.(*ApplicationConfig)
	myConfig = config.(*ApplicationConfig).Server
	Logger = config.(*ApplicationConfig).Logger

	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	log.Fatal(err)
	//	os.Exit(1)
	//}

	//check chain-storage-api base address
	if len(myConfig.ChainStorageApiBaseAddress) > 0 {
		chainStorageApiBaseAddress := myConfig.ChainStorageApiBaseAddress
		if !strings.HasPrefix(chainStorageApiBaseAddress, "http") {
			fmt.Println("ERROR: invalid chain-storage-api base address in Configuration file, chain-storage-api base address must be a valid http/https url, exiting")
			os.Exit(1)
		}

		if !strings.HasSuffix(chainStorageApiBaseAddress, "/") {
			myConfig.ChainStorageApiBaseAddress += "/"
		}
	} else {
		fmt.Println("ERROR: no chain-storage-api base address provided in Configuration file, at least 1 valid http/https chain-storage-api base address must be given, exiting")
		os.Exit(1)
	}
}

func InitConfig2() {
	//rand.Seed(time.Now().UnixNano())
	config, err := gprofile.Profile(&ApplicationConfig{}, "./chainstorage-sdk.yaml", true)
	if err != nil {
		fmt.Errorf("Profile execute error", err)
	}
	appConfig = *config.(*ApplicationConfig)
	myConfig = config.(*ApplicationConfig).Server
	Logger = config.(*ApplicationConfig).Logger

	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	log.Fatal(err)
	//	os.Exit(1)
	//}

	//check chain-storage-api base address
	if len(myConfig.ChainStorageApiBaseAddress) > 0 {
		chainStorageApiBaseAddress := myConfig.ChainStorageApiBaseAddress
		if !strings.HasPrefix(chainStorageApiBaseAddress, "http") {
			fmt.Println("ERROR: invalid chain-storage-api base address in Configuration file, chain-storage-api base address must be a valid http/https url, exiting")
			os.Exit(1)
		}

		if !strings.HasSuffix(chainStorageApiBaseAddress, "/") {
			myConfig.ChainStorageApiBaseAddress += "/"
		}
	} else {
		fmt.Println("ERROR: no chain-storage-api base address provided in Configuration file, at least 1 valid http/https chain-storage-api base address must be given, exiting")
		os.Exit(1)
	}
}
