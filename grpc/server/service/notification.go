package service

import (
	"fmt"
	"time"

	pb "github.com/mf-sakura/golang_study/grpc/server/proto"
)

// MyNotifierService is implementation of pb.NotifierService
// pb.NotifierServiceの実装
type MyNotifierService struct{}

// PereodicHello returns hello message pereodically
func (s *MyNotifierService) PereodicHello(req *pb.PereodicHelloRequest, stream pb.Notifier_PereodicHelloServer) error {
	firstName := req.GetFirstName()
	lastName := req.GetLastName()
	for {
		if err := stream.Send(&pb.PeriodicHelloReply{
			Message: fmt.Sprintf("Hello %s %s", firstName, lastName),
		}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
