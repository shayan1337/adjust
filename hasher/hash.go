package hasher

import (
	"adjust/app"
	"crypto/md5"
	"encoding/hex"
)

type Hash struct {}

func NewHash() app.Hasher {
	return Hash{}
}

func (this Hash)Hash(bytes []byte) string {
	hasher := md5.New()
	hasher.Write(bytes)
	return hex.EncodeToString(hasher.Sum(nil))
}
