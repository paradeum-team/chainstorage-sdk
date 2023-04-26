package ipfs

import (
	client "chainstorage-sdk/base/basicauthclient"
	"chainstorage-sdk/common/e/code"
	"chainstorage-sdk/conf"
	"chainstorage-sdk/utils"
	model "chainstorage-sdk/vo"
	"encoding/json"
	"net/http"
	"time"
)

var (
	// ifps cluster地址
	ipfsClusterBaseAddress = conf.Config.IpfsClusterAddresses[0]
	// ifps 地址
	ipfsBaseAddress = conf.Config.IpfsAddresses[0]
)

/**
 * 根据摘要值获取cid列表
 */
func GetCidsByDgst(dgst string) (int, int, model.CidInfo, error) {
	defer utils.DevTimeTrack(time.Now(), "GetCidsByDgst")

	cidInfo := model.CidInfo{}
	url := ipfsClusterBaseAddress + "bfs/GetCidInfoByDgst/" + dgst
	httpStatus, body, err := client.RestyGet(url)
	if err != nil {
		return code.ScodeErr, http.StatusInternalServerError, cidInfo, err
	}

	err = json.Unmarshal(body, &cidInfo)
	if err != nil {
		return code.ScodeErr, http.StatusInternalServerError, cidInfo, err
	}
	utils.DevLog("GetCidsByDgst", string(body))

	return code.ScodeOK, httpStatus, cidInfo, nil
}
