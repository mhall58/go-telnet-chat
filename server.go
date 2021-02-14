package main

import "net"

type Server struct {
	Addr string
}

func (s Server) Start() (net.Listener, error) {
	return net.Listen("tcp", s.Addr)
}