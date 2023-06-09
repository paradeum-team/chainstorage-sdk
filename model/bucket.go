package model

import (
	"time"
)

type Bucket struct {
	Id                  int       `json:"id" comment:"桶ID"`
	UserId              int       `json:"-" comment:"用户ID"`
	BucketName          string    `json:"bucketName" comment:"桶名称（3-63字长度限制）"`
	StorageNetworkCode  int       `json:"storageNetworkCode" comment:"存储网络编码（10001-IPFS）"`
	BucketPrincipleCode int       `json:"bucketPrincipleCode" comment:"桶策略编码（10001-公开，10000-私有）"`
	UsedSpace           int64     `json:"usedSpace" comment:"已使用空间（字节）"`
	ObjectAmount        int       `json:"objectAmount" comment:"对象数量"`
	Status              int       `json:"status" comment:"记录状态（0-有效，1-删除）"`
	CreatedAt           time.Time `json:"createdAt" comment:"创建时间"`
	UpdatedAt           time.Time `json:"updatedAt" comment:"最后更新时间"`
}

type BucketPageResponse struct {
	RequestId string     `json:"requestId,omitempty"`
	Code      int32      `json:"code,omitempty"`
	Msg       string     `json:"msg,omitempty"`
	Status    string     `json:"status,omitempty"`
	Data      BucketPage `json:"data,omitempty"`
}

type BucketPage struct {
	Count     int      `json:"count,omitempty"`
	PageIndex int      `json:"pageIndex,omitempty"`
	PageSize  int      `json:"pageSize,omitempty"`
	List      []Bucket `json:"list,omitempty"`
}

type BucketCreateResponse struct {
	RequestId string `json:"requestId,omitempty"`
	Code      int32  `json:"code,omitempty"`
	Msg       string `json:"msg,omitempty"`
	Status    string `json:"status,omitempty"`
	Data      Bucket `json:"data,omitempty"`
}

type BucketEmptyResponse struct {
	RequestId string      `json:"requestId,omitempty"`
	Code      int32       `json:"code,omitempty"`
	Msg       string      `json:"msg,omitempty"`
	Status    string      `json:"status,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

type BucketRemoveResponse struct {
	RequestId string      `json:"requestId,omitempty"`
	Code      int32       `json:"code,omitempty"`
	Msg       string      `json:"msg,omitempty"`
	Status    string      `json:"status,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}
