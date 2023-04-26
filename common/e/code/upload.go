package code

import k2Error "github.com/kingwel-xie/k2/common/error"

// 上传
const (
	// 上传
	errUploadFileEmpty int = iota + 100401
	errUploadFileTooLarge
	errUploadFileIpfsCluster
	errUploadFileIpfs
	errSha256Verify
	errUploadedChunk
	errUploadChunkNotCompleted
	errUploadDagBuilding
	errUploadUserAvailableSpaceLimit
	errUploadUserAvailableFilesLimit
	errUploadDirSubfileRelativePathEmpty
	errUploadDirSubfileNotCompleted
	errUploadDirTotalFileAmountLessThanUploaded
	errUploadObjectSizeTooLarge
)

var (
	// 上传
	ErrUploadFileEmpty                          = k2Error.New(errUploadFileEmpty, "参数无效，文件内容不能为空", "Invalid parameter,file content cannot be empty")
	ErrUploadFileTooLarge                       = k2Error.New(errUploadFileTooLarge, "上传文件过大,不能超过44M", "The uploaded pdf file is too large and cannot exceed 44M")
	ErrUploadFileIpfsCluster                    = k2Error.New(errUploadFileIpfsCluster, "创建副本失败", "Replica creation failure") // 上传到ipfscluster失败
	ErrUploadFileIpfs                           = k2Error.New(errUploadFileIpfs, "上传到ipfs失败", "Description Failed to upload to ipfs")
	ErrSha256Verify                             = k2Error.New(errSha256Verify, "验证sha256失败", "Failed to verify sha256")
	ErrUploadedChunk                            = k2Error.New(errUploadedChunk, "上传分片列表异常", "Failed to verify sha256")
	ErrUploadChunkNotCompleted                  = k2Error.New(errUploadChunkNotCompleted, "块上传未完成", "The chunk upload is not completed.")
	ErrUploadDagBuilding                        = k2Error.New(errUploadDagBuilding, "生成DAG中，请等会再试.", "Generating DAG, please try again later.")
	ErrUploadUserAvailableSpaceLimit            = k2Error.New(errUploadUserAvailableSpaceLimit, "计划存储空间不足上传失败，请升级计划", "Plan storage space insufficient upload failed, please upgrade plan.")
	ErrUploadUserAvailableFilesLimit            = k2Error.New(errUploadUserAvailableFilesLimit, "计划到达数量上限上传失败，请升级计划", "The upload failed to reach the limit of the plan. Please upgrade the plan.")
	ErrUploadDirSubfileRelativePathEmpty        = k2Error.New(errUploadDirSubfileRelativePathEmpty, "目录子文件relativePath不能为空", "The directory subfile relativePath cannot be empty.")
	ErrUploadDirSubfileNotCompleted             = k2Error.New(errUploadDirSubfileNotCompleted, "目录子文件上传未完成", "The directory subfile upload is not completed.")
	ErrUploadDirTotalFileAmountLessThanUploaded = k2Error.New(errUploadDirTotalFileAmountLessThanUploaded, "目录总数比已经上传的目录文件数量少", "The total number of subfiles is less than the number of subfiles that have been uploaded.")
	ErrUploadObjectSizeTooLarge                 = k2Error.New(errUploadObjectSizeTooLarge, "当前上传对象已 7GB 平台限制，上传失败", "The current uploaded object has 7GB platform limit, upload failed.")
)
