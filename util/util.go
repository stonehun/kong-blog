package util

import (
	"crypto/md5"
	"encoding/hex"
)

func MdString(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
