package adapter

import (
	"context"
	pb "go-server/internal/apps/adapter/gen_grpc"
)

type SampleAdapter struct {}

func (s *SampleAdapter) GetMySample(ctx context.Context, message *pb.Sample_GetMySampleMessage) (*pb.Sample_MySampleResponse, error) {
	return &pb.Sample_MySampleResponse{
		Name: "hogehoge",
		Kind: "fugafuga",
	}, nil
}
