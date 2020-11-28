package util_test

import (
	"testing"

	"github.com/murer/desolation/util"
	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	//str := "dBU/s+c="
	str := "dBU_s-c"
	//assert.Equal(t, 0, len(str)%4)
	decoded := util.B64Dec(str)
	assert.Equal(t, str, util.B64Enc(decoded))
}
