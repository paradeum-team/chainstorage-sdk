package dgst

import "strings"

const AfsAlgoTypeMd5 = "md5"
const AfsAlgoTypeSha1 = "sha1"
const AfsAlgoTypeSha256 = "sha256"
const AfsAlgoTypeSha512 = "sha512"
const AfsAlgoTypeFileLength = "file_length"
const AfsAlgoTypeAfId = "afid"
const AfsAlgoTypeAfIdMini = "afid_mini"
const AfsAlgoTypeAfIdLite = "afid_lite"
const AfsAlgoTypeAfidSeed = "seed_afid_1e"
const AfsAlgoTypeAfidRaw = "raw_afid_1e"
const AfsAlgoTypeGfid = "gfid"

func GetDataTypeByDgst(dgst string) string {
	dataType := ""
	dataLength := len(dgst)

	switch dataLength {
	case 34:
		//gfid
		dataType = AfsAlgoTypeGfid
	case 32:
		//md5
		dataType = AfsAlgoTypeMd5
	case 24:
		//afid_mini
		dataType = AfsAlgoTypeAfIdMini
	case 56:
		//afid_lite
		dataType = AfsAlgoTypeAfIdLite
	case 40:
		//sha1
		dataType = AfsAlgoTypeSha1
	case 64:
		//sha256
		dataType = AfsAlgoTypeSha256
	case 128: //afid & sha512
		if strings.Contains(dgst, "0000") {
			//afid
			dataType = AfsAlgoTypeAfId
		} else {
			//sha512
			dataType = AfsAlgoTypeSha512
		}
	default:
		dataType = ""
	}

	return dataType
}

/**
 * 获取数据类型根据摘要信息（IPFS）
 */
func GetDataTypeByDgst4Ipfs(dgst string) string {
	dataType := ""
	dataLength := len(dgst)

	switch dataLength {
	case 32:
		//md5
		dataType = AfsAlgoTypeMd5
	case 40:
		//sha1
		dataType = AfsAlgoTypeSha1
	case 128: //afid & sha512
		if strings.Contains(dgst, "0000") {
			//afid
			dataType = AfsAlgoTypeAfId
		}
	default:
		dataType = ""
	}

	return dataType
}
