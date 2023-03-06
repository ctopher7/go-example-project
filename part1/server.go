package part1

import (
	"log"
	"net"

	"github.com/ctopher7/gltc/v2/part1/logic/handler"
	pb "github.com/ctopher7/gltc/v2/part1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func ServeGrpc(env string) {
	cfg, err := ReadConfig(env)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", cfg.GrpcPort)
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	res := InitResource(cfg)
	initLogic(res)
	pb.RegisterOhlcServer(srv, &handler.GrpcHandler{
		OhlcUsecase: ohlcUsecase,
	})
	reflection.Register(srv)

	if err = srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve, error: %v", err)
	}

}

func ServeConsumer(env string) {
	cfg, err := ReadConfig(env)
	if err != nil {
		panic(err)
	}
	res := InitResource(cfg)
	initLogic(res)
	initSarama(cfg)
	mainSarama()
}
