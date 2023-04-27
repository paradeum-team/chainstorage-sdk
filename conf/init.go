package conf

import (
	"bufio"
	"chainstorage-sdk/base/client"
	"chainstorage-sdk/conf/uri"
	"encoding/json"
	"fmt"
	"github.com/flyleft/gprofile"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// test merge
func init() {
	rand.Seed(time.Now().UnixNano())
	config, err := gprofile.Profile(&ApplicationConfig{}, "./application.yaml", true)
	if err != nil {
		fmt.Errorf("Profile execute error", err)
	}
	AppConfig = *config.(*ApplicationConfig)
	Config = config.(*ApplicationConfig).Server
	Logger = config.(*ApplicationConfig).Logger
	Leveldb = config.(*ApplicationConfig).LevelDb
	Config.RNodesInfoPath = uri.TN_RNODESINFO_URL

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	//check uploadFolder and doanloadFolder exist
	afsDataFolder := filepath.Join(dir, Config.APIDataFolder)
	if _, err := os.Stat(afsDataFolder); os.IsNotExist(err) {
		os.Mkdir(afsDataFolder, os.ModePerm)
	} else if err != nil {
		log.Fatal(err)
	}
	UploadDir = filepath.Join(dir, Config.APIDataFolder, UploadFolder)
	if _, err := os.Stat(UploadDir); os.IsNotExist(err) {
		os.Mkdir(UploadDir, os.ModePerm)
	} else if err != nil {
		log.Fatal(err)
	}
	DownloadDir = filepath.Join(dir, Config.APIDataFolder, DownloadFolder)
	if _, err := os.Stat(DownloadDir); os.IsNotExist(err) {
		os.Mkdir(DownloadDir, os.ModePerm)
	} else if err != nil {
		log.Fatal(err)
	}

	MaxExecTime = Config.MaxExecTime

	//----check afs-----------
	IsAFSExist = true

	if _, err := os.Stat(filepath.Join(Config.AbsAFSDir, Config.AFSProgram)); os.IsExist(err) {
		IsAFSExist = false
	} else if err != nil {
		IsAFSExist = false
	}
	//check port
	if Config.Port > 1023 && Config.Port <= 65535 {
		Port = Config.Port
	} else {
		fmt.Printf("WARNING: invalid port number(port) in Configuration file, using default:%s\n", Port)
	}

	//check rnodeInfoPath
	if Config.RNodesInfoPath != "" {
		Config.RNodesInfoPath = strings.TrimPrefix(Config.RNodesInfoPath, "/")
		Config.RNodesInfoPath = strings.TrimSuffix(Config.RNodesInfoPath, "/")
	} else {
		fmt.Println("ERROR: invalid path for RNodes info json(rnodesInfoPath) in Configuration file, exiting")
		os.Exit(1)
	}

	//check TN
	if len(Config.TrackRoots) > 0 {
		for i, tnRoot := range Config.TrackRoots {
			if tnRoot != "" {
				if !strings.HasPrefix(tnRoot, "http") {
					fmt.Println("ERROR: invalid track node roots(trackRoots) in Configuration file, track node root must be a valid http/https url, exiting")
					os.Exit(1)
				}
				if !strings.HasSuffix(tnRoot, "/") {
					Config.TrackRoots[i] += "/"
				}
			}
		}
	} else {
		fmt.Println("ERROR: no track node roots(trackRoots) provided in Configuration file, at least 1 valid http/https track node root must be given, exiting")
		os.Exit(1)
	}

	//check ipfs cluster
	if len(Config.IpfsClusterAddresses) > 0 {
		for i, ipfsCluster := range Config.IpfsClusterAddresses {
			if ipfsCluster != "" {
				if !strings.HasPrefix(ipfsCluster, "http") {
					fmt.Println("ERROR: invalid ipfs cluster address in Configuration file, ipfs cluster address must be a valid http/https url, exiting")
					os.Exit(1)
				}

				if !strings.HasSuffix(ipfsCluster, "/") {
					Config.IpfsClusterAddresses[i] += "/"
				}
			}
		}
	} else {
		fmt.Println("ERROR: no ipfs cluster address provided in Configuration file, at least 1 valid http/https ipfs cluster address must be given, exiting")
		os.Exit(1)
	}

	//check ipfs
	if len(Config.IpfsAddresses) > 0 {
		for i, ipfs := range Config.IpfsAddresses {
			if ipfs != "" {
				if !strings.HasPrefix(ipfs, "http") {
					fmt.Println("ERROR: invalid ipfs address in Configuration file, ipfs address must be a valid http/https url, exiting")
					os.Exit(1)
				}

				if !strings.HasSuffix(ipfs, "/") {
					Config.IpfsClusterAddresses[i] += "/"
				}
			}
		}
	} else {
		fmt.Println("ERROR: no ipfs address provided in Configuration file, at least 1 valid http/https ipfs address must be given, exiting")
		os.Exit(1)
	}

	//check gatewayTimeout
	if Config.GatewayTimeout <= 0 || Config.GatewayTimeout > 300 {
		fmt.Printf("WARNING: invalid gateway timeout(gatewayTimeout) in Configuration file, using default:%d\n", GatewayTimeout)
	} else {
		GatewayTimeout = Config.GatewayTimeout
	}

	//check uploadBoath
	UploadFieldModel = Config.UploadFieldModel

	//check chain-storage-api base address
	if len(Config.ChainStorageApiBaseAddress) > 0 {
		chainStorageApiBaseAddress := Config.ChainStorageApiBaseAddress
		if !strings.HasPrefix(chainStorageApiBaseAddress, "http") {
			fmt.Println("ERROR: invalid chain-storage-api base address in Configuration file, chain-storage-api base address must be a valid http/https url, exiting")
			os.Exit(1)
		}

		if !strings.HasSuffix(chainStorageApiBaseAddress, "/") {
			Config.ChainStorageApiBaseAddress += "/"
		}
	} else {
		fmt.Println("ERROR: no chain-storage-api base address provided in Configuration file, at least 1 valid http/https chain-storage-api base address must be given, exiting")
		os.Exit(1)
	}

	//设置ContentType
	setUpContentTypeMap(Config)

	//api 适配
	apiAdapter()
}

func setUpContentTypeMap(server Configuration) {
	ctypeAddress := server.ContentTypeAddr
	if ctypeAddress != "" {
		httpStatus, body, err := client.RestyGet(ctypeAddress)
		if err == nil && httpStatus == http.StatusOK {
			if len(body) < 1<<20 {
				err = ioutil.WriteFile("content-type.json", body, 0777)
				if err == nil {
					err = json.Unmarshal(body, &ContentTypeMap)
					if err == nil {
						fmt.Println("init get content-type.json is success.")
						return
					}
				}
			} else {
				fmt.Printf("WARNING: content-type-init.json The configuration file is too large, more than 1m, unable to read...")
			}
		}
		fmt.Printf("ERROR: init get content-type.json is error. httpcode is %d, err is %v", httpStatus, err)
	}

	contentTypeMap := readContentTypeFile()
	if contentTypeMap == nil {
		fmt.Println("ERROR: Not find content-type-init.json")
		os.Exit(1)
	}
	ContentTypeMap = contentTypeMap
	return
}

func readContentTypeFile() map[string]string {
	contentTypeMap := make(map[string]string, 20)
	file, err := os.Open("./content-type-init.json")
	if err != nil {
		return nil
	}
	defer file.Close()

	var bytes []byte
	inputReader := bufio.NewReader(file)
	if inputReader.Size() > 1<<20 {
		fmt.Printf("WARNING: content-type-init.json The configuration file is too large, more than 1m, unable to read...")
		return nil
	}
	for {
		inputByte, readerError := inputReader.ReadByte()
		if readerError == io.EOF {
			break
		}
		bytes = append(bytes, inputByte)
	}
	json.Unmarshal(bytes, &contentTypeMap)
	return contentTypeMap

}
