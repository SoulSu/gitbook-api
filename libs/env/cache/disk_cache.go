package cache

import (
	"os"
	"path/filepath"
	"time"
)

type DiskCache struct {
	cacheDir string
}

func NewDiskCache(cacheDir string) BookCacher {
	return &DiskCache{cacheDir: filepath.Dir(cacheDir)}
}

func (d *DiskCache) Init() error {
	// setup cache dir
	_, err := os.Stat(d.cacheDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(d.cacheDir, 0755)
	}
	if err != nil {
		return err
	}

	return nil
}

func (d *DiskCache) GetIfExists(name string, expireTime time.Time) ([]byte, error) {
	panic("implement me")
}

func (d *DiskCache) Save(name string, data []byte) error {
	panic("implement me")
}
