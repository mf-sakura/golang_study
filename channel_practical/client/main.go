package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	pb "github.com/mf-sakura/golang_study/channel_practical/client/proto"
	"google.golang.org/grpc"
)

func main() {
	var userName string
	flag.StringVar(&userName, "u", "sakura", "")
	flag.Parse()

	// ローカルなのでInsecureで接続する
	conn, err := grpc.Dial("127.0.0.1:5502", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection establish failed. %v", err)
	}
	// connectionのcloseを忘れない
	defer conn.Close()
	// clientの生成
	chatClient := pb.NewChatClient(conn)
	go func() error {
		enterReq := &pb.EnterRequest{
			UserName: userName,
		}
		stream, err := chatClient.Enter(context.Background(), enterReq)
		if err != nil {
			return err
		}
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Fatalf("Recieve Stream failed. %v", err)
			}
			fmt.Printf("---------------\n%s\n", res.GetMessage())
		}
	}()

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()
		fmt.Printf("Scaned messae is %s\n", text)
		_, err := chatClient.SendMessage(context.Background(), &pb.SendMessageRequest{
			Message:    text,
			SenderName: userName,
		})
		if err != nil {
			fmt.Printf("SendMessage error. %v\n", err)
		}
	}
}
