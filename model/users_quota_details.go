package model

import (
	"time"
)

type UsersQuotaDetails struct {
	Id          int `json:"id" comment:"ID"`
	UserQuotaId int `json:"userQuotaId" comment:"用户资源限额主表Id"`
	//UserId           int       `json:"userId" comment:"用户Id"`
	Sequence       int    `json:"sequence" comment:"显示顺序"`
	NetCode        int    `json:"netCode" comment:"网络类型code"`
	NetName        string `json:"netName" comment:"网络类型名称"`
	ConstraintId   int    `json:"constraintId" comment:"限额资源Id"`
	ConstraintName string `json:"constraintName" comment:"限额资源名称（cn,冗余）"`
	//ConstraintNameCN string    `json:"constraintNameCN" comment:"资源名称"`
	//ConstraintNameEN string    `json:"constraintNameEN" comment:"资源名称"`
	LimitedQuota int64     `json:"limitedQuota" comment:"最大限额"`
	UsedQuota    int64     `json:"usedQuota" comment:"已经使用量"`
	Available    int64     `json:"available" comment:"可用量(余量，冗余)"`
	CreatedAt    time.Time `json:"createdAt" comment:"创建时间"`
	UpdatedAt    time.Time `json:"updatedAt" comment:"最后更新时间"`
}
