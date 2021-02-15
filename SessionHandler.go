package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"regexp"
	"strings"
	"time"
)

type SessionHandler struct {
}

func (session SessionHandler) handleSession(connection net.Conn, rooms chan DataEvent) {

	read, write := session.getReaderAndWriter(connection)

	write.PrintfLine("Welcome to Simple Go Chat")
	session.writeHR(write)

	handle := session.prompt(read, write, "What is your handle?")

	eb.Subscribe("general", rooms)

	currentRoom := session.joinRoom("general", handle, write)
	go session.monitorRoom(rooms, currentRoom, write)

	//Input Detection
	for {
		input, _ := read.ReadLine()
		input = session.cleanInput(input)

		if input == "/exit" {
			connection.Close()
			break
		}

		log.Println(fmt.Sprintf("#%s - %s - %s", currentRoom, handle, input))
		eb.Publish(currentRoom, fmt.Sprintf("%s<<%s>> %s", session.getDate(), handle, input))
	}

}

func (SessionHandler) getReaderAndWriter(connection net.Conn) (*textproto.Reader, *textproto.Writer) {
	bufReader := bufio.NewReader(connection)
	read := textproto.NewReader(bufReader)
	bufWriter := bufio.NewWriter(connection)
	write := textproto.NewWriter(bufWriter)

	return read, write
}

func (SessionHandler) writeHR(write *textproto.Writer) {
	write.PrintfLine("---------------------------------------------------------")
}

func (SessionHandler) prompt(read *textproto.Reader, write *textproto.Writer, question string) string {
	answer := ""

	for answer == "" {
		write.PrintfLine(question)
		answer, _ = read.ReadLine()
	}

	return answer
}

func (session SessionHandler) joinRoom(roomName string, handle string, write *textproto.Writer) string {
	write.PrintfLine("Hi %s, welcome to #%s:", handle, roomName)
	eb.Publish(roomName, fmt.Sprintf("%s has joined #%s", handle, roomName))
	session.writeHR(write)
	return roomName
}
func (session SessionHandler) getDate() string {
	return time.Now().Format("01/02/06 3:04:05 PM MST")
}

func (session SessionHandler) cleanInput(input string) string {
	pattern, _ := regexp.Compile(`[^[:alpha:][:punct:][:space:]]`)
	input = pattern.ReplaceAllString(input, "")
	return strings.Trim(input, "' ")
}

func (session SessionHandler) monitorRoom(rooms chan DataEvent, room string, write *textproto.Writer) {
	for {
		input := <-rooms
		if input.Topic == room {
			write.PrintfLine(input.Data)
		}
	}
}
