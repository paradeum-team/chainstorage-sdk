package afs

import (
	"chainstorage-sdk/common/dict"
	"chainstorage-sdk/common/e/code"
	"chainstorage-sdk/common/plogger"

	"chainstorage-sdk/entity"
	"chainstorage-sdk/utils"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
)

/*
*
处理TN寻址结果
*/
func handleQueryDgstResponse(dgstType, dgstStr, nodeType string, response entity.ResponseVO) (httpcode int, dgstResp *entity.RespQueryDgst) {

	dgstResp = &entity.RespQueryDgst{}
	if response.HttpCode != http.StatusOK || response.Code != code.ScodeOK {
		dgstResp.Scode = response.Code
		dgstResp.TrackMsg = response.Msg
		dgstResp.Err = errors.New(response.Msg)
		return response.HttpCode, dgstResp
	}
	resp := entity.TNQueryDgstResp{}
	responseData := response.Data
	if response.Data != nil {
		bytes, e := json.Marshal(responseData)
		if e != nil {
			dgstResp.Scode = code.ScodeErr
			dgstResp.TrackMsg = "QueryDgstLocation respBody to json is error "
			dgstResp.Err = e
			utils.DevLog("queryTrackNode", fmt.Sprintf("parse track response failed %v", e))
			utils.LogError(fmt.Sprintf("cannot parse response from track node querying with dgst %s:%s", dgstType, dgstStr))
			return http.StatusInternalServerError, dgstResp
		}
		json.Unmarshal(bytes, &resp)
		if resp.Afid == "" {
			dgstResp.TrackMsg = fmt.Sprintf("cannot get afid from track node querying with %s:%s", dgstType, dgstStr)
			utils.LogError(dgstResp.TrackMsg)
			dgstResp.Scode = code.ScodeRNodesInfoErr
			dgstResp.TrackMsg = "file not found"
			dgstResp.Err = fmt.Errorf(code.ScodeMsgMapping[code.ScodeRNodesInfoErr])
			return http.StatusNotFound, dgstResp
		}
		//拼装TN Response对象参数
		dgstResp = assemblyDgstResp(nodeType, &resp)
		if dgstResp.Err == nil {
			dgstResp.Scode = response.Code
			return response.HttpCode, dgstResp
		} else {
			dgstResp.Scode = code.ScodeNoFNDat
			return http.StatusNotFound, dgstResp
		}

	} else {
		dgstResp.TrackMsg = fmt.Sprintf("cannot get afid from track node querying with %s:%s", dgstType, dgstStr)
		utils.LogError(dgstResp.TrackMsg)
		dgstResp.Scode = code.ScodeRNodesInfoErr
		dgstResp.TrackMsg = "file not found"
		dgstResp.Err = fmt.Errorf(code.ScodeMsgMapping[code.ScodeRNodesInfoErr])
		return http.StatusNotFound, dgstResp
	}
}

/*
*
拼装dgstResp
*/
func assemblyDgstResp(nodeType string, tnQueryDgstResp *entity.TNQueryDgstResp) (dgstResp *entity.RespQueryDgst) {
	dgstResp = &entity.RespQueryDgst{}
	dgstResp.Afid = tnQueryDgstResp.Afid
	dgstResp.IsBigFile = tnQueryDgstResp.IsBigFile
	dgstResp.LargeFileAfid = tnQueryDgstResp.LargeFileAfid
	//判断文件在哪些node上，区分不同的node类型
	rNodes := tnQueryDgstResp.RNodes
	dgstResp = assemblyRandomRnodes(nodeType, rNodes, dgstResp)
	if dgstResp.Err == nil {
		dgstResp.IsExist = tnQueryDgstResp.IsExist
	} else {
		dgstResp.IsExist = false
	}
	return dgstResp
}

/*
*
返回的是随机RN
*/
func assemblyRandomRnodes(nodeType string, rNodes []entity.RnInfo, dgstResp *entity.RespQueryDgst) *entity.RespQueryDgst {
	var rfsNode []string
	var afsNode []string
	var agfsNode []string
	var unknowNode []string
	dgstResp.NodeType = map[string]bool{}
	for _, rnode := range rNodes {
		isArfs := rnode.IsInArfs
		isAfs := rnode.IsInAfs
		isAgfs := rnode.IsInAgfs
		if isArfs {
			rfsNode = append(rfsNode, rnode.Api)
			if isAfs {
				afsNode = append(afsNode, rnode.Api)
			}
			continue
		} else if isAfs {
			afsNode = append(afsNode, rnode.Api)
			continue
		} else if isAgfs {
			agfsNode = append(agfsNode, rnode.Api)
			continue
		} else {
			unknowNode = append(unknowNode, rnode.Api)
			continue
		}
	}

	//判断agfs
	if len(agfsNode) > 0 {
		dgstResp.NodeType[dict.NODE_TYPE_GFS] = true
		dgstResp.RnodeRoots = agfsNode
	} else if nodeType == dict.NODE_TYPE_RFS && len(rfsNode) > 0 {
		dgstResp.RnodeRoots = rfsNode
		dgstResp.NodeType[dict.NODE_TYPE_RFS] = true
	} else if nodeType == dict.NODE_TYPE_BFS && len(afsNode) > 0 {
		dgstResp.NodeType[dict.NODE_TYPE_BFS] = true
		dgstResp.RnodeRoots = afsNode
	} else {
		dgstResp.RnodeRoots = unknowNode
	}

	if len(dgstResp.RnodeRoots) > 0 {
		dgstResp.RnodeRootsTmp = dgstResp.RnodeRoots
		randomRn := []string{}
		randomIndex := rand.Intn(len(dgstResp.RnodeRoots))
		random := dgstResp.RnodeRoots[randomIndex]
		randomRn = append(randomRn, random)
		dgstResp.RnodeRoots = randomRn
		utils.DevLog("getRNodeRoots", fmt.Sprintf("successfully got %d len rnodes", len(dgstResp.RnodeRootsTmp)))
		return dgstResp

	}
	dgstResp.Err = errors.New(fmt.Sprintf("File does not exist on %s", nodeType))
	plogger.NewInstance().GetLogger().Errorf("File does not exist on %s", nodeType)
	return dgstResp

	/*if nodeType == ""{
		if len(rfsNode) >0{
			dgstResp.NodeType[dict.NODE_TYPE_RFS] = true
			dgstResp.RnodeRoots = rfsNode
		}else if len(afsNode)>0 {
			dgstResp.NodeType[dict.NODE_TYPE_BFS] = true
			dgstResp.RnodeRoots = afsNode
		}else if len(agfsNode)>0 {
			dgstResp.NodeType[dict.NODE_TYPE_GFS] = true
			dgstResp.RnodeRoots = agfsNode
		} else {
			dgstResp.RnodeRoots = unknowNode
		}
		if len(dgstResp.RnodeRoots)>0{
			dgstResp.RnodeRootsTmp = dgstResp.RnodeRoots

			randomRn := []string{}
			randomIndex := rand.Intn(len(dgstResp.RnodeRoots))
			random := dgstResp.RnodeRoots[randomIndex]
			randomRn = append(randomRn,random)
			dgstResp.RnodeRoots = randomRn
		}
		utils.DevLog("getRNodeRoots", fmt.Sprintf("successfully got %d len rnodes", len(dgstResp.RnodeRootsTmp)))
		return dgstResp

	}else {


	}*/
}
