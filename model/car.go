package model

import (
	"github.com/ipfs/go-cid"
	ipldfmt "github.com/ipfs/go-ipld-format"
)

type CarResponse struct {
	RequestId string      `json:"requestId,omitempty"`
	Code      int32       `json:"code,omitempty"`
	Msg       string      `json:"msg,omitempty"`
	Status    string      `json:"status,omitempty"`
	Data      interface{} `json:"data"`
}

type RootLink struct {
	ipldfmt.Link
	RootCid cid.Cid
}
