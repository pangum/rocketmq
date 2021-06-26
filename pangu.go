package rocketmq

import (
	`github.com/storezhang/pangu`
)

func init() {
	if err := pangu.New().Provides(getProducer, getPushConsumer); nil != err {
		panic(err)
	}
}
