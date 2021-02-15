# Go Telnet App
## Description

This is a telnet chat prototype written in go.

This is my first go application ever.

The server listens to a port which is defined in the .env file
The server logs to a port which is defined in the .env file

**.env**
```
GO_CHAT_ADDR=":5555"
GO_CHAT_LOG_FILE="./var/logs/gochat.log"
```

# Approach

- Used the build in net library to listen to a given port
- used goroutines to handle sessions for each user
- used a chan and some borrowed event bus code to share messages between users
  - this could probably be simplified with some more time
  - I knew that i wanted a basic pub/sub pattern to solve the chat piece. Chat users are both publishers and subscribers in our case
  - event bus code: https://levelup.gitconnected.com/lets-write-a-simple-event-bus-in-go-79b9480d8997


## Current Limitations & Bugs
- the input should clear what you currently typed after hitting enter before it reprints it out, But i haven't found the correct combinations of ansi control sequences to get that to happen yet
- backspace isn't working yet if a user changes their input. Probably more ansi control sequences to get that to be supported

## ToDo
- multiple rooms
  - I feel the event bus could support multiple rooms with some mild changes. Topic == Room
  - I would need to add an unsubcribe method to the event bus when they leave one room and enter another
- REST POST
  - this is just a matter of a new listener and using the event bus to subscribe and publish data to the same chan
- REST GET (search)
 - this might be something that scans the log file for messages and returns matching results since we don't have a database.

