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
