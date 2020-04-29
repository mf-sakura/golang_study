package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mf-sakura/golang_study/channel_practical/server/internal/message"
	pb "github.com/mf-sakura/golang_study/channel_practical/server/proto"
)

// MyChatService is implementation of pb.ChatService interface
type MyChatService struct {
	ctx         context.Context
	ctxCancel   context.CancelFunc
	processor   message.Processor
	messageCh   chan<- message.ChatMessage
	streams     map[string]pb.Chat_EnterServer
	mu          *sync.Mutex
	wg          *sync.WaitGroup
	stopSending bool
}

// NewMyChatService reutrns MyChatService
func NewMyChatService() *MyChatService {
	ctx, cancel := context.WithCancel(context.Background())
	processor := message.NewProcessor()
	s := MyChatService{
		ctx:       ctx,
		ctxCancel: cancel,
		processor: processor,
		messageCh: processor.MessageChan(),
		streams:   make(map[string]pb.Chat_EnterServer),
		mu:        &sync.Mutex{},
		wg:        &sync.WaitGroup{},
	}
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.start()
	}()
	return &s
}

// Stop stops MyChatService
func (s *MyChatService) Stop() {
	s.stopSending = true
	// processorの処理を終了し、終了を待つ
	s.processor.StopAndWait()
	go func() {
		// 一定時間の後にCancelする。
		// Sleepは望ましくないので、deferで呼ぶ事が多い。
		// 例) ワーカーでメインの処理関数が終わる際に、deferでcanncelを呼ぶ
		// time.Sleep(1 * time.Millisecond)
		time.Sleep(10 * time.Second)
		s.ctxCancel()
	}()
	// 送信処理の完了を待つ
	// time.Sleep(1 * time.Second)
	s.wg.Wait()
}

// Enter enters a chat room
func (s *MyChatService) Enter(req *pb.EnterRequest, stream pb.Chat_EnterServer) error {
	s.addStreamMap(req.GetUserName(), stream)
	// ユーザーがChatの受信を辞めるにはExitの様な処理が必要
	for {
		// 停止処理されると抜ける
		if s.stopSending {
			break
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *MyChatService) addStreamMap(userName string, stream pb.Chat_EnterServer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.streams[userName] = stream
}

func (s *MyChatService) start() {
	s.processor.Start()
	resultCh := s.processor.ResultChan()
loop:
	for {
		select {
		case r, ok := <-resultCh:
			if !ok {
				break loop
				// continue loop
			}
			// errorをchannelに含める事で、エラーハンドリングが容易になる
			// error用のchannelを分けると処理が複雑になる
			if r.Err != nil {
				fmt.Printf("processed message has error. : %v\n", r.Err)
				continue loop
			}
			for name, stream := range s.streams {
				if err := stream.Send(&pb.EnterReply{
					Message: fmt.Sprintf("[TO: %s]\n%s", name, r.Message),
				}); err != nil {
					fmt.Printf("Send Error. %v\n", r)
					return
				}
			}
		// 強制終了
		case <-s.ctx.Done():
			fmt.Printf("Conetxt Done. %v\n", s.ctx.Err())
			return
		}
	}
}

// SendMessage Send Message
func (s *MyChatService) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageReply, error) {
	// `message :=`と定義すると`package message`を呼べなくなるので、別の名前にする
	m := req.GetMessage()
	senderName := req.GetSenderName()
	if senderName == "" {
		return &pb.SendMessageReply{
			Accepted:     false,
			ErrorMessage: "Sender Name is Empty",
		}, nil
	}
	// Closeされたchannelに送らない様に、boolを用意する
	if !s.stopSending {
		s.messageCh <- message.ChatMessage{
			Message: m,
			Sender:  senderName,
		}
		return &pb.SendMessageReply{
			Accepted: true,
		}, nil
	}
	// gRPCサーバーのgracefull shutdownまでの間にリクエストが来るケース
	return &pb.SendMessageReply{
		ErrorMessage: "Server Shutdown is in progress",
	}, nil
}
