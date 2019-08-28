package cache

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestDiskCache_Init(t *testing.T) {
	type fields struct {
		cacheDir string
	}
	existsDir, _ := ioutil.TempDir("", "cache")

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "not exists dir",
			fields:  fields{cacheDir: fmt.Sprintf("/tmp/cache.%d", rand.Int())},
			wantErr: false,
		},
		{
			name:    "exists dir",
			fields:  fields{cacheDir: existsDir},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DiskCache{
				cacheDir: tt.fields.cacheDir,
			}
			if err := d.Init(); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
