package service

import (
	"context"
	"fmt"

	pb "github.com/mf-sakura/golang_study/grpc/server/proto"
)

// MyGreeterService is implementation of pb.GreeterService
// pb.GreeterServiceの実装
type MyGreeterService struct{}

// SayHello returns hello message
func (s *MyGreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {

	return &pb.HelloReply{
		// reqがpointerなのでメソッド経由で取得する
		Message: fmt.Sprintf("Hello %s %s", req.GetFirstName(), req.GetLastName()),
	}, nil
}
