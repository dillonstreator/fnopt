package api

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/dillonstreator/fnopt"
)

type ServerE struct {
	listener net.Listener
	timeout  time.Duration
	maxConns int
	cert     *tls.Certificate
}

func NewServerE(addr string, optFns ...serverEOptFn) (*ServerE, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	srv := &ServerE{
		listener: l,
		maxConns: 50,
		timeout:  time.Minute,
	}

	err = fnopt.FromE(srv, optFns...)
	if err != nil {
		return nil, err
	}

	return srv, nil
}

type serverEOptFn = fnopt.OptFnE[ServerE]

func ServerEWithTimeout(timeout time.Duration) serverEOptFn {
	return func(cfg *ServerE) error {
		if timeout < 0 {
			return fmt.Errorf("invalid timeout less than 0: %s", timeout)
		}

		cfg.timeout = timeout
		return nil
	}
}

func ServerEWithMaxConns(maxConns int) serverEOptFn {
	return func(cfg *ServerE) error {
		if maxConns < 0 {
			return fmt.Errorf("invalid max conns less than 0: %d", maxConns)
		}

		cfg.maxConns = maxConns
		return nil
	}
}

func ServerEWithCert(cert *tls.Certificate) serverEOptFn {
	return func(cfg *ServerE) error {
		if cert == nil {
			return errors.New("invalid nil cert")
		}

		cfg.cert = cert
		return nil
	}
}
