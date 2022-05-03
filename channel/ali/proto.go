package ali

type NotifyParam struct {
	NotifyTime string `json:"notify_time"`  // 通知的发送时间。格式为 yyyy-MM-dd HH:mm:ss
	NofifyType string `json:"notify_type"`  // 通知类型。同步异步
	NotifyId   string `json:"notify_id"`    // 通知检验 ID
	AppId      string `json:"app_id"`       // 支付宝分配给开发者的应用 APPID
	Charset    string `json:"charset"`      // 编码
	Version    string `json:"version"`      // 调用的接口版本
	SignType   string `json:"sign_type"`    // 商家生成签名字符所使用的签名算法类型: RSA、RSA2
	Sign       string `json:"sign"`         // 签名
	TradeNo    string `json:"trade_no"`     // 支付宝交易号
	OutTradeNo string `json:"out_trade_no"` // 原支付请求的商家订单号
	OutBizNo   string `json:"out_biz_no"`
}

type QueryAlipayTradeResp struct {
	Response struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		SubMsg       string `json:"sub_msg"`      // code=10000 时为成功，否则该字段返回报错信息
		TradeNo      string `json:"trade_no"`     // 支付宝交易号
		OutTradeNo   string `json:"out_trade_no"` // 上家订单号
		BuyerLogonId string `json:"buyer_logon_id"`
		TradeStatus  string `json:"trade_status"`
		PayCurrency  string `json:"pay_currency"`  // 订单支付币种
		PayAmount    string `json:"pay_amount"`    // 支付币种订单金额
		SendPayDate  string `json:"send_pay_date"` // 本次交易打款给卖家的时间 2014-11-27 15:45:57
	} `json:"alipay_trade_query_response"`
}
