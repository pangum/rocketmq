package rocketmq

import (
	mq `github.com/apache/rocketmq-client-go/v2`
	`github.com/apache/rocketmq-client-go/v2/consumer`
	`github.com/apache/rocketmq-client-go/v2/primitive`
	`github.com/storezhang/pangu`
)

func getConsumer(config *pangu.Config) (pc mq.PushConsumer, err error) {
	panguConfig := new(panguConfig)
	if err = config.Load(panguConfig); nil != err {
		return
	}

	aliyun := panguConfig.Aliyun
	rocketmq := panguConfig.Rocketmq

	credentials := primitive.Credentials{
		AccessKey: aliyun.AccessKey,
		SecretKey: aliyun.SecretKey,
	}
	pc, err = mq.NewPushConsumer(
		consumer.WithCredentials(credentials),
		consumer.WithGroupName(rocketmq.Group),
		consumer.WithNameServer(rocketmq.Endpoints),
		consumer.WithNamespace(rocketmq.Namespace),
		consumer.WithInstance(rocketmq.Namespace),
		consumer.WithConsumerModel(parseModel(rocketmq.Consumer.Model)),
		consumer.WithAutoCommit(rocketmq.Consumer.Commit),
		consumer.WithStrategy(parseStrategy(rocketmq.Consumer.Strategy)),
		consumer.WithRetry(rocketmq.Retry),
		// consumer.WithVIPChannel(true),
		// 每次只消费一条数据
		consumer.WithConsumeMessageBatchMaxSize(1),
		// 消息追踪
		// 注意：阿里云线上的Rocketmq一定要加此选项，不然会导致消费不成功
		consumer.WithTrace(&primitive.TraceConfig{
			GroupName:    rocketmq.Group,
			Access:       primitive.Cloud,
			NamesrvAddrs: rocketmq.Endpoints,
			Credentials:  credentials,
		}),
	)

	return
}
