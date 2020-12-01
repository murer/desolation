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
	cr := util.ChannelReader(r, 4)

	block := <-cr
	assert.Equal(t, []byte{0x30, 0x31, 0x32, 0x33}, block.Data)
	assert.Equal(t, true, block.Open)

	block = <-cr
	assert.Equal(t, []byte{0x34, 0x35, 0x36, 0x37}, block.Data)
	assert.Equal(t, true, block.Open)

	block = <-cr
	assert.Equal(t, []byte{0x38, 0x39}, block.Data)
	assert.Equal(t, true, block.Open)

	block = <-cr
	assert.Nil(t, block.Data)
	assert.Equal(t, false, block.Open)

	block = <-cr
	assert.Nil(t, block.Data)
	assert.Equal(t, false, block.Open)
}

func TestChannelReaderFor(t *testing.T) {
	r := bytes.NewReader([]byte("0123456789"))
	cr := util.ChannelReader(r, 4)
	for elem := range cr {
		log.Printf("x: %v", elem)
	}
	block := <-cr
	assert.Nil(t, block.Data)
	assert.Equal(t, false, block.Open)

	block = <-cr
	assert.Nil(t, block.Data)
	assert.Equal(t, false, block.Open)
}
