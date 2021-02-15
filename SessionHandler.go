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

	// This next line is what picks up anything added to the currentRoom by this or any other users.
	go session.monitorRoom(rooms, currentRoom, write)

	//this loop detects input by the current connected user and publishes it to the rooms channel for the currentRoom
	for {
		input, _ := read.ReadLine()
		input = session.cleanInput(input)

		//Commands Area Break or Continue
		if input == "/exit" {
			connection.Close()
			break
		}

		// todo: this is where we would implement multiple channels
		//	if strings.Contains(input, "/join") {
		//
		//	}

		log.Println(fmt.Sprintf("#%s - %s - %s", currentRoom, handle, input))

		eb.Publish(currentRoom, fmt.Sprintf("%s<<%s>> %s", session.getDate(), handle, input))
	}

}

// getReaderAndWriter returns a new textproto.Reader and textproto.Writer for the given connection.
func (SessionHandler) getReaderAndWriter(connection net.Conn) (*textproto.Reader, *textproto.Writer) {
	bufReader := bufio.NewReader(connection)
	read := textproto.NewReader(bufReader)
	bufWriter := bufio.NewWriter(connection)
	write := textproto.NewWriter(bufWriter)

	return read, write
}

// writeHR outputs a horizontal line
func (SessionHandler) writeHR(write *textproto.Writer) {
	write.PrintfLine("---------------------------------------------------------")
}

// prompt asks the user a question and returns their input.
func (session SessionHandler) prompt(read *textproto.Reader, write *textproto.Writer, question string) string {
	answer := ""

	for answer == "" {
		write.PrintfLine(question)
		answer, _ = read.ReadLine()
		answer = session.cleanInput(answer)
	}

	return answer
}

// joinRoom alerts others that the handle has entered the room
func (session SessionHandler) joinRoom(roomName string, handle string, write *textproto.Writer) string {
	write.PrintfLine("Hi %s, welcome to #%s:", handle, roomName)
	eb.Publish(roomName, fmt.Sprintf("%s has entered #%s", handle, roomName))
	session.writeHR(write)
	return roomName
}

// getDate returns a the date and time in a human readable format
func (session SessionHandler) getDate() string {
	return time.Now().Format("01/02/06 3:04:05 PM MST")
}

// cleanInput scrubs the string of dirty characters.
func (session SessionHandler) cleanInput(input string) string {
	pattern, _ := regexp.Compile(`[^[:alpha:][:punct:][:space:]]`)
	input = pattern.ReplaceAllString(input, "")
	return strings.Trim(input, "' ")
}

// infinitely loops and checks the rooms channel for input and writes it to the screen if it's for the given room.
func (session SessionHandler) monitorRoom(rooms chan DataEvent, room string, write *textproto.Writer) {
	for {
		input := <-rooms
		if input.Topic == room {
			write.PrintfLine(input.Data)
		}
	}
}
