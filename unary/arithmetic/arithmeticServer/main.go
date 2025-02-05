package main

import (
	"context"
	"net"

	"github.com/Sadham-Hussian/go-gRPC/unary/arithmetic/proto"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	// reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) Add(ctx *context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Subtract(ctx *context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a - b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx *context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}

func (s *server) Divide(ctx *context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a / b

	return &proto.Response{Result: result}, nil
}
