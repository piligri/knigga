package server

import (
	"log"
	"net"
)

type Server struct {
	Address  string
	Listener net.Listener
}

func NewServer(address string) *Server {
	return &Server{
		Address: address,
	}
}

func (s *Server) Start() error {
	var err error
	s.Listener, err = net.Listen("tcp", s.Address)
	if err != nil {
		log.Printf("Error accept %v", err)
		return err
	}
	log.Printf("Server started")
	go s.acceptLoop()
	select {}
}

func (s *Server) acceptLoop() {
	conn, err := s.Listener.Accept()
	if err != nil {
		log.Printf("Failed accept client: %v", err)
	}
	log.Printf("Welcome %v", conn.RemoteAddr().String())
}
