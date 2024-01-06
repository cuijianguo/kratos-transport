package nsq

import (
	"sync"

	"github.com/cuijianguo/kratos-transport/broker"
	"github.com/nsqio/go-nsq"
)

type subscriber struct {
	sync.RWMutex

	topic string

	n       *nsqBroker
	options broker.SubscribeOptions

	consumer    *nsq.Consumer
	handlerFunc nsq.HandlerFunc

	concurrency int
	closed      bool
}

func (s *subscriber) Options() broker.SubscribeOptions {
	s.RLock()
	defer s.RUnlock()

	return s.options
}

func (s *subscriber) Topic() string {
	s.RLock()
	defer s.RUnlock()

	return s.topic
}

func (s *subscriber) Unsubscribe() error {
	//s.Lock()
	//defer s.Unlock()

	if s.consumer != nil {
		s.consumer.Stop()
	}

	s.closed = true

	if s.n != nil && s.n.subscribers != nil {
		_ = s.n.subscribers.Remove(s.topic)
	}

	return nil
}

func (s *subscriber) IsClosed() bool {
	s.RLock()
	defer s.RUnlock()

	return s.closed
}
