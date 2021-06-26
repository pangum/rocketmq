package rocketmq

import (
	`github.com/apache/rocketmq-client-go/v2/consumer`
)

func parseModel(config string) consumer.MessageModel {
	model := consumer.Clustering
	switch config {
	case "cluster":
		model = consumer.Clustering
	case "broadcast":
		model = consumer.BroadCasting
	default:
		model = consumer.Clustering
	}

	return model
}
