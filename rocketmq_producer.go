package rocketmq

import (
	mq `github.com/apache/rocketmq-client-go/v2`
	`github.com/apache/rocketmq-client-go/v2/primitive`
	`github.com/apache/rocketmq-client-go/v2/producer`
	`github.com/storezhang/pangu`
)

func getProducer(config *pangu.Config) (pd mq.Producer, err error) {
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
	pd, err = mq.NewProducer(
		producer.WithCredentials(credentials),
		producer.WithGroupName(rocketmq.Group),
		producer.WithNameServer(rocketmq.Endpoints),
		producer.WithNamespace(rocketmq.Namespace),
		producer.WithRetry(rocketmq.Retry),
		// 消息追踪
		// 注意：阿里云线上的Rocketmq一定要加此选项，不然会导致消费不成功
		producer.WithTrace(&primitive.TraceConfig{
			GroupName:    rocketmq.Group,
			Access:       primitive.Cloud,
			NamesrvAddrs: rocketmq.Endpoints,
			Credentials:  credentials,
		}),
	)

	return
}
