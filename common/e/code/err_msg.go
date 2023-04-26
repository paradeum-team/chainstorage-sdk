package code



var ScodeMsgMapping = map[int]string{
	ScodeOK:             "",
	ScodeErr:            "server error",
	ScodeRnErr:          "retrieval error",
	ScodeApiParamErr:    "invalid parameter",
	ScodeSysErr:         "server os/io error",
	ScodeNotExist:       "file does not exist",
	ScodeTrackingErr:    "genereal tracking error，please check trackRoots config...",
	ScodeTnErr:          "track error",
	ScodeInvalidRenewal: "cannot submit a file again on its creation day or renew expiration on the last day",

	//网络连接异常
	ScodeNetConnectionError: "network connection exception",
	ScodeRnNetConnectionError: "connect  RN Server network error",
	ScodeTnNetConnectionError: "connect TN Server network error",

	//服务器代码
	ScodeAFSErrNull:              "AFS null output",
	ScodeAFSErr:                  "AFS error",
	ScodeAPIErrParam:             "invalid parameter",
	ScodeAPIErrFile:              "invalid file",
	ScodeAFSErrExec:              "AFS core execution error",
	ScodeAFSTimeout:              "AFS core timed out",
	ScodeNotEnoughFund:           "not enough fund",
	ScodeErrOS:                   "server os error",
	ScodeErrIO:                   "server io error",
	ScodeOpOnNotExist:            "operation on non-existing file",
	ScodeInvalidContract:         "invalid signed contract",
	ScodeRNodesInfoErr:           "failed to fetch rnodes info",
	ScodeNotFoundFNDat:           "cannot locate fnode-daily data",
	ScodeNotReadyFNDat:           "fnode-daily data not ready",
	ScodeNoFNDat:                 "no fnode data available",
	ScodeCriticalErr:             "server preparing broadcast request failed after payment",
	ScodeCriticalAFSTimeout:      "AFS core timed out after payment",
	ScodeCriticalAFCErr:          "broadcast request to AFC failed after payment",
	ScodeCriticalTrackErr:        "failed to query rnode info from track node for extension synchronization",
	ScodeCriticalTrackTimeout:    "timed out querying track node for extension synchonization",
	ScodeCriticalRnodeSyncExt:    "failed to synchronize extension on rnode",
	ScodeAFCErr:                  "request to AFC failed",
	ScodeAFCTimeout:              "request to AFC timed out",
	ScodeContractVerificationErr: "contract verification error",
	ScodeReusingContract:         "contract already processed and consumed before",
	ScodeTrackErr:                "failed to request track node api",
	ScodeConfigErr:               "current server configuration does not support this request",



	//业务异常
	ScodeCopiesError:  "The number of BFS copies must be greater than or equal to the value set with ",
	ScodeCopiesExpirationTimeError: "BFS replica update expiration time failed",
	ScodeJsonError: "JSON conversion exception",
	ScodeIsTodyFileExpirationTimeError: "Today's files can't change the expiration time，Try again in 24 hours",
	ScodeDownloadRangeErr: "invalid range: failed to overlap",
	ScodeDbError:"leveldb is error",
	ScodeTimeStampExpiredInvalidError: "the time_stamp_expired param is invalid;Parse time Error. It must can be format eg:20060102150405 .",
	ScodeTimeStampExpiredInvalidLess3650Error: "Expiration time must be less than 3650 days",
	ScodeTimeStampExpiredInvalidLaterThanCurrentError: "The expired time must be later than current time",
	ScodeExpirationTimeError : "the expired time must be 3 days later than created time",

	ScodeAsymmetricEncryptionTypeError:"Encrypt type must be in [RSA,ECC].",
	ScodeAsymmetricEncryptionValueError:"Asympubkey is invalid",

}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := ScodeMsgMapping[code]
	if ok {
		return msg
	}

	return ScodeMsgMapping[ERROR]
}

func GetCodeByMsg(value string) int {
	for k, x := range ScodeMsgMapping {
		if x == value {
			return k
		}
	}
	return -1
}