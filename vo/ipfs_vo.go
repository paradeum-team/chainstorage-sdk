package vo

import (
	"fmt"
	"net/url"
	"time"
)

var pinOptionsMetaPrefix = "meta-"

type CidInfo struct {
	Cids       []CidWithSn `json:"cids" codec:"cs"`
	AlgoType   string      `json:"algo_type" codec:"a"`
	Value      string      `json:"value" codec:"v"`
	ExtendData string      `json:"extend_data" codec:"e,omitempty"`
	IsRawID    bool        `json:"is_rawid" codec:"r,omitempty"`
}

/**
 * Cid顺序列表项目
 */
type CidWithSn struct {
	Sn  uint32 `json:"sn" codec:"s,omitempty"`
	Cid string `json:"cid" codec:"c"`
}

type Pin struct {
	Cid     string    `json:"cid" codec:"c"`
	Created time.Time `json:"created" codec:"t,omitempty"`

	PinOptions
}

// PinOptions wraps user-defined options for Pins
type PinOptions struct {
	ReplicationFactorMin int               `json:"replication_factor_min" codec:"rn,omitempty"`
	ReplicationFactorMax int               `json:"replication_factor_max" codec:"rx,omitempty"`
	Name                 string            `json:"name" codec:"n,omitempty"`
	ExpireAt             time.Time         `json:"expire_at" codec:"e,omitempty"`
	Metadata             map[string]string `json:"metadata" codec:"m,omitempty"`
}

// ToQuery returns the PinOption as query arguments.
func (po PinOptions) ToQuery() (string, error) {
	q := url.Values{}
	q.Set("name", po.Name)

	if !po.ExpireAt.IsZero() {
		v, err := po.ExpireAt.MarshalText()
		if err != nil {
			return "", err
		}
		q.Set("expire-at", string(v))
		//expireAt := po.ExpireAt.Format(dict.SysTimeFmt)
		//q.Set("expire-at", expireAt)
	}

	for k, v := range po.Metadata {
		if k == "" {
			continue
		}

		q.Set(fmt.Sprintf("%s%s", pinOptionsMetaPrefix, k), v)
	}

	return q.Encode(), nil
}

type AddedOutput struct {
	Afid string `json:"afid" codec:"ad,omitempty"`
	Name string `json:"name" codec:"n,omitempty"`
	Cid  string `json:"cid" codec:"c"`
	//Bytes       uint64   `json:"bytes,omitempty" codec:"b,omitempty"`
	//Size        uint64   `json:"size,omitempty" codec:"s,omitempty"`
	//Allocations []string `json:"allocations,omitempty" codec:"a,omitempty"`
}

type Metric struct {
	Name          string `json:"name" codec:"n,omitempty"`
	Peer          string `json:"peer" codec:"p,omitempty"`
	Value         string `json:"value" codec:"v,omitempty"`
	Expire        int64  `json:"expire" codec:"e,omitempty"`
	Valid         bool   `json:"valid" codec:"d,omitempty"`
	Weight        int64  `json:"weight" codec:"w,omitempty"`
	Partitionable bool   `json:"partitionable" codec:"o,omitempty"`
	ReceivedAt    int64  `json:"received_at" codec:"t,omitempty"` // ReceivedAt contains a UnixNano timestamp
}

type VersionResp struct {
	Verison string `json:"version"`
}
