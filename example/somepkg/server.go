package somepkg

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

func ServerWithTimeout(timeout time.Duration) fnopt.OptFn[Server] {
	return func(cfg *Server) {
		cfg.timeout = timeout
	}
}

func ServerWithMaxConns(maxConns int) fnopt.OptFn[Server] {
	return func(cfg *Server) {
		cfg.maxConns = maxConns
	}
}

func ServerWithCert(cert *tls.Certificate) fnopt.OptFn[Server] {
	return func(cfg *Server) {
		cfg.cert = cert
	}
}

func NewServer(addr string, optFns ...fnopt.OptFn[Server]) (*Server, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	srv := &Server{
		listener: l,
		maxConns: 50,
		timeout:  time.Minute,
	}

	fnopt.NewFrom(srv, optFns...)

	return srv, nil
}
