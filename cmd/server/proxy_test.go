package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SoulSu/gitbook-api/libs/env/cache"
)

func TestProxy_ServeHTTP(t *testing.T) {
	cacheDir := "/tmp/gitbook-api/cache"
	p, err := NewProxy(cache.NewDiskCache(cacheDir))
	if err != nil {
		t.Errorf(err.Error())
	}
	server := httptest.NewServer(p)
	reqUrl := server.URL + "/book/wizardforcel/vbird-linux-basic-4e"
	t.Log(reqUrl)
	ret, err := http.Get(reqUrl)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer ret.Body.Close()

	resp, _ := ioutil.ReadAll(ret.Body)
	t.Logf("response : %s", resp)
}
