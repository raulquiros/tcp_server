package server

import (
	"bufio"
	"context"
	"fmt"
	"github.com/raulquiros/tcp_server/internal/sku"
	"github.com/raulquiros/tcp_server/kit/command"
	"log"
	"net"
	"strings"
)

const maxConn = 5

var numConn int

// Server ...
type Server struct {
	host       string
	port       string
	commandBus command.Bus
}

// Config ...
type Config struct {
	Host string
	Port string
}

// New ...
func New(config *Config, commandBus command.Bus) *Server {
	return &Server{
		host:       config.Host,
		port:       config.Port,
		commandBus: commandBus,
	}
}

// Run ...
func (server *Server) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		if numConn < maxConn {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}

			go handleRequest(conn, server.commandBus)
			numConn++
		}
	}
}

func handleRequest(conn net.Conn, commandBus command.Bus) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			numConn--
			return
		}

		param := strings.TrimSpace(string(message))
		if param == "terminate" {
			conn.Close()
		}

		err = commandBus.Dispatch(context.Background(), sku.NewCreateSkuCommand(param))
		if err != nil {
			conn.Write([]byte(fmt.Sprintf("Error: %s\n", err)))
		}
		fmt.Printf("Message incoming: %s", string(message))
		conn.Write([]byte("---\n"))
	}
}
