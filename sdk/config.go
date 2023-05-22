package chainstoragesdk

import (
	"fmt"
	"github.com/flyleft/gprofile"
	"os"
	"strings"
)

var appConfig ApplicationConfig
var cssConfig Configuration
var cssLoggerConfig LoggerConf

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
	// 链存服务API地址
	ChainStorageApiBaseAddress string `profile:"chainStorageApiBaseAddress" profileDefault:"http://127.0.0.1:8821" json:"chainStorageApiBaseAddress"`

	// CAR文件工作目录
	CarFileGenerationPath string `profile:"carFileGenerationPath" profileDefault:"./temp/carfile" json:"carFileGenerationPath"`

	// CAR文件分片阈值
	CarFileShardingThreshold int `profile:"carFileShardingThreshold" profileDefault:"1048576" json:"carFileShardingThreshold"`

	// 链存服务API token
	ChainStorageApiToken string `profile:"chainStorageApiToken" profileDefault:"" json:"chainStorageApiToken"`

	// HTTP request user agent (K2请求需要)
	HttpRequestUserAgent string `profile:"httpRequestUserAgent" profileDefault:"" json:"httpRequestUserAgent"`
}

type LoggerConf struct {
	LogPath      string `profile:"logPath" profileDefault:"./logs" json:"logPath"`
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
	//if len(configFile) == 0 {
	//	configFile = "./github.com/paradeum-team/chainstorage-sdk.yaml"
	//}

	configFile := "./chainstorage-sdk.yaml"
	config, err := gprofile.Profile(&ApplicationConfig{}, configFile, true)
	if err != nil {
		fmt.Errorf("Profile execute error", err)
	}

	appConfig = *config.(*ApplicationConfig)
	cssConfig = config.(*ApplicationConfig).Server
	cssLoggerConfig = config.(*ApplicationConfig).Logger

	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	log.Fatal(err)
	//	os.Exit(1)
	//}

	//check chain-storage-api base address
	if len(cssConfig.ChainStorageApiBaseAddress) > 0 {
		chainStorageApiBaseAddress := cssConfig.ChainStorageApiBaseAddress
		if !strings.HasPrefix(chainStorageApiBaseAddress, "http") {
			fmt.Println("ERROR: invalid chain-storage-api base address in Configuration file, chain-storage-api base address must be a valid http/https url, exiting")
			os.Exit(1)
		}

		if !strings.HasSuffix(chainStorageApiBaseAddress, "/") {
			cssConfig.ChainStorageApiBaseAddress += "/"
		}
	} else {
		fmt.Println("ERROR: no chain-storage-api base address provided in Configuration file, at least 1 valid http/https chain-storage-api base address must be given, exiting")
		os.Exit(1)
	}

	if len(cssConfig.ChainStorageApiToken) == 0 {
		fmt.Println("ERROR: invalid chain-storage-api token in Configuration file, chain-storage-api token must not be empty")
		os.Exit(1)
	} else if !strings.HasPrefix(cssConfig.ChainStorageApiToken, "Bearer ") {
		cssConfig.ChainStorageApiToken = "Bearer " + cssConfig.ChainStorageApiToken
	}

	//if _, err := os.Stat(filepath.Join(Config.AbsAFSDir, Config.AFSProgram)); os.IsExist(err) {
	//	IsAFSExist = false
	//} else if err != nil {
	//	IsAFSExist = false
	//}
	////check port
	//if Config.Port > 1023 && Config.Port <= 65535 {
	//	Port = Config.Port
	//} else {
	//	fmt.Printf("WARNING: invalid port number(port) in Configuration file, using default:%s\n", Port)
	//}

}

func InitConfig2() {
	//rand.Seed(time.Now().UnixNano())
	config, err := gprofile.Profile(&ApplicationConfig{}, "./chainstorage-sdk.yaml", true)
	if err != nil {
		fmt.Errorf("Profile execute error", err)
	}
	appConfig = *config.(*ApplicationConfig)
	cssConfig = config.(*ApplicationConfig).Server
	cssLoggerConfig = config.(*ApplicationConfig).Logger

	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	log.Fatal(err)
	//	os.Exit(1)
	//}

	//check chain-storage-api base address
	if len(cssConfig.ChainStorageApiBaseAddress) > 0 {
		chainStorageApiBaseAddress := cssConfig.ChainStorageApiBaseAddress
		if !strings.HasPrefix(chainStorageApiBaseAddress, "http") {
			fmt.Println("ERROR: invalid chain-storage-api base address in Configuration file, chain-storage-api base address must be a valid http/https url, exiting")
			os.Exit(1)
		}

		if !strings.HasSuffix(chainStorageApiBaseAddress, "/") {
			cssConfig.ChainStorageApiBaseAddress += "/"
		}
	} else {
		fmt.Println("ERROR: no chain-storage-api base address provided in Configuration file, at least 1 valid http/https chain-storage-api base address must be given, exiting")
		os.Exit(1)
	}
}
