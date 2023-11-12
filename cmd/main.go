package main

import (
	"github.com/jmarc101/grpc-go-server/internal/adapter/grpc"
	app "github.com/jmarc101/grpc-go-server/internal/application"
)

func main() {
	hs := &app.GreetingService{}
	grpcAdapter := grpc.NewGrpcAdapter(hs, 8080)

	grpcAdapter.Run()


}
