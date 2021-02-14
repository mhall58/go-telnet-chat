package main

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
)

func main() {

	s := Server{Addr: ":5555"}

	l, _ := s.Start()
	//defer l.Close()

	for {
		c, _ := l.Accept()

		//session handler
		go func(c net.Conn) {
			bufReader := bufio.NewReader(c)
			read := textproto.NewReader(bufReader)
			bufWriter := bufio.NewWriter(c)
			write := textproto.NewWriter(bufWriter)
			write.PrintfLine("Welcome to Simple Go Chat\r\n")


			for{
				input, _ := read.ReadLine()
				fmt.Println(input)
				if input == "exit" {
					break
				}
			}



		}(c)

	}
}
