package code

const ScodeOK = 0
const ERROR          = 500

//-------------------------------------------
const ScodeErr = 10000 //未定义错误
const ScodeSysErr = 10001
const ScodeApiParamErr = 10002
const ScodeTnErr = 10003
const ScodeRnErr = 10004
const ScodeTrackingErr = 10005

//----------执行成功，返回无效---
const ScodeInvalidRenewal = 20002
const ScodeNotExist = 20001


// 服务链接超时
const ScodeNetConnectionError = 50010
const ScodeRnNetConnectionError  = 50011
const ScodeTnNetConnectionError  = 50012


//-----------TN Exception
const(
	//服务器代码
	ScodeAPIErrParam     = 30004
	ScodeAPIErrFile      = 30005
	ScodeInvalidContract = 30013
	ScodeAFSErrNull      = 30002
	ScodeAFSErr          = 30003
	ScodeAFSErrExec      = 30006
	ScodeAFSTimeout      = 30007
	ScodeNotEnoughFund   = 30008
	ScodeOpOnNotExist    = 30010

	ScodeErrOS = 30011
	ScodeErrIO = 30012

	ScodeNotFoundFNDat = 30051
	ScodeNotReadyFNDat = 30052
	ScodeNoFNDat       = 30053

	ScodeAFCErr                  = 30061
	ScodeAFCTimeout              = 30063
	ScodeContractVerificationErr = 30064
	ScodeReusingContract         = 30065

	ScodeTrackErr      = 30087
	ScodeRNodesInfoErr = 30088

	ScodeConfigErr = 30099

	ScodeCriticalErr        = 39021
	ScodeCriticalAFSErr     = 39022
	ScodeCriticalAFSTimeout = 39023

	ScodeCriticalAFCErr       = 39031
	ScodeCriticalRnodeSyncExt = 39036

	ScodeCriticalTrackErr     = 39054
	ScodeCriticalTrackTimeout = 39055



	//业务异常
	ScodeCopiesError = 40001
	ScodeCopiesExpirationTimeError = 40002
	ScodeExpirationTimeError = 40022

	ScodeJsonError = 40003
	ScodeIsTodyFileExpirationTimeError = 40004
	ScodeDownloadRangeErr = 40006
	ScodeDbError = 40007

	//过期时间系列
	ScodeTimeStampExpiredInvalidError = 40008
	ScodeTimeStampExpiredInvalidLess3650Error = 40009
	ScodeTimeStampExpiredInvalidLaterThanCurrentError = 40010

	// 非对称校验不过
	ScodeAsymmetricEncryptionTypeError = 40011
	ScodeAsymmetricEncryptionValueError = 40012


)

