package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/mf-sakura/golang_study/grpc/go_client/proto"
	"google.golang.org/grpc"
)

func main() {
	// ローカルなのでInsecureで接続する
	conn, err := grpc.Dial("127.0.0.1:5502", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection establish failed. %v", err)
	}
	// connectionのcloseを忘れない
	defer conn.Close()
	// clientの生成
	helloClient := pb.NewGreeterClient(conn)
	helloReq := &pb.HelloRequest{
		FirstName: "yuto",
		LastName:  "sakura",
	}
	// ここでgRPCサーバーへリクエストがされる
	res, err := helloClient.SayHello(context.Background(), helloReq)

	if err != nil {
		log.Fatalf("SayHello failed. %v", err)
	}
	fmt.Println(res.GetMessage())

	// clientの生成
	// 前と同じコネクションを使い回す(TCPのハンドシェイクが不要)
	noticicationClient := pb.NewNotifierClient(conn)
	preodicReq := &pb.PereodicHelloRequest{
		FirstName: "yuto",
		LastName:  "sakura",
	}
	// Server Side Streamingなのでstreamを受け取る。
	stream, err := noticicationClient.PereodicHello(context.Background(), preodicReq)
	if err != nil {
		log.Fatalf("PereodicHello failed. %v", err)
	}
	for {
		// サーバーからストリームを受け取ると、resに値が入り処理が進む
		// 内部実装としては、サーバーからストリームを受け取るまで、処理がブロックされる様になっている。
		// 具体的には、チャンネルの受信待ちになっており、別Goroutine(スレッド)でサーバーからのストリームを受信できればチャンネルに送信される。
		res, err := stream.Recv()
		if err != nil {
			log.Fatalf("Recieve Stream failed. %v", err)
		}
		fmt.Printf("Stream Recieved at %v. Message: %s\n", time.Now(), res.GetMessage())
	}
}
