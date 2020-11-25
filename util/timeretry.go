package util

import (
	"log"
	"time"
)

type TimeRetry struct {
	Created int64
	Expires int64
}

func TimeRetryCreate(duration int64) *TimeRetry {
	now := time.Now().Unix()
	return &TimeRetry{
		Created: now,
		Expires: now + duration,
	}
}

func (me *TimeRetry) Expired() bool {
	now := time.Now().Unix()
	log.Printf("test: %d, %d", now, me.Expires)
	return now >= me.Expires
}
