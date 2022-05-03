package model

type ChargeConfig struct {
	ClinetId   int64 // 租户ID
	AppId      string
	AppSecret  string
	PriKey     string
	PubKey     string
	ChargeType string
	NotifyUrl  string // 回调地址
	Extra      string // json
}
