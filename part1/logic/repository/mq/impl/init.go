package impl

import (
	"github.com/Shopify/sarama"
	mqRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/mq"
)

type repository struct {
	mqProducer sarama.SyncProducer
}

func New(
	mqProducer sarama.SyncProducer,
) mqRepo.Repository {
	return &repository{
		mqProducer: mqProducer,
	}
}
