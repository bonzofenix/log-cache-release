package integrationfakes

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"

	logcache "code.cloudfoundry.org/go-log-cache/v2/rpc/logcache_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type FakeLogCache struct {
	port int
	addr string
	c    *tls.Config
	s    *grpc.Server

	serveCh chan error

	logcache.UnimplementedEgressServer
	logcache.UnimplementedIngressServer
	logcache.UnimplementedPromQLQuerierServer
}

func NewFakeLogCache(port int, c *tls.Config) *FakeLogCache {
	return &FakeLogCache{
		port:    port,
		c:       c,
		serveCh: make(chan error),
	}
}

func (f *FakeLogCache) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", f.port))
	if err != nil {
		return err
	}
	f.addr = lis.Addr().String()

	f.s = grpc.NewServer()
	if f.c != nil {
		f.s = grpc.NewServer(grpc.Creds(credentials.NewTLS(f.c)))
	}

	logcache.RegisterEgressServer(f.s, f)
	logcache.RegisterIngressServer(f.s, f)
	logcache.RegisterPromQLQuerierServer(f.s, f)

	go func() {
		f.serveCh <- f.s.Serve(lis)
	}()

	return nil
}

func (f *FakeLogCache) Address() string {
	return f.addr
}

func (f *FakeLogCache) Read(ctx context.Context, req *logcache.ReadRequest) (*logcache.ReadResponse, error) {
	return nil, nil
}

func (f *FakeLogCache) Stop() error {
	f.s.Stop()
	return <-f.serveCh
}
