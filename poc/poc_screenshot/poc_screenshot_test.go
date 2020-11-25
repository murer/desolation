package poc_screenshot_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/murer/desolation/poc/poc_screenshot"
)

func TestScreenshot(t *testing.T) {
	out := poc_screenshot.Screenshot()
	assert.Equal(t, 1, 1)
	assert.Less(t, 1, len(out))

	os.MkdirAll("target", 0755)
	f, err := os.Create("target/poc_screenshot_test.png")
	assert.Nil(t, err)
	defer f.Close()
	f.Write(out)
}
