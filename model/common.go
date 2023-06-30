package model

type VersionResponse struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data Version `json:"data"`
}
type Version struct {
	Code    int    `json:"code"`
	Version string `json:"version"`
}

// 按照网络类型所有桶容量统计响应参数
type BucketStorageTypeStatResp struct {
	StorageNetworkCode int   `json:"storageNetworkCode" comment:"存储网络编码（10001-IPFS）"`
	UsedSpace          int64 `json:"usedSpace" comment:"已使用空间（字节）"`
	UsedSpaceQuota     int64 `json:"usedSpaceQuota" comment:"用户空间上限"`
	ObjectAmount       int   `json:"objectAmount" comment:"对象数量"`
	ObjectAmountQuota  int   `json:"objectAmountQuota" comment:"IPFS总个数上限"`
}
