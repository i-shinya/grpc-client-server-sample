package main

import (
	"go-server/internal/apps/adapter"
	pb "go-server/internal/apps/adapter/gen_grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	log.Println("Starting main.")
	s := grpc.NewServer()

	// gRpcのメソッドを登録する
	registerGrpcAdapter(s)

	listener, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Server starting.")
	err = s.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}

// gRpcで生成したadapterを登録する
func registerGrpcAdapter(s *grpc.Server) {
	// SampleAdapterを登録する
	pb.RegisterSampleAdapterServer(s, &adapter.SampleAdapter{})
}
