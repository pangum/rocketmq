package rocketmq

type aliyunConfig struct {
	// 授权，相当于用户名
	AccessKey string `json:"accessKey" yaml:"accessKey" validate:"required"`
	// 授权，相当于密码
	SecretKey string `json:"secretKey" yaml:"secretKey" validate:"required"`
}
