package ali

// 请求参数说明doc: https://opendocs.alipay.com/open/204/105465/
type AppChargeCommonParam struct {
	AppId      string `binding:"required" json:"app_id"` // 支付宝分配给开发者的应用ID
	Method     string `binding:"required" json:"method"` // 接口名称
	Format     string `json:"format" default:"json"`     // 仅支持JSON
	Charset    string `json:"charset" default:"utf-8"`   // 请求使用的编码格式，如utf-8,gbk,gb2312等
	SignType   string `json:"sign_type"`                 // 商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA
	Sign       string `json:"sign"`                      // 商户请求参数的签名串
	Timestamp  string `json:"timestamp"`                 // 发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
	Version    string `json:"version"`                   // 调用的接口版本，固定为：1.0
	NotifyUrl  string `json:"notify_url"`                // 支付宝服务器主动通知商户服务器里指定的页面http/https路径
	BizContent string `json:"biz_content"`               // 请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
	// AppAuthToken string `json:"app_auth_token"`            // 应用授权
	// ExtendParams ExtendParams `json:"extend_params"`
}
type AppChargeParam struct {
	OutTradeNo  string        `binding:"required" json:"out_trade_no"` // 商户网站唯一订单号，由商家自定义，64个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复
	TotalAmount string        `binding:"required" json:"total_amount"` // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]，金额不能为0
	Subject     string        `binding:"required" json:"subject"`      // 订单标题，注意：不可使用特殊字符，如 /，=，& 等。
	ProductCode string        `json:"product_code"`                    // 销售产品码，商家和支付宝签约的产品码
	GoodsDetail []GoodsDetail `json:"goods_detail"`                    // 订单包含的商品列表信息
	TimeExpire  string        `json:"time_expire"`                     // 绝对超时时间，格式为yyyy-MM-dd HH:mm:ss
}

type GoodsDetail struct {
	GoodsId        string `binding:"required" json:"goods_id"`   // 商品的编号
	AlipayGoodsId  string `json:"alipay_goods_id"`               // 支付宝定义的统一商品编号
	GoodsName      string `binding:"required" json:"goods_name"` // 商品名称
	Quantity       int    `binding:"required" json:"quantity"`   // 商品数量
	Price          string `binding:"required" json:"price"`      // 商品单价，单位为元
	GoodsCategory  string `json:"goods_category"`                // 商品类目
	CategoriesTree string `json:"categories_tree"`               // 商品类目树，从商品类目根节点到叶子节点的类目id组成，类目id值使用|分割
	ShowUrl        string `json:"show_url"`                      // 商品的展示地址
}

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
