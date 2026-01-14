package kafka

import (
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaAdmin struct {
	Broker string
}

func NewKafkaAdmin(broker string) *KafkaAdmin {
	return &KafkaAdmin{Broker: broker}
}

// creates topic if it does not exist
func (ka *KafkaAdmin) EnsureTopic(topic string, partitions int) error {
	conn, err := kafka.Dial("tcp", ka.Broker)
	if err != nil {
		return err
	}
	defer conn.Close()

	// check if topic exists
	partitionsInfo, err := conn.ReadPartitions()
	if err != nil {
		return err
	}

	for _, partition := range partitionsInfo {
		if partition.Topic == topic {
			log.Printf("Topic %s already exists", topic)
			return nil
		}
	}

	// create topic
	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	ctrlAddr := fmt.Sprintf("%s:%d", controller.Host, controller.Port)

	ctrlConn, err := kafka.Dial("tcp", ctrlAddr)
	if err != nil {
		return err
	}
	defer ctrlConn.Close()

	return ctrlConn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     partitions,
		ReplicationFactor: 1,
	})
}

// NewWriter returns a kafka.Writer for publishing
func NewWriter(broker, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}
