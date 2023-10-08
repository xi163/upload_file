package kafka

import (
	"errors"

	"github.com/cwloo/gonet/logs"
	config "github.com/cwloo/uploader/src/config"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"

	promePkg "github.com/cwloo/uploader/src/global/pkg/common/prometheus"
)

type Producer struct {
	topic    string
	addr     []string
	config   *sarama.Config
	producer sarama.SyncProducer
}

func NewKafkaProducer(addr []string, topic string) *Producer {
	p := Producer{}
	p.config = sarama.NewConfig()
	p.config.Producer.Return.Successes = true
	p.config.Producer.Return.Errors = true
	p.config.Producer.RequiredAcks = sarama.WaitForAll
	p.config.Producer.Partitioner = sarama.NewHashPartitioner
	if config.Config.Kafka.SASLUserName != "" && config.Config.Kafka.SASLPassword != "" {
		p.config.Net.SASL.Enable = true
		p.config.Net.SASL.User = config.Config.Kafka.SASLUserName
		p.config.Net.SASL.Password = config.Config.Kafka.SASLPassword
	}
	p.addr = addr
	p.topic = topic
	producer, err := sarama.NewSyncProducer(p.addr, p.config)
	if err != nil {
		logs.Fatalf(err.Error())
		return nil
	}
	p.producer = producer
	return &p
}

func (p *Producer) SendMessage(m proto.Message, key string, operationID string) (int32, int64, error) {
	logs.Infof("%v %v %v", key, m.String(), p.producer)
	kMsg := &sarama.ProducerMessage{}
	kMsg.Topic = p.topic
	kMsg.Key = sarama.StringEncoder(key)
	bMsg, err := proto.Marshal(m)
	if err != nil {
		logs.Errorf(err.Error())
		return -1, -1, err
	}
	if len(bMsg) == 0 {
		return 0, 0, errors.New(logs.SprintErrorf("error"))
	}
	kMsg.Value = sarama.ByteEncoder(bMsg)
	if kMsg.Key.Length() == 0 || kMsg.Value.Length() == 0 {
		logs.Errorf("error")
		return -1, -1, errors.New(logs.SprintErrorf("error"))
	}
	a, b, err := p.producer.SendMessage(kMsg)
	if err == nil {
		promePkg.PromeInc(promePkg.SendMsgCounter)
	}
	return a, b, err
}
