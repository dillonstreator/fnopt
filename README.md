# `fnopt` - Generic Functional Options for Friendly API Creation

`fnopt` is a Go package that provides a generic-enabled adaptation of Functional Options, following the principles described by Dave Cheney in his article [Functional Options for Friendly APIs](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis). This package makes it easier to create flexible and easy-to-use APIs for configuring Go structs. As seen in [Generic Functional Options for Friendly API Creation](https://blog.streator.me/posts/generic-functional-options/).

## Installation

```sh
go get github.com/dillonstreator/fnopt
```

## Usage

The `fnopt` package consists of three primary components:

### `fnopt.OptFn`

The `OptFn` type represents a functional option that modifies a configuration struct. A functional option is simply a function that takes a pointer to the configuration struct and applies the desired changes to it.

### `fnopt.New`

The `New` function is useful for creating instances of a struct with functional options. It takes the type `T` of the target struct and applies the provided option functions to it, allowing you to initialize and configure the struct in a single call.

### `fnopt.NewFrom`

The `NewFrom` function allows you to modify an existing struct of type `T` by applying the provided option functions to it. This is helpful when you have an existing instance that requires further configuration.

### Error-enabled Versions

In addition to the basic components, `fnopt` also provides error-enabled counterparts, which can be useful when you need to handle errors during the application of functional options.

- `fnopt.OptFnE`: The `OptFnE` type is the error-enabled version of `OptFn`. It allows for functional options that may return an error when modifying the configuration struct.

- `fnopt.NewE`: Similar to `fnopt.New`, `NewE` creates instances of a struct with error-enabled functional options, enabling the initialization and configuration of the struct while handling potential errors.

- `fnopt.NewFromE`: The `NewFromE` function is the error-enabled version of `NewFrom`. It allows you to modify an existing struct with functional options that may return errors during the configuration process.

## Example

To illustrate the usage of `fnopt`, consider the following example:

```go
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

func NewServer(addr string, optFns ...serverFnOpt) (*Server, error) {
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

type serverFnOpt = fnopt.OptFn[Server]

func ServerWithTimeout(timeout time.Duration) serverFnOpt {
	return func(cfg *Server) {
		cfg.timeout = timeout
	}
}

func ServerWithMaxConns(maxConns int) serverFnOpt {
	return func(cfg *Server) {
		cfg.maxConns = maxConns
	}
}

func ServerWithCert(cert *tls.Certificate) serverFnOpt {
	return func(cfg *Server) {
		cfg.cert = cert
	}
}
```

In this example, we define a `Server` struct that we want to configure with functional options. The package provides three functional options (`ServerWithTimeout`, `ServerWithMaxConns`, and `ServerWithCert`) to customize the `Server` instance. We can then use these functional options to create a new `Server` instance with specific configurations after first defining some defaults.

### Error-enabled

```go
package somepkg

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

func NewServerE(addr string, optFns ...serverEFnOpt) (*ServerE, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	srv := &ServerE{
		listener: l,
		maxConns: 50,
		timeout:  time.Minute,
	}

	err = fnopt.NewFromE(srv, optFns...)
	if err != nil {
		return nil, err
	}

	return srv, nil
}

type serverEFnOpt = fnopt.OptFnE[ServerE]

func ServerEWithTimeout(timeout time.Duration) serverEFnOpt {
	return func(cfg *ServerE) error {
		if timeout < 0 {
			return fmt.Errorf("invalid timeout less than 0: %s", timeout)
		}

		cfg.timeout = timeout
		return nil
	}
}

func ServerEWithMaxConns(maxConns int) serverEFnOpt {
	return func(cfg *ServerE) error {
		if maxConns < 0 {
			return fmt.Errorf("invalid max conns less than 0: %d", maxConns)
		}

		cfg.maxConns = maxConns
		return nil
	}
}

func ServerEWithCert(cert *tls.Certificate) serverEFnOpt {
	return func(cfg *ServerE) error {
		if cert == nil {
			return errors.New("invalid nil cert")
		}

		cfg.cert = cert
		return nil
	}
}
```
