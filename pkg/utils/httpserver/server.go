package httpserver

import (
	"net"
	"net/http"
)

func Serve(serverAddress, protocol string, handler http.Handler) error {
	srv := http.Server{
		Addr: serverAddress,
	}

	srv.Handler = handler

	ln, err := net.Listen(protocol, srv.Addr)
	if err != nil {
		return err
	}

	return srv.Serve(ln)
}
