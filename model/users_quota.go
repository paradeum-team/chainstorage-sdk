package model

import (
	"time"
)

type UsersQuota struct {
	Id int `json:"id" comment:"ID"`
	//UserId           int                 `json:"userId" comment:"用户Id"`
	NetCode       int    `json:"netCode" comment:"网络类型code"`
	NetName       string `json:"netName" comment:"网络类型名称"`
	PackagePlanId int    `json:"packagePlanId" comment:"套餐Id"`
	//PackagePlanName  string              `json:"packagePlanName" comment:"套餐名称"`
	//ServiceModelCode int                 `json:"serviceModelCode" comment:"服务类型(试用/购买)"`
	//ServiceModelName string              `json:"serviceModelName" comment:"服务类型名称(1=免费/2=试用/3=购买)"`
	OrderId int `json:"orderId" comment:"订单id"`
	//TransactionId int                 `json:"transactionId" comment:"交易id"`
	StartTime   time.Time           `json:"startTime" comment:"生效时间"`
	ExpiredTime time.Time           `json:"expiredTime" comment:"到期时间"`
	Details     []UsersQuotaDetails `json:"details" comment:"明细数据"`
	CreatedAt   time.Time           `json:"createdAt" comment:"创建时间"`
	UpdatedAt   time.Time           `json:"updatedAt" comment:"最后更新时间"`
}

type UsersQuotaResponse struct {
	RequestId string     `json:"requestId,omitempty"`
	Code      int32      `json:"code,omitempty"`
	Msg       string     `json:"msg,omitempty"`
	Status    string     `json:"status,omitempty"`
	Data      UsersQuota `json:"data,omitempty"`
}
