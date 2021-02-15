package main

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"regexp"
	"strings"
	"time"
)

func main() {

	s := Server{Addr: ":5555"}

	l, _ := s.Start()
	//defer l.Close()

	rooms := make(chan DataEvent)

	//this for loop accepts new connections.
	for {
		c, _ := l.Accept()

		//we pass the connection off to a goroutine.
		//SESSION Handler
		go func(c net.Conn, rooms chan DataEvent) {
			eb.Subscribe("general", rooms)
			bufReader := bufio.NewReader(c)
			read := textproto.NewReader(bufReader)
			bufWriter := bufio.NewWriter(c)
			write := textproto.NewWriter(bufWriter)
			write.PrintfLine("Welcome to Simple Go Chat")
			write.PrintfLine("-----------------------------------")

			handle := ""

			for handle == "" {
				write.PrintfLine("Type a handle :")
				handle, _ = read.ReadLine()
			}

			write.PrintfLine("Hi %s, welcome to #general:", handle)
			write.PrintfLine("-----------------------------------")

			// Channel Listener
			go func(rooms chan DataEvent) {
				for {
					input := <-rooms
					write.PrintfLine(input.Data)
				}
			}(rooms)

			//Input Detection
			for {
				input, _ := read.ReadLine()
				input = cleanInput(input)

				fmt.Println(input)

				if input == "exit" {
					c.Close()
					break
				}

				eb.Publish("general", fmt.Sprintf("%s<<%s>> %s", getDate(), handle, input))
			}

		}(c, rooms)

	}

}
func getDate() string {
	return time.Now().Format("01/02/06 3:04:05 PM MST")
}

func cleanInput(input string) string {

	pattern, _ := regexp.Compile(`[^[:alpha:][:punct:][:space:]]`)
	//there always seems to be 1 extra quote character left over in the beginning
	input = strings.Replace(input, "'", "", 1)
	input = strings.Replace(input, " ", "", 1)
	input = textproto.TrimString(input)
	return pattern.ReplaceAllString(input, "")
}
