package rocketmq

type consumerConfig struct {
	// 消费模式
	Model string `default:"cluster" json:"model" yaml:"model" validate:"requried,oneof=cluster broadcast"`
	// 使用自动提交
	Commit bool `default:"true" json:"commit" yaml:"commit"`
	// 消费策略
	Strategy string `default:"averagely" json:"strategy" yaml:"strategy" validate:"required,oneof=averagely circle nearby"`
}
