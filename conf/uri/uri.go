package uri

//to TN
const TrackLocationPathPrefixGFID = "tn/gfs/file/location/agfid/"

const TrackLocationList = "pn/addresses"
const TrackLocation = "pn/address"

//RnV1
const CreateGfidPath = "rn/gfs/gfid"
const UpdateGfidPath = "rn/gfs/param/"
const UpdateGfidPathSign = "rn/gfs/sign/param/"

//由于bfs接口已经屏蔽，调整获取系统版本接口路径
const GetVersionRnBfsPath = "rn/sys/version"

const RnSysTime = "rn/sys/time"

//RN Api
const (

	//上传

	RN_COMBINE_UPLOAD_URL = "rn/combine/file"

	RN_RFS_UPLOAD_URL="rn/rfs/file"

	RN_BFS_UPLOAD_URL="rn/bfs/file"

	// 下载
	RN_COMBINE_DOWNLOAD_URL = "rn/combine/file/%s"

	RN_BFS_DOWNLOAD_URL = "rn/bfs/file/%s"

	RN_RFS_DOWNLOAD_URL = "rn/rfs/file/%s"

	// 查raw
	RN_COMBINE_SEED_INDEX_URL = "rn/combine/seed/%s"

	// 创建index
	RN_COMBINE_INDEX_URL = "rn/combine/idx/%s/%s"

	RN_BFS_INDEX_URL = "rn/bfs/idx/%s/%s"

	RN_RFS_INDEX_URL = "rn/rfs/idx/%s/%s"

	// set params
	RN_SET_BFS_PARAMS_URL = "rn/bfs/file/%s/params"

	RN_SET_RFS_PARAMS_URL = "rn/rfs/file/%s/params"

	// read params
	RN_READ_RFS_PARAMS_URL = "rn/rfs/param/%s/%s"

	RN_READ_BFS_PARAMS_URL = "rn/bfs/param/%s/%s"

	// dgst
	RN_COMBINE_AFID2DGST_URL = "rn/combine/afid2dgst/%s/%s"

	RN_COMBINE_DGST2AFID_URL = "rn/combine/dgst2afid/%s/%s"

	// 改过期时间
	RN_BFS_PARAMS_EXPIRED_URL = "rn/bfs/file/%s/params/expired"

	RN_RFS_PARAMS_EXPIRED_URL = "rn/rfs/file/%s/params/expired"


//TN Api
	TN_RNODESINFO_URL = "tn/static/rnCluster"

	TN_GFS_LOCATION_BY_AGFID_URL = "tn/gfs/file/location/agfid/{agfid}"

	TN_FILE_LOCATION_QUERYTYPE_URL = "tn/file/location/%s/%s"

	TN_LOCATION_SEEDID_BY_DGST_URL = "tn/location/seedid/%s"

	TN_QUERY_DGST_URL = "tn/query/%s"

	TN_QUERY_PNADDRESS_URL = "pn/addresses"

	TN_SYS_HEALTHY_URL = "tn/sys/healthy"

	TN_DEL_CACHE_BYDGST_URL = "tn/cache/dgst/%s"
)


