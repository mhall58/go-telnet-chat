package main

import (
	"log"
	"net"
)

type Server struct {
	Addr string
}

func (s Server) Start() {

	log.Println("starting server")

	listen, err := net.Listen("tcp", s.Addr)

	if err != nil {
		panic(err)
	}

	rooms := make(chan DataEvent)

	// This for loop accepts new incoming connections and assigns them to a new goroutine
	for {
		connection, listenErr := listen.Accept()

		if listenErr != nil {
			panic(listenErr)
		}
		go SessionHandler{}.handleSession(connection, rooms)

	}

}
