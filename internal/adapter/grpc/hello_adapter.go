package grpc

import (
	"context"
	"github.com/jmarc101/grpc-protobuf/protogen/go/greeting"
)

func (g *GrpcAdapter) Greet(ctx context.Context, request *greeting.GreetingsRequest) (*greeting.GreetingsResponse, error) {
	greet := g.greetingService.Greet(request.Name)

	return &greeting.GreetingsResponse{
		Message: greet,
	}, nil
}
