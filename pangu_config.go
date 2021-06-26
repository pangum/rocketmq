package rocketmq

type panguConfig struct {
	// 阿里云配置
	Aliyun aliyunConfig `json:"aliyun" yaml:"aliyun" validate:"required"`
	// Rocketmq配置
	Rocketmq config `json:"rocketmq" yaml:"rocketmq" validate:"required"`
}
