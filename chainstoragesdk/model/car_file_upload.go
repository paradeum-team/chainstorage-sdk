package model

type CarFileUploadReq struct {
	BucketId        int    `json:"bucketId" binding:"required" comment:"桶主键"`
	ObjectName      string `json:"objectName" comment:"对象名称（255字限制）"`
	ObjectTypeCode  int    `json:"objectTypeCode" comment:"对象类型编码"`
	ObjectSize      int64  `json:"objectSize" comment:"对象大小（字节）"`
	ObjectCid       string `json:"objectCid" comment:"对象CID"`
	FileDestination string `json:"-" comment:"文件路径"`
	RawSha256       string `json:"rawSha256" binding:"required" comment:"原始文件sha256"`
	ShardingSha256  string `json:"shardingSha256" comment:"分片sha256"`
	ShardingNo      int    `json:"shardingNo" comment:"分片序号"`
	ShardingAmount  int    `json:"shardingAmount" comment:"分片数量"`
	CarFileCid      string `json:"carFileCid" comment:"Car文件CID"`
}

type ShardingCarFileUploadResponse struct {
	RequestId string        `json:"requestId,omitempty"`
	Code      int32         `json:"code,omitempty"`
	Msg       string        `json:"msg,omitempty"`
	Status    string        `json:"status,omitempty"`
	Data      CarFileUpload `json:"data,omitempty"`
}

type CarFileUpload struct {
	BucketId       int    `json:"bucketId" binding:"required" comment:"桶主键"`
	ObjectName     string `json:"objectName" comment:"对象名称（255字限制）"`
	ObjectTypeCode int    `json:"objectTypeCode" comment:"对象类型编码"`
	ObjectSize     int64  `json:"objectSize" comment:"对象大小（字节）"`
	ObjectCid      string `json:"objectCid" comment:"对象CID"`
	RawSha256      string `json:"rawSha256" binding:"required" comment:"原始文件sha256"`
	ShardingSha256 string `json:"shardingSha256" comment:"分片sha256"`
	ShardingNo     int    `json:"shardingNo" comment:"分片序号"`
	ShardingAmount int    `json:"shardingAmount" comment:"分片数量"`
	CarFileCid     string `json:"carFileCid" comment:"Car文件CID"`
}

type ShardingCarFilesVerifyResponse struct {
	RawSha256      string   `json:"rawSha256" binding:"required" comment:"原始文件sha256"`
	ObjectName     string   `json:"objectName" comment:"文件名"`
	ShardingAmount int      `json:"shardingAmount" comment:"分片数量"`
	UploadStatus   int      `json:"uploadStatus" comment:"上传状态"`
	Uploaded       []string `json:"uploaded" comment:"上传完成分片列表"`
}
