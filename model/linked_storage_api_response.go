package model

/**
 *简单操作
 **/
// swagger:model
type OptResponse struct {
	Status bool   `json:"status"`
	Note   string `json:"note"`
	Raw    string `json:"raw"`
}

// 返回对象
type ResponseVO struct {
	Code     int         `json:"code"`
	HttpCode int         `json:"http_code"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data"`
}
