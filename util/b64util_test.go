package util_test

import (
	"testing"

	"github.com/murer/desolation/util"
	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	assert.Equal(t, "eyJOYW1lIjoiZWNobyIsIkhlYWRlcnMiOnsicmlkIjoiMiJ9LCJQYXlsb2FkIjoiY2hlY2t0ZXh0In0=", util.B64Enc(util.B64Dec("eyJOYW1lIjoiZWNobyIsIkhlYWRlcnMiOnsicmlkIjoiMiJ9LCJQYXlsb2FkIjoiY2hlY2t0ZXh0In0=")))
}
