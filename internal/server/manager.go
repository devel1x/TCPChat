package server

import (
	"flag"
	"net"
)

type server struct {
	port string
	ln   net.Listener
	ver  string
}

type ServerI interface {
	Start() error
	Serve()
}

func NewServer() ServerI {
	var port string
	var ver string
	flag.StringVar(&port, "port", "2525", "Usage: go run cmd/main.go [PORT] \n\nExample: go run cmd/main.go -port=8080")
	flag.StringVar(&ver, "ver", "classic", "Usage: go run cmd/main.go [VER] \n\nExample: go run cmd/main.go -ver=right")
	flag.Parse()
	port = "localhost:" + port

	return &server{
		port: port,
		ln:   nil,
		ver:  ver,
	}
}
