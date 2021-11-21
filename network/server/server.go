package server

import (
	"fmt"
	"ip_reporter/common"
	"log"
	"net"
)

type Server struct {
	ip   string
	port int
	conn *net.UDPConn
}

func init() {
	log.SetPrefix("[IP_REPORTER]")
}

func NewServer(ip string, port int) *Server {
	return &Server{
		ip:   ip,
		port: port,
	}
}

func (s *Server) Listen() int {
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		log.Println(err)
		return common.RESULT_NOK
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Println(err)
		return common.RESULT_NOK
	}
	s.conn = conn
	return common.RESULT_OK
}

func (s *Server) RecvFrom(buffer []byte) (int, *net.UDPAddr, error) {
	return s.conn.ReadFromUDP(buffer)
}

func (s *Server) Sendto(buffer []byte, addr *net.UDPAddr) (int, error) {
	return s.conn.WriteToUDP(buffer, addr)
}

func (s *Server) Close() {
	s.conn.Close()
}
