package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/mf-sakura/golang_study/channel_practical/server/internal/service"
	pb "github.com/mf-sakura/golang_study/channel_practical/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listenPort, err := net.Listen("tcp", ":5502")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	chatService := service.NewMyChatService()

	gracefulSignals := make(chan os.Signal, 1)
	signal.Notify(gracefulSignals, syscall.SIGTERM)
	go func() {
		s := <-gracefulSignals
		fmt.Printf("Got signal: %s\n", s)
		fmt.Println("start graceful shutdown")
		// メッセージ処理の停止
		chatService.Stop()
		// gRPCサーバーの停止
		grpcServer.GracefulStop()
		fmt.Println("Gracefully Shutdown")
	}()

	pb.RegisterChatServer(grpcServer, chatService)
	// Serverの起動
	grpcServer.Serve(listenPort)
}
