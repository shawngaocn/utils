package auth

import (
	"encoding/json"

	"github.com/shawngaocn/utils/strutils"
)

type BasicAuth struct {
	Key       string `json:"key"`
	Secret    string `json:"-"`
	Nonce     string `json:"nonce"`
	Signature string `json:"sign"`
}

func (b *BasicAuth) Verify(key, secret string) bool {
	if len(b.Key) == 0 || len(b.Nonce) == 0 || len(b.Signature) == 0 {
		return false
	}

	if b.Key != key {
		return false
	}

	sign := strutils.MD5String(b.Key + b.Nonce + secret)
	if b.Signature != sign {
		return false
	}

	return true
}

func BasicAuthVerifyWithJSON(key, secret string, body []byte) (*BasicAuth, bool) {
	auth := &BasicAuth{}
	if len(body) == 0 {
		return nil, false
	}
	if err := json.Unmarshal(body, auth); err != nil {
		return nil, false
	}
	return auth, auth.Verify(key, secret)
}
