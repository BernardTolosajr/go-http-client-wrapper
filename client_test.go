package serviceprovider

import (
	"net"
	"net/http"
	"net/http/httptest"
)

var (
	HOST = "http://localhost:3000/"
)

func server(fn func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	const HOST = "127.0.0.1:3000"

	ts := httptest.NewUnstartedServer(http.HandlerFunc(fn))

	l, _ := net.Listen("tcp", HOST)
	ts.Listener = l

	ts.Start()

	return ts
}
