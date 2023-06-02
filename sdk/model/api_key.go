package model

import "time"

type ApiKey struct {
	Id                            int       `json:"id" comment:"ApiKeyID"`
	UserId                        int       `json:"-" comment:"用户ID"`
	ApiName                       string    `json:"apiName" comment:"Api名称（3-63字长度限制）"`
	ApiKey                        string    `json:"apiKey" comment:"comment:ApiKey（默认20字节）"`
	ApiSecret                     string    `json:"apiSecret" comment:"Api密钥（默认40字节）"`
	PermissionTypeCode            int       `json:"permissionTypeCode" comment:"权限类型编码（管理员设置）（10001-管理员权限，10000-自定义设置）"`
	PermissionIdList              []int     `json:"permissionIdList" comment:"API服务权限ID列表"`
	PinninServicePermissionIdList []int     `json:"pinninServicePermissionIdList" comment:"Pinning服务权限ID列表"`
	DataScope                     int       `json:"dataScope" comment:"数据范围(桶ID)"`
	Status                        int       `json:"status" comment:"记录状态（0-有效，1-删除）"`
	CreatedAt                     time.Time `json:"createdAt" comment:"创建时间"`
	UpdatedAt                     time.Time `json:"updatedAt" comment:"最后更新时间"`
}

type ApiKeyPageResponse struct {
	RequestId string     `json:"requestId,omitempty"`
	Code      int32      `json:"code,omitempty"`
	Msg       string     `json:"msg,omitempty"`
	Status    string     `json:"status,omitempty"`
	Data      ApiKeyPage `json:"data,omitempty"`
}

type ApiKeyPage struct {
	Count     int      `json:"count,omitempty"`
	PageIndex int      `json:"pageIndex,omitempty"`
	PageSize  int      `json:"pageSize,omitempty"`
	List      []ApiKey `json:"list,omitempty"`
}
