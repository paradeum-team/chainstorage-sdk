package code

// 登录注册
const (
	errLogin int = iota + 100101
	errLoginVerifySignature
	errLoginInvalidWalletAddr
	errLoginWalletAddrOrSignature
	errUserNotFound
	errUserTokenExpired
	errNoPermissionFailed
	errUpdateUserIdNoSame
	errUserRefreshTokenExpired
	errUserInvalidMailbox
	errUserInvalidMailVerificationInfo
	errUserMailAlreadyUsed
	errUploadAvatarFail
	errUserProhibitChangeVerifiedMail
	errUserVerifiedMailSendLimitExceeded
	errUserTokenInvalid
)

// 桶
const (
	errBucketNotFound int = iota + 100201
	errBucketObjectNotFound
	errInvalidBucketName
	errBucketNameConflict
	errStorageNetworkMustSet
	errBucketPrincipleMustSet
	errBucketMustBeEmpty
	errBucketVolumnStatFail
	errBucketQuotaFetchFail
	errUserQuotaUpdateFail
	errOnlyCreate1BucketForSameStorageNetwork
)

// 文件列表
const (
	errObjectNotFound int = iota + 100301
	//errBucketObjectNotFound
	errInvalidObjectName
	errObjectNameConflict
	errObjectSetReferenceCounterFail
)

// ApiKey
const (
	errApiKeyNotFound int = iota + 100801
	errInvalidApiKeyName
	errApiKeyNameConflict
	errApiKeyGenerateFail
	errApiKeyPermissionTypeMustSet
	errApiKeyPermissionMustSet
	errApiKeyDataScopeMustSet
	errApiKeyPinningServicePermissionMustSet
)

var (

	//登录注册
	ErrWalletInvalidFailed               = NewBizError(errLogin, "钱包地址非法空", "Invalid wallet")
	ErrLoginFailed                       = NewBizError(errLogin, "登录失败", "Login failed")
	ErrLoginVerifySignatureFailed        = NewBizError(errLoginVerifySignature, "验证签名失败", "Verification signature failed")
	ErrLoginInvalidWalletAddrFailed      = NewBizError(errLoginInvalidWalletAddr, "无效的钱包地址", "Invalid wallet address")
	ErrLoginWalletAddrOrSignatureFailed  = NewBizError(errLoginWalletAddrOrSignature, "钱包地址或者签名格式错误", "Wrong wallet address or signature format")
	ErrUserNotFound                      = NewBizError(errUserNotFound, "用户不存在", "User not found")
	ErrUserTokenExpired                  = NewBizError(errUserTokenExpired, "登录已过期，请重新登录", "User token is expired, please login again")
	ErrNoPermissionFailed                = NewBizError(errNoPermissionFailed, "您无权访问此资源", "You don't have permission to access this resource")
	ErrUpdateUserIdNoSame                = NewBizError(errUpdateUserIdNoSame, "不是本人操作，不能修改", "It cannot be modified unless you operate it yourself")
	ErrUserRefreshTokenExpired           = NewBizError(errUserRefreshTokenExpired, "身份验证会话已过期，请重新登录", "The authentication session has expired, please sign-in again")
	ErrUserInvalidMailbox                = NewBizError(errUserInvalidMailbox, "无效的电子邮箱地址", "无效的电子邮箱地址")
	ErrUserInvalidMailVerificationInfo   = NewBizError(errUserInvalidMailVerificationInfo, "无效的电子邮箱认证信息", "无效的电子邮箱认证信息")
	ErrUserMailAlreadyUsed               = NewBizError(errUserMailAlreadyUsed, "电子邮箱已经被使用", "电子邮箱已经被使用")
	ErrUploadAvatarFail                  = NewBizError(errUploadAvatarFail, "上传头像失败", "上传头像失败")
	ErrUserProhibitChangeVerifiedMail    = NewBizError(errUserProhibitChangeVerifiedMail, "禁止修改已验证邮箱", "禁止修改已验证邮箱")
	ErrUserVerifiedMailSendLimitExceeded = NewBizError(errUserVerifiedMailSendLimitExceeded, "验证邮件发送过于频繁", "验证邮件发送过于频繁")
	ErrUserTokenInvalid                  = NewBizError(errUserTokenInvalid, "用户Token无效，请重新登录", "User token is invalid, please login again")

	// 桶
	ErrBucketNotFound                         = NewBizError(errBucketNotFound, "桶数据不存在", "桶数据不存在")
	ErrInvalidBucketName                      = NewBizError(errInvalidBucketName, "桶名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试", "桶名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试")
	ErrBucketNameConflict                     = NewBizError(errBucketNameConflict, "桶名称冲突，桶名称必须全平台唯一，请重新尝试", "桶名称冲突，桶名称必须全平台唯一，请重新尝试")
	ErrStorageNetworkMustSet                  = NewBizError(errStorageNetworkMustSet, "存储网络名称必须正确设置", "存储网络名称必须正确设置")
	ErrBucketPrincipleMustSet                 = NewBizError(errBucketPrincipleMustSet, "桶策略必须正确设置", "桶策略必须正确设置")
	ErrBucketMustBeEmpty                      = NewBizError(errBucketMustBeEmpty, "不能删除非空桶", "不能删除非空桶")
	ErrBucketObjectNotFound                   = NewBizError(errBucketObjectNotFound, "桶对象数据不存在", "桶对象数据不存在")
	ErrBucketVolumnStatFail                   = NewBizError(errBucketVolumnStatFail, "桶容量统计失败", "桶容量统计失败")
	ErrBucketQuotaFetchFail                   = NewBizError(errBucketQuotaFetchFail, "桶容量配额获取失败", "桶容量配额获取失败")
	ErrUserQuotaUpdateFail                    = NewBizError(errUserQuotaUpdateFail, "用户容量配额更新失败", "用户容量配额更新失败")
	ErrOnlyCreate1BucketForSameStorageNetwork = NewBizError(errOnlyCreate1BucketForSameStorageNetwork, "基础版本中，一种网络类型只允许创建一个桶", "基础版本中，一种网络类型只允许创建一个桶")

	// 文件列表
	ErrObjectNotFound                = NewBizError(errObjectNotFound, "对象数据不存在", "对象数据不存在")
	ErrInvalidObjectName             = NewBizError(errInvalidObjectName, "对象名称异常，名称范围必须在 1-255 个字符之间，并且不能包含非法字符，以及使用操作系统保留字，请重新尝试", "对象名称异常，名称范围必须在 1-255 个字符之间，并且不能包含非法字符，以及使用操作系统保留字，请重新尝试")
	ErrObjectNameConflictInBucket    = NewBizError(errObjectNameConflict, "对象名称冲突，对象名称必须在桶内唯一，请重新尝试或者确认进行覆盖操作", "对象名称冲突，对象名称必须在桶内唯一，请重新尝试或者确认进行覆盖操作")
	ErrObjectSetReferenceCounterFail = NewBizError(errObjectSetReferenceCounterFail, "设置对象引用计数器异常", "设置对象引用计数器异常")

	// ApiKey
	ErrApiKeyNotFound                        = NewBizError(errApiKeyNotFound, "Api数据不存在", "Api数据不存在")
	ErrInvalidApiKeyName                     = NewBizError(errInvalidApiKeyName, "Api名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试", "Api名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试")
	ErrApiKeyNameConflict                    = NewBizError(errApiKeyNameConflict, "Api名称冲突，ApiKey名称必须唯一，请重新尝试", "Api名称冲突，ApiKey名称必须唯一，请重新尝试")
	ErrApiKeyGenerateFail                    = NewBizError(errApiKeyGenerateFail, "ApiKey生成失败", "ApiKey生成失败")
	ErrApiKeyPermissionTypeMustSet           = NewBizError(errApiKeyPermissionTypeMustSet, "管理员设置必须正确设置", "管理员设置必须正确设置")
	ErrApiKeyPermissionMustSet               = NewBizError(errApiKeyPermissionMustSet, "API服务权限必须正确设置", "API服务权限必须正确设置")
	ErrApiKeyDataScopeMustSet                = NewBizError(errApiKeyDataScopeMustSet, "数据范围(桶)必须正确设置", "数据范围(桶)必须正确设置")
	ErrApiKeyPinningServicePermissionMustSet = NewBizError(errApiKeyPinningServicePermissionMustSet, "PinningServiceAPI权限必须正确设置", "PinningServiceAPI权限必须正确设置")
)

const (
	//收藏
	errFavoriteUnbound int = iota + 100401
	errFavoriteBoundFailed

	//上传
	//errUploadFileEmpty
	errUploadFilePdfTooLarge
	errUploadFileThumbnailTooLarge
	errUploadFileToLocal
	errUploadFileIpfsPathEmpty
	errUploadFileIpfsPathIllegal
)

// 校验字段
const (
	errVerifyTitle int = iota + 100501
	errVerifyPublishDescription
	errVerifyBio
	errVerifyNickName
	errVerifyArticleAbstract
	errVerifyBountyDetail
	errVerifyWalletAddrFormat
	errVerifyArticleContents
	errVerifyAvatarEmpty
	errVerifyAvatarTooLarge
	errVerifyAvatarType
)

// system
const (
	errSysUuidGenFailed = iota + 100601
	errSysUuidNotEmpty
	errSysUuidIdempotent
)

var (
	// ipfs
	//ErrUploadFileEmpty             = NewBizError(errUploadFileEmpty, "参数无效，文件内容不能为空", "Invalid parameter,file content cannot be empty")
	ErrUploadFilePdfTooLarge       = NewBizError(errUploadFilePdfTooLarge, "上传pdf文件过大,不能超过20M", "The uploaded pdf file is too large and cannot exceed 20M")
	ErrUploadFileThumbnailTooLarge = NewBizError(errUploadFileThumbnailTooLarge, "上传缩图文件过大,不能超过5M", "The uploaded thumbnail file is too large and cannot exceed 5M")
	ErrUploadFileToLocal           = NewBizError(errUploadFileToLocal, "本地上传文件失败", "Failed to upload files locally")
	ErrUploadFileIpfsPathEmpty     = NewBizError(errUploadFileIpfsPathEmpty, "Ipfs路径不能为空", "Ipfs path cannot be empty")
	ErrUploadFileIpfsPathIllegal   = NewBizError(errUploadFileIpfsPathIllegal, "Ipfs路径格式错误,必须是 /ipfs/${cid}?filename=xxx.xx", "The format of the Ipfs path is incorrect. It must be/ipfs/${cid}? filename=xxx.xx")

	//校验字段
	ErrVerifyTitle              = NewBizError(errVerifyTitle, "验证标题最多80字符", "Validation title can be up to 80 characters")
	ErrVerifyPublishDescription = NewBizError(errVerifyPublishDescription, "期刊描述最多300字符", "The publication description can be 300 characters at most")
	ErrVerifyBio                = NewBizError(errVerifyBio, "个人简介最多150字符", "User profle can be up to 150 characters")
	ErrVerifyNickName           = NewBizError(errVerifyNickName, "用户昵称最多50字符", "User nickname can be up to 50 characters")
	ErrVerifyArticleAbstract    = NewBizError(errVerifyArticleAbstract, "文章描述最多600字符", "Article description up to 600 characters")
	ErrVerifyBountyDetail       = NewBizError(errVerifyBountyDetail, "征稿描述长度最多5000字符", "The length of bounty description is 5000 characters at most")
	ErrVerifyWalletAddrFormat   = NewBizError(errVerifyWalletAddrFormat, "钱包地址格式错误", "Wallet address format error")
	ErrVerifyArticleContents    = NewBizError(errVerifyArticleContents, "文章内容最多20000字符", "Article contents up to 20000 characters")
	ErrVerifyAvatarEmpty        = NewBizError(errVerifyAvatarEmpty, "头像文件不能为空", "The avatar file cannot be empty")
	ErrVerifyAvatarTooLarge     = NewBizError(errVerifyAvatarTooLarge, "头像文件过大,不能超过2M", "The avatar file is too large to exceed 2M")
	ErrVerifyAvatarType         = NewBizError(errVerifyAvatarType, "上传的头像图片类型不支持", "The type of avatar image uploaded is not supported.")

	//sys
	ErrSysUuidGenFailed  = NewBizError(errSysUuidGenFailed, "生成UUID失败,请重新生成", "Failed to generate UUID, please regenerate")
	ErrSysUuidNotEmpty   = NewBizError(errSysUuidNotEmpty, "UUID必须不为空", "Header uuid must not be empty")
	ErrSysUuidIdempotent = NewBizError(errSysUuidIdempotent, "接口幂等,不做处理", "Interface idempotent, no processing")
)
