package conf

var AppConfig ApplicationConfig
var Config Configuration
var Logger LoggerConf
var Leveldb LevelDb
var GatewayTimeout = 60
var Port int
var MaxExecTime = 300
var IsAFSExist bool
var DownloadDir string
var UploadDir string
var UploadFieldModel string

const DownloadFolder = "download/"
const UploadFolder = "upload/"

var ContentTypeMap = make(map[string]string, 20)

type Configuration struct {
	Port                 int      `profile:"port" profileDefault:"5145" json:"port"`
	PProfPort            int      `profile:"pprof-port" profileDefault:"9122" json:"pprof-port"`
	TrackRoots           []string `profile:"trackRoots" profileDefault:"["http://127.0.0.1:5143"]" json:"trackRoots"`
	IpfsClusterAddresses []string `profile:"ipfsClusterAddresses" profileDefault:"["http://127.0.0.1:9094"]" json:"ipfsClusterAddresses"`
	IpfsAddresses        []string `profile:"ipfsAddresses" profileDefault:"["http://127.0.0.1:8081"]" json:"ipfsAddresses"`
	GatewayTimeout       int      `profile:"gatewayTimeout" json:"gatewayTimeout"`
	MaxExecTime          int      `profile:"maxExecTime" json:"maxExecTime"`
	UploadFieldModel     string   `profile:"uploadFieldModel" profileDefault:"combine" json:"uploadFieldModel"`
	RNodesInfoPath       string   `profile:"_"`
	APIDataFolder        string   `profile:"_"`
	//判断本地是否存在AFS客户端，计算AFID
	AbsAFSDir  string `profile:"_"` // Absolute path of AFS core.
	AFSProgram string `profile:"_"` // Name of AFS core program.

	//限流
	RateLimitMax float64 `profile:"rateLimitMax" profileDefault:"1000" json:"rateLimitMax"`

	UploadLimitRate   float64 `profile:"uploadLimitRate" profileDefault:"1" json:"uploadLimitRate"`
	DownloadLimitRate float64 `profile:"downloadLimitRate" profileDefault:"1" json:"downloadLimitRate"`

	DataPath string `profile:"dataPath" profileDefault:"afs_data" json:"dataPath"`

	ContentTypeAddr string `profile:"contentTypeAddr",profileDefault:"" json:"contentTypeAddr"`

	BfsCopies int `profile:"bfsCopies" profileDefault:"3" json:"bfsCopies"`

	UploadLimitSize int64 `profile:"uploadLimitSize" profileDefault:"10" json:"uploadLimitSize"`

	UseField string `profile:"useField" profileDefault:"rfs" json:"useField"`

	PProf bool `profile:"pprof" profileDefault:"false" json:"pprof"`

	LinkedStorageApiBaseAddress string `profile:"linkedStorageApiBaseAddress" profileDefault:"http://127.0.0.1:8821" json:"linkedStorageApiBaseAddress"`
}

type LoggerConf struct {
	Mode         string `prfile:"mode" profileDefault:"release" json:"mode"`
	Level        string `prfile:"level" profileDefault:"info" json:"level"`
	IsOutPutFile bool   `profile:"isOutPutFile" profileDefault:"false" json:"isOutPutFile"`
	MaxAgeDay    int64  `profile:"maxAgeDay" profileDefault:"7" json:"maxAgeDay"`
	RotationTime int64  `profile:"rotationTime" profileDefault:"1" json:"rotationTime"`
}

type LevelDb struct {
	Openpath string `profile:"openpath" profileDefault: "./afs_data/datafile"`
}

const AfidLength = 128

var DevUglyLogCounter = 0 //devUglyLogCounter and devOp are tags for logging only.

type ApplicationConfig struct {
	Server  Configuration `profile:"server"`
	Logger  LoggerConf    `profile:"logger"`
	LevelDb LevelDb       `profile:"leveldb"`
}
