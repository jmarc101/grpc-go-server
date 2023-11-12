package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jmarc101/grpc-go-server/internal/port"
	"github.com/jmarc101/grpc-protobuf/protogen/go/greeting"
	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	greetingService port.GreetingService
	grpcPort        int
	server 			*grpc.Server
	greeting.GreetingServiceServer
}

func NewGrpcAdapter(greetingService port.GreetingService, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		greetingService: greetingService,
		grpcPort:        grpcPort,
	}
}

func (g *GrpcAdapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", g.grpcPort))
	if err != nil {
		log.Fatalln("Failed to listen on port", g.grpcPort, err)
	}

	log.Println("Starting gRPC server on port", g.grpcPort)

	grpcServer := grpc.NewServer()
	g.server = grpcServer

	greeting.RegisterGreetingServiceServer(grpcServer, g)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve grpc server on port %d \n\n%v", g.grpcPort, err)
	}
}

func (g *GrpcAdapter) Stop() {
	g.server.Stop()
}
