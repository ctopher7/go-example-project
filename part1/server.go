package part1

import (
	"log"
	"net"

	"github.com/ctopher7/gltc/v2/part1/logic/handler"
	pb "github.com/ctopher7/gltc/v2/part1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Serve() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	initLogic()
	pb.RegisterOhlcServer(srv, &handler.GrpcHandler{
		OhlcUsecase: ohlcUsecase,
	})
	reflection.Register(srv)

	if err = srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve, error: %v", err)
	}

}
