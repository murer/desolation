package util

import (
	"hash/crc32"
)

func Crc32(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}
