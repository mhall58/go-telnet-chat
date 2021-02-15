package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	Addr string
}

func (server Server) Start() {

	log.Println(fmt.Sprintf("Starting server on %server", server.Addr))

	listen, err := net.Listen("tcp", server.Addr)

	if err != nil {
		panic(err)
	}

	//rooms is the channel which is used by the event bus and allows communication between goroutines.
	rooms := make(chan DataEvent)

	// This for loop accepts new incoming connections and assigns them to a new goroutine with access to the rooms channel
	for {
		connection, listenErr := listen.Accept()

		if listenErr != nil {
			panic(listenErr)
		}
		go SessionHandler{}.handleSession(connection, rooms)

	}

}
