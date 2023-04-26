package dgstseed

import (
	"bytes"
	"chainstorage-sdk/entity"
	"chainstorage-sdk/utils"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	sys_path "path"
	"path/filepath"
)

/**
 * 把 seed file 解析为 seed 描述对象
 */
func ToSeedObjoct(reader io.Reader) (seeds *entity.SeedDescriptioFile, err error) {
	seeds = &entity.SeedDescriptioFile{}
	data, err := ioutil.ReadAll(reader)
	/*afidsLen := len(data)
	if afidsLen%64 != 0 {
		return seeds, errors.New("Not seedfile format")
	}*/
	if isUtf8(data) {
		err := json.Unmarshal(data, seeds)
		if err != nil {
			return seeds, err
		}
	} else {
		for i := 0; i < len(data); i += 64 {
			nn := i + 64
			afid := hex.EncodeToString(data[i:nn])
			dgstObj := ConvertAfig2Dgst(afid)
			seed := entity.Seed{
				Sequence: uint64(i / 64),
				Afid:     afid,
				Size:     dgstObj.FileSize,
			}
			seeds.Seeds = append(seeds.Seeds, seed)
		}
		return seeds, nil
	}
	return seeds, nil
}

/**
 * 拆分大文件，以及转换成seed描述文件
 */
func SplitBigFile(file_path string) (splitPath string, seeds entity.SeedDescriptioFile, err error) {

	if utils.IsDir(file_path) {
		return
	}
	fileDir := sys_path.Dir(file_path)
	fileName := sys_path.Base(file_path)
	splitPath, err = ioutil.TempDir(fileDir, fileName+".seed.")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("splitPath: %s\n", splitPath)
	log.Printf("upload filePath: %s\n", file_path)

	fileSize := utils.GetFileSize(file_path)

	// calculate total number of parts the filePath will be chunked into
	var fileChunk uint64 = 10 << 10 << 10
	log.Printf("fileChunk: %d\n", fileChunk)
	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	totalAfid, err := GetAfidLocal(file_path)
	if err != nil {
		return splitPath, seeds, err
	}
	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)
	file, err := os.Open(file_path)
	if err != nil {
		return splitPath, seeds, err
	}
	var seedContent []byte

	for i := uint64(0); i < totalPartsNum; i++ {
		//获取最小容量
		partSize := uint64(math.Min(float64(fileChunk), float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)
		file.Read(partBuffer)
		afid := Parse2Afid(partBuffer)

		fileName := filepath.Join(splitPath, afid)

		_, err = os.Create(fileName)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		// write/save buffer to disk
		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
		seed := entity.Seed{
			Sequence: i,
			Afid:     afid,
			Size:     partSize,
		}
		seeds.Seeds = append(seeds.Seeds, seed)

		seedHex, err := hex.DecodeString(fmt.Sprintf("%s", afid))
		if err != nil {
			log.Fatal(err)
		}
		seedContent = BytesCombine(seedContent, seedHex)
	}

	seedPath := filepath.Join(splitPath, "afids.afs")

	_, err = os.Create(seedPath)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// write/save buffer to disk
	ioutil.WriteFile(fileName, seedContent, os.ModeAppend)

	seeds.Afid = totalAfid
	return splitPath, seeds, err
}
func preNUm(data byte) int {
	var mask byte = 0x80
	var num int = 0
	//8bit中首个0bit前有多少个1bits
	for i := 0; i < 8; i++ {
		if (data & mask) == mask {
			num++
			mask = mask >> 1
		} else {
			break
		}
	}
	return num
}
func isUtf8(data []byte) bool {
	i := 0
	for i < len(data) {
		if (data[i] & 0x80) == 0x00 {
			// 0XXX_XXXX
			i++
			continue
		} else if num := preNUm(data[i]); num > 2 {
			// 110X_XXXX 10XX_XXXX
			// 1110_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
			// preNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
			i++
			for j := 0; j < num-1; j++ {
				//判断后面的 num - 1 个字节是不是都是10开头
				if (data[i] & 0xc0) != 0x80 {
					return false
				}
				i++
			}
		} else {
			//其他情况说明不是utf-8
			return false
		}
	}
	return true
}

// 二进制拼接
func BytesCombine(pBytes ...[]byte) []byte {
	len := len(pBytes)
	s := make([][]byte, len)
	for index := 0; index < len; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

// 4 12
// 1e00 000000a00000 f44dcd0835097973865cd86ebf486ccd11875d50c8d5619ea6c97095523bbfb168c82823182f77c3c11e947e1e04eb31c1791de3e31b9fe1
// 1e00 000000a00000 f43447967f203a80c4d6d3471498fd84b153c73d50af3663a5829549f946d5804a4f9448d09cc0c28e88cacdb5c5bd06446e4b63deefab3d
func Parse2Afid(data []byte) string {
	const bfsVersionCode = "1e"
	verisonCode := fmt.Sprintf("%s00", bfsVersionCode)
	fileLength := len(data)

	fileLengthHex := fmt.Sprintf("%0*s", 12, fmt.Sprintf("%x", fileLength))
	h := sha1.New()
	h.Write(data)
	fileSha := fmt.Sprintf("%x", h.Sum(nil))
	m := md5.New()
	m.Write(data)
	fileMd5 := fmt.Sprintf("%x", m.Sum(nil))

	str88 := verisonCode + fileLengthHex + fileSha + fileMd5

	h2 := sha1.New()
	h2.Write([]byte(str88))
	str88Sha := fmt.Sprintf("%x", h2.Sum(nil))

	afid := str88 + str88Sha
	return afid
}
