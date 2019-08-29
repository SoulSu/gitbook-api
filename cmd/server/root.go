package server

import (
	"net/http"
)

type SrvHandle struct {
	*http.Server
}

func RunServer(port string) (err error) {
	srvHandle := SrvHandle{
		Server: new(http.Server),
	}
	srvHandle.Addr = ":" + port
	srvHandle.Handler, err = NewProxy(nil)
	return srvHandle.ListenAndServe()
}
