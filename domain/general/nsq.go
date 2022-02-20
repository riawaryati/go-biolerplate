package general

import "github.com/nsqio/go-nsq"

type NSQConsumerInput struct {
	Topic       string
	Channel     string
	MaxInFlight int
	MaxAttempts uint16
	Concurrency int
	Handler     nsq.Handler
}
