package dict

/**
	leveldb 数据前缀
*/
const (
	//metric 修改大文件过期时间异常数据LargeFileModifyExtimeFailed
	LARGETFILE_MODIFY_EXPIRED_FAILED = "/largefile/expired/failed/%s"

	//queue 等待更新过期时间的大文件
	LARGETFILE_TOBE_PROCESSED = "/largefile/expired/processed/%s"
)
