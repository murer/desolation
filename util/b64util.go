package util

import "encoding/base64"

func B64Enc(data []byte) string {
	if data == nil {
		return ""
	}
	return base64.RawURLEncoding.EncodeToString(data)
}

func B64Dec(data string) []byte {
	ret, err := base64.RawURLEncoding.DecodeString(data)
	Check(err)
	return ret
}
