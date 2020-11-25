package util

import "encoding/base64"

func B64Enc(data []byte) string {
	if data == nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(data)
}

func B64Dec(data string) []byte {
	ret, err := base64.StdEncoding.DecodeString(data)
	Check(err)
	return ret
}
