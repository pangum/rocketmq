package rocketmq

import (
	`github.com/apache/rocketmq-client-go/v2/consumer`
)

func parseStrategy(config string) consumer.AllocateStrategy {
	strategy := consumer.AllocateByAveragely
	switch config {
	case "averagely":
		strategy = consumer.AllocateByAveragely
	case "circle":
		strategy = consumer.AllocateByAveragelyCircle
	case "nearby":
		strategy = consumer.AllocateByMachineNearby
	default:
		strategy = consumer.AllocateByAveragely
	}

	return strategy
}
