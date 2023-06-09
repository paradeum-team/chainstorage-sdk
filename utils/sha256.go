package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/pkg/errors"
	"io"
	"os"
)

func EncodeSha256(value string) string {
	m := sha256.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func GetFileSha256ByReader(reader io.Reader) (string, error) {
	fileSha256 := sha256.New()
	io.Copy(fileSha256, reader)
	return hex.EncodeToString(fileSha256.Sum(nil)), nil
}

func GetFileSha256ByPath(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", errors.Wrapf(err, "GetFileSha256ByPath open %s failed", filePath)
	}
	defer file.Close()

	fileSha256 := sha256.New()
	io.Copy(fileSha256, file)

	return hex.EncodeToString(fileSha256.Sum(nil)), nil
}
