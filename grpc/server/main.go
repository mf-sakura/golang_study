package main

import (
	"log"
	"net"

	pb "github.com/mf-sakura/golang_study/grpc/server/proto"
	"github.com/mf-sakura/golang_study/grpc/server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// ListenするPortの設定
	listenPort, err := net.Listen("tcp", ":5502")
	if err != nil {
		log.Fatalln(err)
	}
	// gRPCのサーバーの初期化
	server := grpc.NewServer()
	// grpcurlを使うためにreflectionを有効化する
	// これによりネットワーク経由でprotobufの情報が分かるので、本番環境では設定しない方が良い。
	reflection.Register(server)
	greeterService := &service.MyGreeterService{}
	notifierService := &service.MyNotifierService{}
	// protobufに基づいて自動生成されたInterfaceを満たすServiceの登録
	pb.RegisterGreeterServer(server, greeterService)
	pb.RegisterNotifierServer(server, notifierService)
	// Serverの起動
	server.Serve(listenPort)
}
