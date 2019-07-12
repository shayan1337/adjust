package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

type Hash struct {}

func NewHash() Hash{
	return Hash{}
}

func (this Hash)Hash(bytes []byte) string {
	hasher := md5.New()
	hasher.Write(bytes)
	return hex.EncodeToString(hasher.Sum(nil))
}
