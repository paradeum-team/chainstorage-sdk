package tnapi

import "chainstorage-sdk/entity"

type TnApiService interface {
	CheckSysHealthy(trackRoot string) (response entity.RnClusterResponse)

	GetRnCluster(trcakRoot string) (response entity.RnClusterResponse)

	TnQueryLocationByDgst(trackRoot, dgst string) (response entity.ResponseVO)

	TnQueryBigfileLocationByDgst(trackRoot, dgst string) (response entity.ResponseVO)

	TnFileQueryLocationDgst2Afid(trackRoot, dgstType, dgst string) (response entity.ResponseVO)
}

type tnApiService struct {
}

var tnApiServiceInstace *tnApiService

func NewTnApiServiceInstace() TnApiService {
	if tnApiServiceInstace == nil {
		tnApiServiceInstace = &tnApiService{}
	}
	return tnApiServiceInstace
}
