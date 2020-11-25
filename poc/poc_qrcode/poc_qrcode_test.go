package poc_qrcode_test

import (
	"testing"

	"github.com/murer/desolation/poc/poc_qrcode"
	"github.com/stretchr/testify/assert"
)

func TestScreenshot(t *testing.T) {
	txt := poc_qrcode.Parse()
	assert.Equal(t, "My simple text\n", txt)
}
