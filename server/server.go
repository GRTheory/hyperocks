package server

import "crypto/tls"

type Server struct {
	Config    *Config
	TLSConfig *tls.Config
}

func NewServer(protocol, addr string) *Server {
	return &Server{
		Config: &Config{
			Protocol: protocol,
			Addr:     addr,
		},
	}
}

type Config struct {
	Protocol   string
	Addr       string
	Verify     func(string, string) bool
	HttpPath   string
	WsPath     string
	WsCompress bool
}
