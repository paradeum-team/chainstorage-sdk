package vo

type UpdateParamGfs struct {
	//参数是自定义的参数，不能是系统参数
	ParameterName  string `json:"parameterName"  binding:"required" example:"service"`
	//参数对应的值
	ParameterValue string `json:"parameterValue"  binding:"required" example:"gateway"`
}
type UpdateParamGfsSign struct {
	//参数支持[afid, time_stamp_expired]或者是自定义参数
	ParameterName    string `json:"parameterName"  binding:"required" example:"parameterName"`
	//time_stamp_expired 格式 20191212235959
	ParameterValue   string `json:"parameterValue"  binding:"required" example:"20191212235959"`
	//签名
	CombinationSign  string `json:"CombinationSign"  binding:"required" example:"30819f300d06092a864886f70d010101050003818d0030818902818100b1e77e09efe3e587829617723280da60cd330c63eaac1ee8f5b246ac2257646970f0ff7e08ac4d889d4e89d40d6f5f7a4fa3880c1068ac1b5bfe46c96a1f15f3f2e3b03a7369894347ee50ed6197a5f0547ec4b7945154582a8df012672b4f653037b164c2b1f99555c327f2b10746a33fd4834c50f404b14c3543648798ce2f0203010001" `
	//需要更新的值
	TimeStampUpdated string `json:"TimeStampUpdated"  binding:"required" example:"20201212235959"`
}

type CreateParamGfs struct {
	//加密算法只支持RSA、ECC
	AsymAlgo   string `json:"asymalgo" example:"RSA"`

	//RSA公钥需要做16进制转换
	AsymPubkey string `json:"asympubkey" example:"30819f300d06092a864886f70d010101050003818d0030818902818100b1e77e09efe3e587829617723280da60cd330c63eaac1ee8f5b246ac2257646970f0ff7e08ac4d889d4e89d40d6f5f7a4fa3880c1068ac1b5bfe46c96a1f15f3f2e3b03a7369894347ee50ed6197a5f0547ec4b7945154582a8df012672b4f653037b164c2b1f99555c327f2b10746a33fd4834c50f404b14c3543648798ce2f0203010001"`
}

type getVersionDataResp struct {
	Core    string `json:"core"`
	Verison string `json:"version"`
	Raw     string `json:"raw"`
}

type GetVersionRespSpecific struct {
	Code int                `json:"code"`
	Data getVersionDataResp `json:"data"`
	Msg  string             `json:"msg"`
}

type GetVersionResp struct {
	Bfs getVersionDataResp `json:"bfs"`
}
type UpdateParam struct {
	TimeStampExpired string `json:"time_stamp_expired"  binding:"required"`
}

type GeneralResp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type DumpFileVO struct {
	TimeStampExpired string `json:"time_stamp_expired"`
}

type SeedIndexVO struct {
	Algotype string `json:"algotype" binding:"required"`
	Dgst string `json:"dgst" binding:"required"`
}
type PNAddressVO struct {
	Addresses []string `json:"addresses"`
}
