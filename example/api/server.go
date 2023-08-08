package api

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/dillonstreator/fnopt"
)

type Server struct {
	listener net.Listener
	timeout  time.Duration
	maxConns int
	cert     *tls.Certificate
}

func NewServer(addr string, optFns ...serverOptFn) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	srv := &Server{
		listener: l,
		maxConns: 50,
		timeout:  time.Minute,
	}

	fnopt.From(srv, optFns...)

	return srv, nil
}

type serverOptFn = fnopt.OptFn[Server]

func ServerWithTimeout(timeout time.Duration) serverOptFn {
	return func(cfg *Server) {
		cfg.timeout = timeout
	}
}

func ServerWithMaxConns(maxConns int) serverOptFn {
	return func(cfg *Server) {
		cfg.maxConns = maxConns
	}
}

func ServerWithCert(cert *tls.Certificate) serverOptFn {
	return func(cfg *Server) {
		cfg.cert = cert
	}
}
