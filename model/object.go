package model

import (
	"time"
)

type Object struct {
	Id                      int                    `json:"id" comment:"对象ID"`
	UserId                  int                    `json:"userId" comment:"用户ID"`
	BucketId                int                    `json:"bucketId" comment:"桶主键"`
	ObjectName              string                 `json:"objectName" comment:"对象名称（255字限制）"`
	ObjectTypeCode          int                    `json:"objectTypeCode" comment:"对象类型编码"`
	ObjectSize              int64                  `json:"objectSize" comment:"对象大小（字节）"`
	IsMarked                int                    `json:"isMarked" comment:"星标（1-已标记，0-未标记）"`
	ObjectCid               string                 `json:"objectCid" comment:"对象CID"`
	Status                  int                    `json:"status" comment:"记录状态（0-有效，1-删除）"`
	LinkedStorageObjectCode string                 `json:"linkedStorageObjectCode" comment:"链存类型对象编码"`
	LinkedStorageObject     map[string]interface{} `json:"linkedStorageObject" comment:"链存类型对象"`
	CreatedAt               time.Time              `json:"createdAt" comment:"创建时间"`
	UpdatedAt               time.Time              `json:"updatedAt" comment:"最后更新时间"`
}

type ObjectPageResponse struct {
	RequestId string     `json:"requestId,omitempty"`
	Code      int32      `json:"code,omitempty"`
	Msg       string     `json:"msg,omitempty"`
	Status    string     `json:"status,omitempty"`
	Data      ObjectPage `json:"data,omitempty"`
}

type ObjectPage struct {
	Count     int      `json:"count,omitempty"`
	PageIndex int      `json:"pageIndex,omitempty"`
	PageSize  int      `json:"pageSize,omitempty"`
	List      []Object `json:"list,omitempty"`
}

type ObjectCreateResponse struct {
	RequestId string `json:"requestId,omitempty"`
	Code      int32  `json:"code,omitempty"`
	Msg       string `json:"msg,omitempty"`
	Status    string `json:"status,omitempty"`
	Data      Object `json:"data,omitempty"`
}

type ObjectRemoveResponse struct {
	RequestId string      `json:"requestId,omitempty"`
	Code      int32       `json:"code,omitempty"`
	Msg       string      `json:"msg,omitempty"`
	Status    string      `json:"status,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

type ObjectRenameResponse struct {
	RequestId string      `json:"requestId,omitempty"`
	Code      int32       `json:"code,omitempty"`
	Msg       string      `json:"msg,omitempty"`
	Status    string      `json:"status,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

type ObjectMarkResponse struct {
	RequestId string      `json:"requestId,omitempty"`
	Code      int32       `json:"code,omitempty"`
	Msg       string      `json:"msg,omitempty"`
	Status    string      `json:"status,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}
