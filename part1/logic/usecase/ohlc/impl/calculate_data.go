package impl

import (
	"context"

	"github.com/Shopify/sarama"
)

func (u *usecase) CalculateData(ctx context.Context) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true

	// On the broker side, you may want to change the following settings to get
	// stronger consistency guarantees:
	// - For your broker, set `unclean.leader.election.enable` to false
	// - For the topic, you could increase `min.insync.replicas`.

	producer, _ := sarama.NewSyncProducer([]string{"redpanda-0:9092"}, config)
	producer.SendMessage(&sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.StringEncoder("testset"),
	})

	return nil
}
