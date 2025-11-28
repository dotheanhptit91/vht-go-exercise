package sharedcomponent

import (
	"flag"
	"net"
	"vht-go/shared"

	sctx "github.com/viettranx/service-context"
	"google.golang.org/grpc"
)

const (
	DefaultGrpcServerPort = "5555"
)

type GrpcServerComp struct {
	id     string
	port   string
	server *grpc.Server
	logger sctx.Logger
}

func NewGrpcServerComp(id string) *GrpcServerComp {
	return &GrpcServerComp{id: id}
}

func (g *GrpcServerComp) ID() string {
	return g.id
}

func (g *GrpcServerComp) InitFlags() {
	flag.StringVar(&g.port, "grpc-server-port", DefaultGrpcServerPort, "gRPC server port")
}

func (g *GrpcServerComp) Activate(sctx sctx.ServiceContext) error {
	g.logger = sctx.Logger("grpc-server")
	g.server = grpc.NewServer()
	return nil
}

func (g *GrpcServerComp) Stop() error {
	return nil
}

func (g *GrpcServerComp) Port() string {
	return g.port
}

func (g *GrpcServerComp) GetServer() *grpc.Server {
	return g.server
}

func (g *GrpcServerComp) Serve() {
	lis, err := net.Listen("tcp", ":"+g.port)
	if err != nil {
		g.logger.Errorln("Failed to listen on port", g.port, err)
		return
	}

	go func() {
		defer shared.RecoverApp()
		if err := g.server.Serve(lis); err != nil {
			g.logger.Errorln("Failed to serve gRPC server", err)
		}
	}()
}

func (g *GrpcServerComp) Register(fnc func(s *grpc.Server)) {
	fnc(g.server)
}

type IGrpcServerComp interface {
	Port() string
	GetServer() *grpc.Server
	Register(fnc func(s *grpc.Server))
	Serve()
}