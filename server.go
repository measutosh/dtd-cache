package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net"
	
//     "main/cache"
// )

// type ServerOpts struct {
// 	ListenAddr string
// 	IsLeader   bool
// 	LeaderAddr string
// }

// type Server struct {
// 	ServerOpts
// 	followers map[net.Conn]struct{}
// 	cache     cache.Cacher
// }

// func NewServer(opts ServerOpts, c cache.Cacher) *Server {
// 	return &Server{
// 		ServerOpts: opts,
// 		cache:      c,
// 	}
// }

// // creates a TCP connection and starts the server
// func (s *Server) Start() error {
// 	ln, err := net.Listen("tcp", s.ListenAddr)
// 	if err != nil {
// 		return fmt.Errorf("listen error: %s\n", err)
// 	}

// 	log.Printf("server starting at port [%s]\n", s.ListenAddr)

// 	// if it's not a leader we will dial the leader
// 	// will open up a connection with the leader to recieve info
// 	if !s.IsLeader {
// 		go func() {
// 			conn, err := net.Dial("tcp", s.LeaderAddr)
// 			fmt.Println("connected with leader: ", s.LeaderAddr)
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			s.handleConn(conn)
// 		}()
// 	}

// 	for {
// 		conn, err := ln.Accept()
// 		if err != nil {
// 			log.Printf("accept error: %s\n", err)
// 			continue
// 		}

// 		go s.handleConn(conn)
// 	}
// }

// // the conn from the above function will be handled here
// // using a goroutine

// func (s *Server) handleConn(conn net.Conn) {
// 	defer conn.Close()

// 	buf := make([]byte, 2048)

//   if s.IsLeader {
//     s.followers[conn] = struct{}{}
//   }
	
//   fmt.Println("connection made: ", conn.RemoteAddr())
//   for {
// 		n, err := conn.Read(buf)
// 		if err != nil {
// 			log.Printf("read conn error: %s\n", err)
// 			break
// 		}

// 		go s.handleCommand(conn, buf[:n])

// 	}
// }

// func (s *Server) handleCommand(conn net.Conn, rawCmd []byte) {
// 	msg, err := parseMessage(rawCmd)
// 	if err != nil {
// 		fmt.Println("failed to parse commmand: ", err)
// 		conn.Write([]byte(err.Error()))
// 		return
// 	}

// 	fmt.Printf("received command %s\n", msg.Cmd)

// 	switch msg.Cmd {
// 	case CMDSet:
// 		err = s.handleSetCmd(conn, msg)
// 	case CMDGet:
// 		err = s.handleGetCmd(conn, msg)
// 	}
// 	if err != nil {
// 		fmt.Println("failed to handle command: ", err)
// 		conn.Write([]byte(err.Error()))
// 	}
// }

// func (s *Server) handleGetCmd(conn net.Conn, msg *Message) error {
// 	val, err := s.cache.Get(msg.Key)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = conn.Write(val)

// 	return err
// }
// func (s *Server) handleSetCmd(conn net.Conn, msg *Message) error {

// 	if err := s.cache.Set(msg.Key, msg.Value, msg.TTL); err != nil {
// 		return err
// 	}

// 	go s.sendToFollowers(context.TODO(), msg)
// 	return nil
// }

// func (s *Server) sendToFollowers(ctx context.Context, msg *Message) error {
// 	// using context so that a deadline can be set
// 	for conn := range s.followers {
//     fmt.Println("forwarding request to folllower")
//     rawMsg := msg.ToBytes()
//     fmt.Println("forwarding rawmsg to follower: ", string(rawMsg))
// 		_, err := conn.Write(rawMsg)
// 		if err != nil {
// 			fmt.Println("Write to follower error: ", err)
// 			continue
// 		}
// 	}
// 	return nil
// }
