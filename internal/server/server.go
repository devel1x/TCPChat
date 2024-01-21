package server

import (
	"fmt"
	"net"
	"net-cat/internal/handler"
	"strings"
	"sync"
)

func (s *server) Start() error {
	ln, err := net.Listen("tcp", s.port)
	if err != nil {
		if strings.Contains(err.Error(), "unknown port") {
			return fmt.Errorf("port should be a number, %w", err)
		}
		if strings.Contains(err.Error(), "permission denied") {
			return fmt.Errorf("port should be a four digit number, %w", err)
		}
		return nil
	}

	defer ln.Close()
	s.ln = ln
	fmt.Printf("Listening on port %v\n", s.ln.Addr())
	s.Serve()
	return nil
}

func (s *server) Serve() {
	var mu sync.Mutex
	for {
		con, err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			return
		}
		u, ok := handler.NewUser(&mu, con, s.ver)
		if !ok {
			con.Write([]byte("The chat is full"))
			con.Close()
		} else {
			go u.AcceptLoop()
		}

	}
}
