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
	errInvalidBucketId
	errStorageNetworkCodeMustSet
	errBucketBindWithGateway
)

// 文件列表
const (
	errObjectNotFound int = iota + 100301
	//errBucketObjectNotFound
	errInvalidObjectName
	errObjectNameConflict
	errObjectSetReferenceCounterFail
	errInvalidObjectCid
	errInvalidObjectId
	errInvalidObjectIds
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

// CAR文件上传
const (
	errCarUploadFileFail int = iota + 100901
	errCarUploadFileInvalidDataPath
	errCarUploadFileCreateCarFileFail
	errCarUploadFileParseCarFileFail
	errCarUploadFileComputeCarFileHashFail
	errCarUploadFileChunkCarFileFail
	errCarUploadFileReferenceObjcetFail
	errCarUploadFileCidNotEqualRawCid
	errCarUploadFileInvalidDataFolder
	errCarUploadFileExccedObjectAmountUsage
	errCarUploadFileExccedStorageSpaceUsage
	errCarUploadFileExccedUploadDirItems
)

var (
	//登录注册
	ErrWalletInvalidFailed               = NewBizError(errLogin, "无效的空钱包地址,请检查后请重试", "Invalid empty wallet address, please check and try again")
	ErrLoginFailed                       = NewBizError(errLogin, "登录失败,请重试", "Login failed, please try again")
	ErrLoginVerifySignatureFailed        = NewBizError(errLoginVerifySignature, "签名验证未通过,请重试", "Signature validation failed, please try again")
	ErrLoginInvalidWalletAddrFailed      = NewBizError(errLoginInvalidWalletAddr, "无效的钱包地址,请重试", "Invalid wallet address, please try again")
	ErrLoginWalletAddrOrSignatureFailed  = NewBizError(errLoginWalletAddrOrSignature, "无效的钱包地址或签名格式,请重试", "Incorrect wallet address or signature format, please try again")
	ErrUserNotFound                      = NewBizError(errUserNotFound, "该用户不存在,请重试", "This user does not exist, please try again")
	ErrUserTokenExpired                  = NewBizError(errUserTokenExpired, "登录已失效，请重新登录", "Login has expired, please log in again")
	ErrNoPermissionFailed                = NewBizError(errNoPermissionFailed, "对不起，您无权访问此资源", "Sorry, you do not have permission to access this resource")
	ErrUpdateUserIdNoSame                = NewBizError(errUpdateUserIdNoSame, "非本人操作，无法修改", "Cannot modify as not performed by the account owner")
	ErrUserRefreshTokenExpired           = NewBizError(errUserRefreshTokenExpired, "登录超时，请重新登录", "Login timed out, please log in again")
	ErrUserInvalidMailbox                = NewBizError(errUserInvalidMailbox, "请输入有效的电子邮箱地址", "Please enter a valid email address")
	ErrUserInvalidMailVerificationInfo   = NewBizError(errUserInvalidMailVerificationInfo, "认证信息不正确,请检查后请重试", "Incorrect authentication information,please check and try again")
	ErrUserMailAlreadyUsed               = NewBizError(errUserMailAlreadyUsed, "该邮箱已被占用", "This email is already in use")
	ErrUploadAvatarFail                  = NewBizError(errUploadAvatarFail, "头像上传失败，请重试", "Avatar upload failed, please try again")
	ErrUserProhibitChangeVerifiedMail    = NewBizError(errUserProhibitChangeVerifiedMail, "邮箱已经验证，无法修改", "Cannot modify verified email")
	ErrUserVerifiedMailSendLimitExceeded = NewBizError(errUserVerifiedMailSendLimitExceeded, "验证发送频率过高，请稍后再试", "Verification sending frequency is too high, please try again later")
	ErrUserTokenInvalid                  = NewBizError(errUserTokenInvalid, "登录已失效，请重新登录", "Login has expired, please log in again")

	// 桶
	ErrBucketNotFound                         = NewBizError(errBucketNotFound, "该数据桶不存在", "This bucket data does not exist")
	ErrInvalidBucketName                      = NewBizError(errInvalidBucketName, "桶名称必须是 3-63 个字符之间的小写字母、数字和破折号,请重新尝试", "Bucket name must be between 3-63 characters and can only contain lowercase letters、numbers and hyphens, please try again")
	ErrBucketNameConflict                     = NewBizError(errBucketNameConflict, "桶名称已被占用，请更换桶名称", "Bucket name is already taken, please choose a different name")
	ErrStorageNetworkMustSet                  = NewBizError(errStorageNetworkMustSet, "存储网络名称设置不正确,请重新尝试", "Incorrect storage network name settings, please try again")
	ErrBucketPrincipleMustSet                 = NewBizError(errBucketPrincipleMustSet, "请正确设置桶策略", "Please set the bucket policy correctly")
	ErrBucketMustBeEmpty                      = NewBizError(errBucketMustBeEmpty, "桶内有数据，无法删除", "The bucket contains data and cannot be deleted")
	ErrBucketObjectNotFound                   = NewBizError(errBucketObjectNotFound, "桶内没有该对象数据", "The object data is not in the bucket")
	ErrBucketVolumnStatFail                   = NewBizError(errBucketVolumnStatFail, "桶容量统计出错，请重试", "Error occurred while calculating bucket capacity, please try again")
	ErrBucketQuotaFetchFail                   = NewBizError(errBucketQuotaFetchFail, "桶容量配额获取出错，请重试", "Error occurred while getting bucket capacity quota, please try again")
	ErrUserQuotaUpdateFail                    = NewBizError(errUserQuotaUpdateFail, "桶容量配额更新出错，请重试", "Error occurred while updating bucket capacity quota, please try again")
	ErrOnlyCreate1BucketForSameStorageNetwork = NewBizError(errOnlyCreate1BucketForSameStorageNetwork, "基础版本限制，每种网络类型只能创建一个桶", "In the basic version, only one bucket can be created for each network type")
	ErrInvalidBucketId                        = NewBizError(errInvalidBucketId, "桶ID无效", "invalid bucket ID")
	ErrStorageNetworkCodeMustSet              = NewBizError(errStorageNetworkCodeMustSet, "存储网络编码设置不正确,请重新尝试", "Incorrect storage network code settings, please try again")
	ErrBucketBindWithGateway                  = NewBizError(errBucketBindWithGateway, "桶已经与网关绑定，无法删除", "The bucket is bound with gateway and cannot be deleted")

	// 文件列表
	ErrObjectNotFound                = NewBizError(errObjectNotFound, "该对象数据不存在", "The object data does not exist")
	ErrInvalidObjectName             = NewBizError(errInvalidObjectName, "对象名称必须是 1-255 个字符之间，不能包含非法字符，请重试", "Object name must be between 1-255 characters and cannot contain invalid characters, please try again")
	ErrObjectNameConflictInBucket    = NewBizError(errObjectNameConflict, "对象名称已存在，是否覆盖原有对象", "Object name already exists, do you want to overwrite the existing object?")
	ErrObjectSetReferenceCounterFail = NewBizError(errObjectSetReferenceCounterFail, "对象引用计数器操作出错", "Error occurred while operating object reference counter")
	ErrInvalidObjectCid              = NewBizError(errInvalidObjectCid, "无效的对象CID", "Invalid object CID")
	ErrInvalidObjectId               = NewBizError(errInvalidObjectId, "无效的对象ID", "invalid object ID")
	ErrInvalidObjectIds              = NewBizError(errInvalidObjectIds, "无效的对象ID列表", "invalid object ID list")

	// ApiKey
	ErrApiKeyNotFound                        = NewBizError(errApiKeyNotFound, "该 APIKey 不存在", "The APIKey does not exist")
	ErrInvalidApiKeyName                     = NewBizError(errInvalidApiKeyName, "APIKey 名称必须是 3-63 个字符之间的小写字母、数字和破折号，请重新尝试", "APIKey name must be between 3-63 characters and can only contain lowercase letters, numbers, and hyphens, please try again")
	ErrApiKeyNameConflict                    = NewBizError(errApiKeyNameConflict, "APIKey 名称已存在，请重新尝试", "APIKey name already exists, please try again")
	ErrApiKeyGenerateFail                    = NewBizError(errApiKeyGenerateFail, "APIKey 创建失败，请重试", "Failed to create APIKey, please try again")
	ErrApiKeyPermissionTypeMustSet           = NewBizError(errApiKeyPermissionTypeMustSet, "管理员设置不正确，请重试", "Incorrect admin settings, please try again")
	ErrApiKeyPermissionMustSet               = NewBizError(errApiKeyPermissionMustSet, "APIKey 权限设置错误，请重试", "APIKey permissions settings are incorrect, please try again")
	ErrApiKeyDataScopeMustSet                = NewBizError(errApiKeyDataScopeMustSet, "APIKey 数据范围设置不正确，请重试", "Incorrect APIKey data range setting, please try again")
	ErrApiKeyPinningServicePermissionMustSet = NewBizError(errApiKeyPinningServicePermissionMustSet, "PinningServiceAPI 权限设置不正确，请重试", "Incorrect PinningServiceAPI permissions setting，please try again")

	// CAR文件上传
	ErrCarUploadFileFail                    = NewBizError(errCarUploadFileFail, "CAR上传文件失败", "Fail to upload CAR file")
	ErrCarUploadFileInvalidDataPath         = NewBizError(errCarUploadFileInvalidDataPath, "无效的上传数据路径", "Invalid uploading data path")
	ErrCarUploadFileCreateCarFileFail       = NewBizError(errCarUploadFileCreateCarFileFail, "创建CAR文件失败", "Fail to create CAR file")
	ErrCarUploadFileParseCarFileFail        = NewBizError(errCarUploadFileParseCarFileFail, "解析CAR文件失败", "Fail to parse CAR file")
	ErrCarUploadFileComputeCarFileHashFail  = NewBizError(errCarUploadFileComputeCarFileHashFail, "CAR文件HASH计算失败", "Fail to compute CAR file HASH")
	ErrCarUploadFileChunkCarFileFail        = NewBizError(errCarUploadFileChunkCarFileFail, "生成CAR文件分片操作失败", "Fail to chunk CAR file")
	ErrCarUploadFileReferenceObjcetFail     = NewBizError(errCarUploadFileReferenceObjcetFail, "执行CID秒传操作失败", "Fail to reference object by CID")
	ErrCarUploadFileInvalidDataFolder       = NewBizError(errCarUploadFileInvalidDataFolder, "上传目录为空或者目录中的数据无效", "Uploading folder is empty, or uploading data is invalid in the folder")
	ErrCarUploadFileExccedObjectAmountUsage = NewBizError(errCarUploadFileExccedObjectAmountUsage, "超过对象存储限制", "Exceed the limitation of object amount")
	ErrCarUploadFileExccedStorageSpaceUsage = NewBizError(errCarUploadFileExccedStorageSpaceUsage, "超过空间存储限制", "Exceed the limitation of storage space")
	ErrCarUploadFileExccedUploadDirItems    = NewBizError(errCarUploadFileExccedUploadDirItems, "超过上传文件夹条目限制", "Exceed the limitation of entries in uploading folder")

	////登录注册
	//ErrWalletInvalidFailed               = NewBizError(errLogin, "钱包地址非法空", "Invalid wallet")
	//ErrLoginFailed                       = NewBizError(errLogin, "登录失败", "Login failed")
	//ErrLoginVerifySignatureFailed        = NewBizError(errLoginVerifySignature, "验证签名失败", "Verification signature failed")
	//ErrLoginInvalidWalletAddrFailed      = NewBizError(errLoginInvalidWalletAddr, "无效的钱包地址", "Invalid wallet address")
	//ErrLoginWalletAddrOrSignatureFailed  = NewBizError(errLoginWalletAddrOrSignature, "钱包地址或者签名格式错误", "Wrong wallet address or signature format")
	//ErrUserNotFound                      = NewBizError(errUserNotFound, "用户不存在", "User not found")
	//ErrUserTokenExpired                  = NewBizError(errUserTokenExpired, "登录已过期，请重新登录", "User token is expired, please login again")
	//ErrNoPermissionFailed                = NewBizError(errNoPermissionFailed, "您无权访问此资源", "You don't have permission to access this resource")
	//ErrUpdateUserIdNoSame                = NewBizError(errUpdateUserIdNoSame, "不是本人操作，不能修改", "It cannot be modified unless you operate it yourself")
	//ErrUserRefreshTokenExpired           = NewBizError(errUserRefreshTokenExpired, "身份验证会话已过期，请重新登录", "The authentication session has expired, please sign-in again")
	//ErrUserInvalidMailbox                = NewBizError(errUserInvalidMailbox, "无效的电子邮箱地址", "无效的电子邮箱地址")
	//ErrUserInvalidMailVerificationInfo   = NewBizError(errUserInvalidMailVerificationInfo, "无效的电子邮箱认证信息", "无效的电子邮箱认证信息")
	//ErrUserMailAlreadyUsed               = NewBizError(errUserMailAlreadyUsed, "电子邮箱已经被使用", "电子邮箱已经被使用")
	//ErrUploadAvatarFail                  = NewBizError(errUploadAvatarFail, "上传头像失败", "上传头像失败")
	//ErrUserProhibitChangeVerifiedMail    = NewBizError(errUserProhibitChangeVerifiedMail, "禁止修改已验证邮箱", "禁止修改已验证邮箱")
	//ErrUserVerifiedMailSendLimitExceeded = NewBizError(errUserVerifiedMailSendLimitExceeded, "验证邮件发送过于频繁", "验证邮件发送过于频繁")
	//ErrUserTokenInvalid                  = NewBizError(errUserTokenInvalid, "用户Token无效，请重新登录", "User token is invalid, please login again")
	//
	//// 桶
	//ErrBucketNotFound                         = NewBizError(errBucketNotFound, "桶数据不存在", "桶数据不存在")
	//ErrInvalidBucketName                      = NewBizError(errInvalidBucketName, "桶名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试", "桶名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试")
	//ErrBucketNameConflict                     = NewBizError(errBucketNameConflict, "桶名称冲突，桶名称必须全平台唯一，请重新尝试", "桶名称冲突，桶名称必须全平台唯一，请重新尝试")
	//ErrStorageNetworkMustSet                  = NewBizError(errStorageNetworkMustSet, "存储网络名称必须正确设置", "存储网络名称必须正确设置")
	//ErrBucketPrincipleMustSet                 = NewBizError(errBucketPrincipleMustSet, "桶策略必须正确设置", "桶策略必须正确设置")
	//ErrBucketMustBeEmpty                      = NewBizError(errBucketMustBeEmpty, "不能删除非空桶", "不能删除非空桶")
	//ErrBucketObjectNotFound                   = NewBizError(errBucketObjectNotFound, "桶对象数据不存在", "桶对象数据不存在")
	//ErrBucketVolumnStatFail                   = NewBizError(errBucketVolumnStatFail, "桶容量统计失败", "桶容量统计失败")
	//ErrBucketQuotaFetchFail                   = NewBizError(errBucketQuotaFetchFail, "桶容量配额获取失败", "桶容量配额获取失败")
	//ErrUserQuotaUpdateFail                    = NewBizError(errUserQuotaUpdateFail, "用户容量配额更新失败", "用户容量配额更新失败")
	//ErrOnlyCreate1BucketForSameStorageNetwork = NewBizError(errOnlyCreate1BucketForSameStorageNetwork, "基础版本中，一种网络类型只允许创建一个桶", "基础版本中，一种网络类型只允许创建一个桶")
	//
	//// 文件列表
	//ErrObjectNotFound                = NewBizError(errObjectNotFound, "对象数据不存在", "对象数据不存在")
	//ErrInvalidObjectName             = NewBizError(errInvalidObjectName, "对象名称异常，名称范围必须在 1-255 个字符之间，并且不能包含非法字符，以及使用操作系统保留字，请重新尝试", "对象名称异常，名称范围必须在 1-255 个字符之间，并且不能包含非法字符，以及使用操作系统保留字，请重新尝试")
	//ErrObjectNameConflictInBucket    = NewBizError(errObjectNameConflict, "对象名称冲突，对象名称必须在桶内唯一，请重新尝试或者确认进行覆盖操作", "对象名称冲突，对象名称必须在桶内唯一，请重新尝试或者确认进行覆盖操作")
	//ErrObjectSetReferenceCounterFail = NewBizError(errObjectSetReferenceCounterFail, "设置对象引用计数器异常", "设置对象引用计数器异常")
	//ErrInvalidObjectCid              = NewBizError(errInvalidObjectCid, "无效的对象CID", "无效的对象CID")
	//
	//// ApiKey
	//ErrApiKeyNotFound                        = NewBizError(errApiKeyNotFound, "Api数据不存在", "Api数据不存在")
	//ErrInvalidApiKeyName                     = NewBizError(errInvalidApiKeyName, "Api名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试", "Api名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试")
	//ErrApiKeyNameConflict                    = NewBizError(errApiKeyNameConflict, "Api名称冲突，ApiKey名称必须唯一，请重新尝试", "Api名称冲突，ApiKey名称必须唯一，请重新尝试")
	//ErrApiKeyGenerateFail                    = NewBizError(errApiKeyGenerateFail, "ApiKey生成失败", "ApiKey生成失败")
	//ErrApiKeyPermissionTypeMustSet           = NewBizError(errApiKeyPermissionTypeMustSet, "管理员设置必须正确设置", "管理员设置必须正确设置")
	//ErrApiKeyPermissionMustSet               = NewBizError(errApiKeyPermissionMustSet, "API服务权限必须正确设置", "API服务权限必须正确设置")
	//ErrApiKeyDataScopeMustSet                = NewBizError(errApiKeyDataScopeMustSet, "数据范围(桶)必须正确设置", "数据范围(桶)必须正确设置")
	//ErrApiKeyPinningServicePermissionMustSet = NewBizError(errApiKeyPinningServicePermissionMustSet, "PinningServiceAPI权限必须正确设置", "PinningServiceAPI权限必须正确设置")
	//
	//// CAR文件上传
	//ErrCarUploadFileParseFail = NewBizError(errCarUploadFileFail, "CAR上传文件解析失败", "CAR上传文件解析失败")
)

//var (
//
//	//登录注册
//	ErrWalletInvalidFailed               = NewBizError(errLogin, "钱包地址非法空", "Invalid wallet")
//	ErrLoginFailed                       = NewBizError(errLogin, "登录失败", "Login failed")
//	ErrLoginVerifySignatureFailed        = NewBizError(errLoginVerifySignature, "验证签名失败", "Verification signature failed")
//	ErrLoginInvalidWalletAddrFailed      = NewBizError(errLoginInvalidWalletAddr, "无效的钱包地址", "Invalid wallet address")
//	ErrLoginWalletAddrOrSignatureFailed  = NewBizError(errLoginWalletAddrOrSignature, "钱包地址或者签名格式错误", "Wrong wallet address or signature format")
//	ErrUserNotFound                      = NewBizError(errUserNotFound, "用户不存在", "User not found")
//	ErrUserTokenExpired                  = NewBizError(errUserTokenExpired, "登录已过期，请重新登录", "User token is expired, please login again")
//	ErrNoPermissionFailed                = NewBizError(errNoPermissionFailed, "您无权访问此资源", "You don't have permission to access this resource")
//	ErrUpdateUserIdNoSame                = NewBizError(errUpdateUserIdNoSame, "不是本人操作，不能修改", "It cannot be modified unless you operate it yourself")
//	ErrUserRefreshTokenExpired           = NewBizError(errUserRefreshTokenExpired, "身份验证会话已过期，请重新登录", "The authentication session has expired, please sign-in again")
//	ErrUserInvalidMailbox                = NewBizError(errUserInvalidMailbox, "无效的电子邮箱地址", "无效的电子邮箱地址")
//	ErrUserInvalidMailVerificationInfo   = NewBizError(errUserInvalidMailVerificationInfo, "无效的电子邮箱认证信息", "无效的电子邮箱认证信息")
//	ErrUserMailAlreadyUsed               = NewBizError(errUserMailAlreadyUsed, "电子邮箱已经被使用", "电子邮箱已经被使用")
//	ErrUploadAvatarFail                  = NewBizError(errUploadAvatarFail, "上传头像失败", "上传头像失败")
//	ErrUserProhibitChangeVerifiedMail    = NewBizError(errUserProhibitChangeVerifiedMail, "禁止修改已验证邮箱", "禁止修改已验证邮箱")
//	ErrUserVerifiedMailSendLimitExceeded = NewBizError(errUserVerifiedMailSendLimitExceeded, "验证邮件发送过于频繁", "验证邮件发送过于频繁")
//	ErrUserTokenInvalid                  = NewBizError(errUserTokenInvalid, "用户Token无效，请重新登录", "User token is invalid, please login again")
//
//	// 桶
//	ErrBucketNotFound                         = NewBizError(errBucketNotFound, "桶数据不存在", "桶数据不存在")
//	ErrInvalidBucketName                      = NewBizError(errInvalidBucketName, "桶名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试", "桶名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试")
//	ErrBucketNameConflict                     = NewBizError(errBucketNameConflict, "桶名称冲突，桶名称必须全平台唯一，请重新尝试", "桶名称冲突，桶名称必须全平台唯一，请重新尝试")
//	ErrStorageNetworkMustSet                  = NewBizError(errStorageNetworkMustSet, "存储网络名称必须正确设置", "存储网络名称必须正确设置")
//	ErrBucketPrincipleMustSet                 = NewBizError(errBucketPrincipleMustSet, "桶策略必须正确设置", "桶策略必须正确设置")
//	ErrBucketMustBeEmpty                      = NewBizError(errBucketMustBeEmpty, "不能删除非空桶", "不能删除非空桶")
//	ErrBucketObjectNotFound                   = NewBizError(errBucketObjectNotFound, "桶对象数据不存在", "桶对象数据不存在")
//	ErrBucketVolumnStatFail                   = NewBizError(errBucketVolumnStatFail, "桶容量统计失败", "桶容量统计失败")
//	ErrBucketQuotaFetchFail                   = NewBizError(errBucketQuotaFetchFail, "桶容量配额获取失败", "桶容量配额获取失败")
//	ErrUserQuotaUpdateFail                    = NewBizError(errUserQuotaUpdateFail, "用户容量配额更新失败", "用户容量配额更新失败")
//	ErrOnlyCreate1BucketForSameStorageNetwork = NewBizError(errOnlyCreate1BucketForSameStorageNetwork, "基础版本中，一种网络类型只允许创建一个桶", "基础版本中，一种网络类型只允许创建一个桶")
//
//	// 文件列表
//	ErrObjectNotFound                = NewBizError(errObjectNotFound, "对象数据不存在", "对象数据不存在")
//	ErrInvalidObjectName             = NewBizError(errInvalidObjectName, "对象名称异常，名称范围必须在 1-255 个字符之间，并且不能包含非法字符，以及使用操作系统保留字，请重新尝试", "对象名称异常，名称范围必须在 1-255 个字符之间，并且不能包含非法字符，以及使用操作系统保留字，请重新尝试")
//	ErrObjectNameConflictInBucket    = NewBizError(errObjectNameConflict, "对象名称冲突，对象名称必须在桶内唯一，请重新尝试或者确认进行覆盖操作", "对象名称冲突，对象名称必须在桶内唯一，请重新尝试或者确认进行覆盖操作")
//	ErrObjectSetReferenceCounterFail = NewBizError(errObjectSetReferenceCounterFail, "设置对象引用计数器异常", "设置对象引用计数器异常")
//
//	// ApiKey
//	ErrApiKeyNotFound                        = NewBizError(errApiKeyNotFound, "Api数据不存在", "Api数据不存在")
//	ErrInvalidApiKeyName                     = NewBizError(errInvalidApiKeyName, "Api名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试", "Api名称异常，名称范围必须在 3-63 个字符之间并且只能包含小写字符、数字和破折号，请重新尝试")
//	ErrApiKeyNameConflict                    = NewBizError(errApiKeyNameConflict, "Api名称冲突，ApiKey名称必须唯一，请重新尝试", "Api名称冲突，ApiKey名称必须唯一，请重新尝试")
//	ErrApiKeyGenerateFail                    = NewBizError(errApiKeyGenerateFail, "ApiKey生成失败", "ApiKey生成失败")
//	ErrApiKeyPermissionTypeMustSet           = NewBizError(errApiKeyPermissionTypeMustSet, "管理员设置必须正确设置", "管理员设置必须正确设置")
//	ErrApiKeyPermissionMustSet               = NewBizError(errApiKeyPermissionMustSet, "API服务权限必须正确设置", "API服务权限必须正确设置")
//	ErrApiKeyDataScopeMustSet                = NewBizError(errApiKeyDataScopeMustSet, "数据范围(桶)必须正确设置", "数据范围(桶)必须正确设置")
//	ErrApiKeyPinningServicePermissionMustSet = NewBizError(errApiKeyPinningServicePermissionMustSet, "PinningServiceAPI权限必须正确设置", "PinningServiceAPI权限必须正确设置")
//)

//const (
//	//收藏
//	errFavoriteUnbound int = iota + 100401
//	errFavoriteBoundFailed
//
//	//上传
//	//errUploadFileEmpty
//	errUploadFilePdfTooLarge
//	errUploadFileThumbnailTooLarge
//	errUploadFileToLocal
//	errUploadFileIpfsPathEmpty
//	errUploadFileIpfsPathIllegal
//)
//
//// 校验字段
//const (
//	errVerifyTitle int = iota + 100501
//	errVerifyPublishDescription
//	errVerifyBio
//	errVerifyNickName
//	errVerifyArticleAbstract
//	errVerifyBountyDetail
//	errVerifyWalletAddrFormat
//	errVerifyArticleContents
//	errVerifyAvatarEmpty
//	errVerifyAvatarTooLarge
//	errVerifyAvatarType
//)
//
//// system
//const (
//	errSysUuidGenFailed = iota + 100601
//	errSysUuidNotEmpty
//	errSysUuidIdempotent
//)

//var (
//	// ipfs
//	//ErrUploadFileEmpty             = NewBizError(errUploadFileEmpty, "参数无效，文件内容不能为空", "Invalid parameter,file content cannot be empty")
//	ErrUploadFilePdfTooLarge       = NewBizError(errUploadFilePdfTooLarge, "上传pdf文件过大,不能超过20M", "The uploaded pdf file is too large and cannot exceed 20M")
//	ErrUploadFileThumbnailTooLarge = NewBizError(errUploadFileThumbnailTooLarge, "上传缩图文件过大,不能超过5M", "The uploaded thumbnail file is too large and cannot exceed 5M")
//	ErrUploadFileToLocal           = NewBizError(errUploadFileToLocal, "本地上传文件失败", "Failed to upload files locally")
//	ErrUploadFileIpfsPathEmpty     = NewBizError(errUploadFileIpfsPathEmpty, "Ipfs路径不能为空", "Ipfs path cannot be empty")
//	ErrUploadFileIpfsPathIllegal   = NewBizError(errUploadFileIpfsPathIllegal, "Ipfs路径格式错误,必须是 /ipfs/${cid}?filename=xxx.xx", "The format of the Ipfs path is incorrect. It must be/ipfs/${cid}? filename=xxx.xx")
//
//	//校验字段
//	ErrVerifyTitle              = NewBizError(errVerifyTitle, "验证标题最多80字符", "Validation title can be up to 80 characters")
//	ErrVerifyPublishDescription = NewBizError(errVerifyPublishDescription, "期刊描述最多300字符", "The publication description can be 300 characters at most")
//	ErrVerifyBio                = NewBizError(errVerifyBio, "个人简介最多150字符", "User profle can be up to 150 characters")
//	ErrVerifyNickName           = NewBizError(errVerifyNickName, "用户昵称最多50字符", "User nickname can be up to 50 characters")
//	ErrVerifyArticleAbstract    = NewBizError(errVerifyArticleAbstract, "文章描述最多600字符", "Article description up to 600 characters")
//	ErrVerifyBountyDetail       = NewBizError(errVerifyBountyDetail, "征稿描述长度最多5000字符", "The length of bounty description is 5000 characters at most")
//	ErrVerifyWalletAddrFormat   = NewBizError(errVerifyWalletAddrFormat, "钱包地址格式错误", "Wallet address format error")
//	ErrVerifyArticleContents    = NewBizError(errVerifyArticleContents, "文章内容最多20000字符", "Article contents up to 20000 characters")
//	ErrVerifyAvatarEmpty        = NewBizError(errVerifyAvatarEmpty, "头像文件不能为空", "The avatar file cannot be empty")
//	ErrVerifyAvatarTooLarge     = NewBizError(errVerifyAvatarTooLarge, "头像文件过大,不能超过2M", "The avatar file is too large to exceed 2M")
//	ErrVerifyAvatarType         = NewBizError(errVerifyAvatarType, "上传的头像图片类型不支持", "The type of avatar image uploaded is not supported.")
//
//	//sys
//	ErrSysUuidGenFailed  = NewBizError(errSysUuidGenFailed, "生成UUID失败,请重新生成", "Failed to generate UUID, please regenerate")
//	ErrSysUuidNotEmpty   = NewBizError(errSysUuidNotEmpty, "UUID必须不为空", "Header uuid must not be empty")
//	ErrSysUuidIdempotent = NewBizError(errSysUuidIdempotent, "接口幂等,不做处理", "Interface idempotent, no processing")
//)
