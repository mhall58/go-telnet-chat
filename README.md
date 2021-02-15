# Go Telnet App
## Description

This is a telnet chat prototype written in go

The server listens to a port which is defined in the .env file

**.env**
```
GO_CHAT_ADDR=":5555"
GO_CHAT_LOG_FILE="./var/logs/gochat.log"
```

The goal of this project is to build a simple chat server in Go (or any language you feel comfortable with).
Multiple clients should be able to connect via telnet and send messages to the server.
When a message is sent to the server it should be relayed to all the connected clients (including a timestamp and the name of the client that sent the message).
All messages should also be logged to a local log file. Basic configuration settings like listening port, ip, and log file location should be read from a local config file.
