package util_test

import (
	"testing"
	"time"

	"github.com/murer/desolation/util"
	"github.com/stretchr/testify/assert"
)

func TestTimeRetry(t *testing.T) {
	timeRetry := util.TimeRetryCreate(1)
	assert.False(t, timeRetry.Expired())
	time.Sleep(100 * time.Millisecond)
	assert.False(t, timeRetry.Expired())
	time.Sleep(1000 * time.Millisecond)
	assert.True(t, timeRetry.Expired())
}
