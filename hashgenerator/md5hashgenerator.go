package hashgenerator

import (
	"crypto/md5"
	"encoding/hex"
)

// Implementation of HashGenerator interface (uses MD5 for hashing).
type MD5HashGenerator struct {
}

func (md5Hash *MD5HashGenerator) GenerateHash(text string) string {
	hash := md5.Sum([]byte(text))
	encodedData := hex.EncodeToString(hash[:])
	return encodedData
}
