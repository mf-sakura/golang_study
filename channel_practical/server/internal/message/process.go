package message

import (
	"errors"
	"fmt"
	"sync"
	"unicode/utf8"

	"github.com/mf-sakura/golang_study/channel_practical/server/pkg/counter"
)

// ChatMessage is struct representing chat message
type ChatMessage struct {
	Message string
	Sender  string
}

// ChatResult is struct representing chat processing result
type ChatResult struct {
	Message string
	ID      int64
	Err     error
}

// Processor is message processor interface
type Processor interface {
	Start() error
	StopAndWait()
	ResultChan() <-chan ChatResult
	MessageChan() chan<- ChatMessage
}

// MyProcessor is my implementation of Processor Interface
// interfaceにして、structはMyのprefixをつける
type MyProcessor struct {
	messageCh chan ChatMessage
	resultCh  chan ChatResult
	counter   *counter.Client
	wg        sync.WaitGroup
}

// NewProcessor returns new MyProcessor
func NewProcessor() Processor {
	return &MyProcessor{
		counter:   counter.New(),
		resultCh:  make(chan ChatResult, 100),
		messageCh: make(chan ChatMessage, 100),
	}
}

// Start statrts to process messages
// 以下の役割を持つ
// 1. 入力channelの受け取り
// 2. メッセージ処理のgoroutineの起動
func (p *MyProcessor) Start() error {
	// nilのchannelからの読み込みはブロックされる為
	if p.messageCh == nil {
		return errors.New("input message channel is nil")
	}
	if p.resultCh == nil {
		return errors.New("output result channel is nil")
	}

	// 重い処理を別Goroutineで並行に動作させる
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		// メッセージの処理が完了した時点で、resultChをCloseする
		defer close(p.resultCh)
		// 送信順を担保するためにここでGoroutineは使わない
		for m := range p.messageCh {
			// メッセージの処理完了を待つ為にWaitGroupを使う
			if utf8.RuneCountInString(m.Message) > 20 {
				p.resultCh <- ChatResult{
					Err: errors.New("Word Count must be less than or equal to 20 "),
				}
			}
			// thread safe
			id := p.counter.Increment()
			p.resultCh <- ChatResult{
				Message: fmt.Sprintf("[FROM: %s]\n%s", m.Sender, m.Message),
				ID:      id,
			}
		}
	}()

	return nil
}

// StopAndWait stops to process message and wait finished
func (p *MyProcessor) StopAndWait() {
	// Processorにおけるmessageの受信を止める
	close(p.messageCh)
	// Processorにおけるmessageの処理の完了を待つ
	p.wg.Wait()
}

// ResultChan returns channel to recieve processed messages
func (p *MyProcessor) ResultChan() <-chan ChatResult {
	return p.resultCh
}

// MessageChan returns channel to send messages to process
func (p *MyProcessor) MessageChan() chan<- ChatMessage {
	return p.messageCh
}
