package hashgenerator

import (
	"crypto/md5"
	"encoding/hex"
)

type MD5HashGenerator struct {
}

func (md5Hash *MD5HashGenerator) GenerateHash(text string) string {
	hash := md5.Sum([]byte(text))
	encodedData := hex.EncodeToString(hash[:])
	return encodedData
}
