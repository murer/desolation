package util_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/murer/desolation/util"
	"github.com/stretchr/testify/assert"
)

func TestChannelReader(t *testing.T) {
	r := bytes.NewReader([]byte("0123456789"))
	q := util.ChannelReader(r, 4)

	log.Printf("f1")
	n := q.WaitShift()
	block, ok := n.([]byte)
	assert.Equal(t, []byte{0x30, 0x31, 0x32, 0x33}, block)
	assert.Equal(t, true, ok)

	log.Printf("f2")
	n = q.WaitShift()
	block, ok = n.([]byte)
	assert.Equal(t, []byte{0x34, 0x35, 0x36, 0x37}, block)
	assert.Equal(t, true, ok)

	log.Printf("f3")
	n = q.WaitShift()
	block, ok = n.([]byte)
	assert.Equal(t, []byte{0x38, 0x39}, block)
	assert.Equal(t, true, ok)

	log.Printf("f4")
	n = q.WaitShift()
	closed, ok := n.(string)
	assert.Equal(t, "closed", closed)
	assert.Equal(t, true, ok)

	log.Printf("f5")
	n = q.Shift()
	assert.Nil(t, n)
}
