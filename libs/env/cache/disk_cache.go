package cache

import (
	"crypto/md5"
	"encoding/hex"
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

// find save file by request name
func (d *DiskCache) findFileByName(name string) (*os.File, error) {

	fileNameFunc := func(name string) string {
		md5Hash := md5.New()
		md5Hash.Write([]byte(name))
		return hex.EncodeToString(md5Hash.Sum(nil))
	}

	fileName := fileNameFunc(name)
	_ = fileName

	return nil, nil
}
