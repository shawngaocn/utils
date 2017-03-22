package strutils

import (
	"crypto/md5"
	crand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	mrand "math/rand"
)

func NonceString() string {
	var bytes [32]byte
	var b int64
	for i := 0; i != 32; i++ {
		b = mrand.Int63n(62)
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

func MD5Bytes(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

// GUID Make MD5 GUID
func GUID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}
	return MD5String(base64.URLEncoding.EncodeToString(b))
}

// GUIDX Make GUID with NonceString
func GUIDX() string {
	str := NonceString()
	b := make([]byte, 48)
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}
	return MD5String(base64.URLEncoding.EncodeToString([]byte(string(b) + str)))
}
