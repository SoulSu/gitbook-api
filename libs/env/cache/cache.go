package cache

import (
	"errors"
	"time"
)

var (
	ExpiredErr = errors.New("Data expired")
)

type BookCacher interface {
	// init cache
	Init() error
	// get from cache by name
	// if data get but expired will return ExpireErr
	GetIfExists(name string, expireTime time.Time) ([]byte, error)
	// save data to cache
	Save(name string, data []byte) error
}
