package vo

import (
	"github.com/gin-gonic/gin"
)

/**
	文件下载参数实体
 */
type DownloadParamsVO struct {
	FileName string `json:"file_name"`
	IsBigfile bool `json:"is_bigfile"`
	IsShortApi bool `json:"is_short_api"`
	GinContext *gin.Context `json:"gin_context"`
	IsSeedFile bool `json:"is_seed_file"`
}

/**
	文件上传参数实体
 */
type UploadParamsVO struct {
	RnodeRoot string `json:"rnode_root"`
	Days string `json:"days"`
	Asympubkey string `json:"asympubkey"`
	Sympassen string `json:"sympassen"`
	EncryptedType string `json:"encrypted_type"`
	LocalFileName string `json:"local_file_name"`
	SourceFileName string `json:"source_file_name"`
	//源文件路径
	SourceFilePath string `json:"source_file_path"`
	GinContext *gin.Context `json:"gin_context"`
}