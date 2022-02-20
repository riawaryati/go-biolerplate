package infra

import (
	"encoding/json"
	"os"
	"os/signal"
	"syscall"

	"github.com/furee/backend/domain/general"
	"github.com/nsqio/go-nsq"
)

// =================== PRODUCER SECTION
type NSQProducer struct {
	producer *nsq.Producer
}

// NewNSQProducer init connection to NSQ Producer
func NewNSQProducer(conf general.NSQProducer) (*NSQProducer, error) {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(conf.NSQD, config)
	if err != nil {
		return nil, err
	}

	return &NSQProducer{
		producer: producer,
	}, nil
}

//Publish needs topic and payload as input
//The payload that needed is struct with json tag
func (np NSQProducer) Publish(topic string, msg interface{}) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return np.producer.Publish(topic, payload)
}

// =================== CONSUMER SECTION
type NSQConsumer struct {
	consumerList []*nsq.Consumer
}

// NewNSQConsumer create & register consumer
func NewNSQConsumer(conf general.NSQConsumer, input []general.NSQConsumerInput) (*NSQConsumer, error) {
	consumerList := make([]*nsq.Consumer, 0)

	for _, in := range input {
		config := nsq.NewConfig()

		//default MaxAttempts = 5
		if in.MaxAttempts != 0 {
			config.MaxAttempts = in.MaxAttempts
		}

		//default MaxInFlight = 200
		if in.MaxInFlight != 0 {
			config.MaxInFlight = in.MaxInFlight
		}

		consumer, err := nsq.NewConsumer(in.Topic, in.Channel, config)
		if err != nil {
			return nil, err
		}

		//default concurrency = 2
		conc := 2
		if in.Concurrency != 0 {
			conc = in.Concurrency
		}

		consumer.SetLoggerLevel(nsq.LogLevelError)
		consumer.AddConcurrentHandlers(in.Handler, conc)
		err = consumer.ConnectToNSQLookupd(conf.NSQLookupD)
		if err != nil {
			return nil, err
		}

		// Append to consumerList
		consumerList = append(consumerList, consumer)
	}

	return &NSQConsumer{
		consumerList: consumerList,
	}, nil
}

// NSQStop will stop the consumer gracefully
func (n *NSQConsumer) NSQStop() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	for _, c := range n.consumerList {
		c.Stop()
	}
}
