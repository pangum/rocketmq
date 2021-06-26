package rocketmq

type config struct {
	// 地址
	Endpoints []string `json:"endpoints" yaml:"endpoints" validate:"required"`
	// 命名空间
	// 阿里云等同于实例Id
	Namespace string `json:"namespace" yaml:"namespace" validate:"required"`
	// 组
	Group string `json:"group" yaml:"group" validate:"required"`
	// 主题
	Topic string `json:"topic" yaml:"topic" validate:"required"`
	// 重试次数
	Retry int `default:"3" json:"retry" yaml:"retry"`
	// 消费配置
	Consumer consumerConfig `json:"consumer" yaml:"consumer" validate:"structonly"`
}
