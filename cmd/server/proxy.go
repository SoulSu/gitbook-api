package server

import (
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/SoulSu/gitbook-api/libs/env/cache"

	"github.com/SoulSu/gitbook-api/libs/log"
)

const API_ROOT = "https://api.gitbook.com/"

type Proxy struct {
	proxy     *httputil.ReverseProxy
	dataCache cache.BookCacher
}

func NewProxy(cacher cache.BookCacher) (http.Handler, error) {
	remote, err := url.Parse(API_ROOT)
	if err != nil {
		log.Fatalf("parse url err:%s", err.Error())
	}
	proxy := &Proxy{
		proxy:     httputil.NewSingleHostReverseProxy(remote),
		dataCache: cacher,
	}
	proxy.proxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	director := proxy.proxy.Director
	proxy.proxy.Director = func(r *http.Request) {
		director(r)
		r.Host = remote.Host
	}
	proxy.proxy.ModifyResponse = func(response *http.Response) error {
		return nil
	}

	return proxy, proxy.dataCache.Init()
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.proxy.ServeHTTP(w, r)
}
