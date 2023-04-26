package utils

import (
	"chainstorage-sdk/common/e/code"
	"chainstorage-sdk/conf"
	model "chainstorage-sdk/vo"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	"net"
	"net/http"
	"time"
)

/*
*
参数为json
*/
func Post(url string, param interface{}) (scode, httpcode int, data interface{}) {
	response, e := resty.SetTimeout(time.Duration(conf.GatewayTimeout)*time.Second).R().
		SetHeader("Connection", "keep-alive").
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		Post(url)
	var respBody model.GeneralResp
	json.Unmarshal(response.Body(), &respBody)
	if e != nil {
		if perr, ok := e.(net.Error); ok && perr.Timeout() {
			DevLog("POST", fmt.Sprintf("POST to %s  timed out", url))
			return code.ScodeTnErr, http.StatusBadGateway, &map[string]interface{}{}
		}
		LogError(e.Error())
		return code.ScodeErr, http.StatusInternalServerError, nil
	} else {
		if respBody.Code != code.ScodeOK {
			return code.ScodeRnErr, http.StatusOK, respBody.Data
		}
		return code.ScodeOK, http.StatusOK, respBody.Data
	}

}

/*
*
参数为json
*/
func Put(url string, param interface{}) (scode, httpcode int, data interface{}) {
	response, e := resty.SetTimeout(time.Duration(conf.GatewayTimeout)*time.Second).R().
		SetHeader("Connection", "keep-alive").
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		Put(url)
	var respBody model.GeneralResp
	json.Unmarshal(response.Body(), &respBody)
	if e != nil {
		if perr, ok := e.(net.Error); ok && perr.Timeout() {
			DevLog("PUT", fmt.Sprintf("PUT to %s  timed out", url))
			return code.ScodeTnErr, http.StatusBadGateway, &map[string]interface{}{}
		}
		LogError(e.Error())
		return code.ScodeErr, http.StatusInternalServerError, nil
	} else {
		if respBody.Code != code.ScodeOK {
			return code.ScodeRnErr, http.StatusOK, respBody.Data
		}
		return code.ScodeOK, http.StatusOK, respBody.Data
	}

}
func Get(url string) (scode, httpcode int, data interface{}) {
	response, e := resty.SetTimeout(time.Duration(conf.GatewayTimeout)*time.Second).R().
		SetHeader("Connection", "keep-alive").
		Get(url)
	var respBody model.GeneralResp
	if e != nil {
		if perr, ok := e.(net.Error); ok && perr.Timeout() {
			DevLog("GET", fmt.Sprintf("GET to %s  timed out", url))
			return code.ScodeTnErr, http.StatusBadGateway, &map[string]interface{}{}
		}
		return code.ScodeErr, http.StatusInternalServerError, nil
	} else {
		json.Unmarshal(response.Body(), &respBody)
		if respBody.Code != code.ScodeOK {
			return code.ScodeRnErr, http.StatusOK, respBody.Data
		}
		return code.ScodeOK, http.StatusOK, respBody.Data
	}

}
