package main

import (
	"fmt"
	"log"
	"main/cache"
	"net"
)

type ServerOpts struct {
	ListenAddr string
	IsLeader   bool
}

type Server struct {
	ServerOpts

	cache cache.Cacher
}

func NewServer(opts ServerOpts, c cache.Cacher) *Server {
	return &Server{
		ServerOpts: opts,
		cache:      c,
	}
}

// creates a TCP connection and starts the server
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return fmt.Errorf("listen error: %s\n", err)
	}

	log.Printf("server starting at port [%s]\n", "3000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %s\n", err)
			continue
		}

		go s.handleConn(conn)
	}
}

// the conn from the above function will be handled here
// using a goroutine

func (s *Server) handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
	}()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("read conn error: %s\n", err)
			break
		}

		go s.handleCommand(conn, buf[:n])

	}
}

func (s *Server) handleCommand(conn net.Conn, rawCmd []byte) {
	msg, err := parseMessage(rawCmd)
	if err != nil {
		fmt.Println("failed to parse commmand\n")
		return
	}

	switch msg.Cmd {
	case CMDSet:
    if err := s.handleSetCmd(conn, msg); err != nil {
      return
    }
	}
}

func (s *Server) handleSetCmd(conn net.Conn, msg *Message) error {
	fmt.Println("hanlding the set command: ", msg)
  
  return nil
}
