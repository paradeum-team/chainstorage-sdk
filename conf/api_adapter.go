package conf

import (
	"chainstorage-sdk/common/dict"
	"chainstorage-sdk/conf/uri"
)

// 上传、下载、创建index、set params、 read params、 改过期时间

var RnUploadUrl string
var RnDownloadUrl string
var RnCreateIndexUrl string
var RnReadParamsUrl string
var RnWriteParamsUrl string
var RnModifyExpiredUrl string

/*
*
初始化根据bfs域，选择对应api
*/
func apiAdapter() {

	switch UploadFieldModel {
	case dict.NODE_TYPE_RFS:
		RnUploadUrl = uri.RN_RFS_UPLOAD_URL
	case dict.NODE_TYPE_BFS:
		RnUploadUrl = uri.RN_BFS_UPLOAD_URL
	default:
		RnUploadUrl = uri.RN_COMBINE_UPLOAD_URL
	}

	switch Config.UseField {
	case dict.NODE_TYPE_RFS:
		RnDownloadUrl = uri.RN_RFS_DOWNLOAD_URL
		RnCreateIndexUrl = uri.RN_RFS_INDEX_URL
		RnReadParamsUrl = uri.RN_READ_RFS_PARAMS_URL
		RnWriteParamsUrl = uri.RN_SET_RFS_PARAMS_URL
		RnModifyExpiredUrl = uri.RN_RFS_PARAMS_EXPIRED_URL
	case dict.NODE_TYPE_BFS:
		RnDownloadUrl = uri.RN_BFS_DOWNLOAD_URL
		RnCreateIndexUrl = uri.RN_BFS_INDEX_URL
		RnReadParamsUrl = uri.RN_READ_BFS_PARAMS_URL
		RnWriteParamsUrl = uri.RN_SET_BFS_PARAMS_URL
		RnModifyExpiredUrl = uri.RN_BFS_PARAMS_EXPIRED_URL
	default:
		RnDownloadUrl = uri.RN_COMBINE_DOWNLOAD_URL
		RnCreateIndexUrl = uri.RN_COMBINE_INDEX_URL
		RnReadParamsUrl = uri.RN_READ_BFS_PARAMS_URL
		RnWriteParamsUrl = uri.RN_SET_BFS_PARAMS_URL
		RnModifyExpiredUrl = uri.RN_BFS_PARAMS_EXPIRED_URL
	}
}
