package webserver

import (
	"net/http"

	"net"
)

type WebServer struct {
	http.Server
	address string
}

func New (host, port string ) *WebServer {

		var ws WebServer

		ws.Addr = net.JoinHostPort(host,port)

		ws.Handler = h

		return &ws
}

func (s *WebServer) Start () error {
	return s.ListenAndServe()
}