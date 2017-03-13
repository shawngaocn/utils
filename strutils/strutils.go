package strutils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

func NonceString() string {
	var bytes [32]byte
	var b int64
	for i := 0; i != 32; i++ {
		b = rand.Int63n(62)
		switch {
		case b < 10:
			b += 48
		case b < 36:
			b += 55
		default:
			b += 61
		}
		bytes[i] = byte(b)
	}
	return string(bytes[:32])
}

func MD5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
