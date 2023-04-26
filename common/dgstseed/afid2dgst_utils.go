package dgstseed

import (
	"chainstorage-sdk/entity"
	"chainstorage-sdk/utils"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strconv"
)

func GetAfidLocal(path string) (rawAfid string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	verisonCode := fmt.Sprintf("%s00", "1e")
	fileLength := utils.GetFileSize(path)
	fileLengthHex := fmt.Sprintf("%0*s", 12, fmt.Sprintf("%x", fileLength))
	h := sha1.New()
	io.Copy(h, f)
	fileSha := fmt.Sprintf("%x", h.Sum(nil))

	f.Seek(0, 0) //公用同一个f指针；重置指针。

	m := md5.New()
	io.Copy(m, f)
	fileMd5 := fmt.Sprintf("%x", m.Sum(nil))
	str88 := verisonCode + fileLengthHex + fileSha + fileMd5
	h2 := sha1.New()
	h2.Write([]byte(str88))
	str88Sha := fmt.Sprintf("%x", h2.Sum(nil))

	afid := str88 + str88Sha
	utils.DevLog("GetAfidLocal", "本地计算afid："+afid)
	return afid, nil
}

/**
 * 从afid 中解析出基本的dgst 数值
 */
func ConvertAfig2Dgst(afid string) entity.AfsSimpleDgst {
	//utils.LogDebug("ConvertAfig2Dgst",fmt.Sprintf("afid.len=%d \n", len(afid)))

	prefix := afid[0:4]
	fileLengthHex := afid[4:16]
	sha1 := afid[16:56]
	md5 := afid[56:88]
	suffix := afid[88:]

	/*utils.LogDebug("ConvertAfig2Dgst",fmt.Sprintf("prefix=%s,length=%d \n ", prefix, len(prefix)))
	utils.LogDebug("ConvertAfig2Dgst",fmt.Sprintf("fileLengthHex=%s ,length=%d\n ", fileLengthHex, len(fileLengthHex)))
	utils.LogDebug("ConvertAfig2Dgst",fmt.Sprintf("sha1=%s ,length=%d\n ", sha1, len(sha1)))
	utils.LogDebug("ConvertAfig2Dgst",fmt.Sprintf("md5=%s ,length=%d\n ", md5, len(md5)))
	utils.LogDebug("ConvertAfig2Dgst",fmt.Sprintf("suffix=%s,length=%d \n ", suffix, len(suffix)))*/

	//计算其他数值
	//afid_mini是前4位加最後20位，共24字节

	afidMini := prefix + suffix[20:]
	//afid_lite是前16位加最後40位，共56字节
	afidLite := prefix + fileLengthHex + suffix
	dgst := entity.AfsSimpleDgst{Afid: afid, AfidLite: afidLite, AfidMini: afidMini, Md5: md5, Sha1: sha1, FileLengthHex: fileLengthHex}

	size, err := strconv.ParseUint(fileLengthHex, 16, 64)

	if err != nil {
		utils.LogError("get file size fail ... ")
	} else {
		dgst.FileSize = size
	}

	return dgst
}
