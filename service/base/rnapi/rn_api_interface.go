package rnapi

import (
	"chainstorage-sdk/entity"
	"chainstorage-sdk/vo"
)

type RnApiService interface {
	RnCombineFileDownload(paramVO *vo.DownloadParamsVO, dgstResp *entity.RespQueryDgst) entity.ResponseVO

	RnCombineFileUpload(paramsVO *vo.UploadParamsVO)

	RnSimpleDownload(dgstResp *entity.RespQueryDgst) entity.ResponseVO

	RnReadParams(url string) *entity.ResponseVO
}

type rnApiService struct{}

var rnApiServiceInstance *rnApiService

func NewRNApiService() RnApiService {
	if rnApiServiceInstance == nil {
		rnApiServiceInstance = &rnApiService{}
	}
	return rnApiServiceInstance
}
