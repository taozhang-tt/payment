package ali

type Ali struct {
	AppId            string
	SignType         string // 签名类型：RSA、RSA2
	AliPubKey        string // 支付宝公钥字符串
	RsaPriKey        string // 自己生成的密钥字符串
	NotifyUrl        string
	IsSandbox        bool
	SpecifiedChannel string
}

func init() {
	//channel.Register("ali", func(conf *model.ChannelConfig) (channel.Channel, error) {
	//	extra := new(Extra)
	//	if err := json.Unmarshal([]byte(conf.Extra), extra); err != nil {
	//		return nil, err
	//	}
	//	return &Ali{
	//		AppId:            conf.AppId,
	//		AliPubKey:        conf.PubKey,
	//		RsaPriKey:        conf.PriKey,
	//		NotifyUrl:        conf.NotifyUrl,
	//		SignType:         extra.SignType,
	//		IsSandbox:        extra.IsSandbox,
	//		SpecifiedChannel: extra.SpecifiedChannel,
	//	}, nil
	//})
}

//func (a *Ali) Pay(order *model.Order) (interface{}, error) {
//	bizContentMap := map[string]string{
//		"out_trade_no": order.OrderId,
//		"product_code": "QUICK_MSECURITY_PAY",
//		"total_amount": order.Amount,
//		"subject":      order.ProductName,
//	}
//	if a.SpecifiedChannel != "" {
//		bizContentMap["specified_channel"] = a.SpecifiedChannel
//	}
//
//	bizContent, err := json.Marshal(bizContentMap)
//	if err != nil {
//		return nil, fmt.Errorf("json.Marshal(%v) with error(%v)", bizContentMap, err)
//	}
//	params := map[string]string{
//		"charset":     "utf-8",
//		"version":     "1.0",
//		"method":      "alipay.trade.app.pay",
//		"app_id":      a.AppId,
//		"sign_type":   a.SignType,
//		"notify_url":  a.NotifyUrl,
//		"biz_content": string(bizContent),
//		"timestamp":   time.Now().Format("2006-01-02 15:04:05"),
//	}
//	vals := url.Values{}
//	for k, v := range params {
//		vals.Add(k, v)
//	}
//
//	sign, err := calSign(vals, a.RsaPriKey)
//	if err != nil {
//		return nil, err
//	}
//	params["sign"] = sign
//	return params, nil
//}
//
//// 支付回调
//func (a *Ali) Callback(vals url.Values) error {
//	sign := vals.Get("sign")
//	vals.Del("sign")
//	vals.Del("sign_type")
//
//	ok, err := verifySign(vals, sign, a.AliPubKey)
//	if err != nil {
//		return err
//	}
//	if !ok {
//		return errors.New("verify sign failed")
//	}
//	tradeNo := vals.Get("trade_no")
//	outTradeNo := vals.Get("out_trade_no")
//	totalAmount := vals.Get("total_amount")
//
//	order, err := model.GetOrderById(outTradeNo)
//	if err != nil {
//		return err
//	}
//	order.ChannelOrderId = tradeNo
//	order.Amount = totalAmount
//	order.Status = int(shared.OrderStatusDeliver)
//
//	// TODO: 上锁
//	// 更新订单
//	if err = model.UpdateOrder(order); err != nil {
//		return err
//	}
//	// 发货
//	if err = channel.Deliver(order); err != nil {
//		return err
//	}
//	return nil
//}
//
//// 交易查询
//func (a *Ali) Query(order *model.Order) (shared.OrderStatus, error) {
//	vals := url.Values{
//		"app_id":      {a.AppId},
//		"method":      {"alipay.trade.query"},
//		"charset":     {"utf-8"},
//		"sign_type":   {"RSA2"},
//		"timestamp":   {time.Now().Format("2006-01-02 15:04:05")},
//		"version":     {"1.0"},
//		"biz_content": {fmt.Sprintf(`{"out_trade_no":"%v"}`, order.OrderId)},
//	}
//	sign, err := calSign(vals, a.RsaPriKey)
//	vals.Add("sign", sign)
//
//	url := fmt.Sprintf("https://openapi.alipay.com/gateway.do?%s", vals.Encode())
//	response, err := http.Get(url)
//	if err != nil {
//		return 0, err
//	}
//	defer response.Body.Close()
//	bs, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		return 0, err
//	}
//
//	resp := new(QueryAlipayTradeResp)
//	if err = json.Unmarshal(bs, resp); err != nil {
//		return 0, err
//	}
//
//	if resp.Response.Code != "10000" {
//		err = fmt.Errorf("query alipay trade failed, msg: %s, sub_msg: %s", resp.Response.Msg, resp.Response.SubMsg)
//		return 0, err
//	}
//	var status shared.OrderStatus
//	switch resp.Response.TradeStatus {
//	case "WAIT_BUYER_PAY":
//		status = shared.OrderStatusInit
//	case "TRADE_CLOSED":
//		status = shared.OrderStatusFail
//	case "TRADE_SUCCESS", "TRADE_FINISHED":
//		status = shared.OrderStatusSuccess
//	}
//	return status, nil
//}
//
//func calSign(vals url.Values, key string) (string, error) {
//	content := vals.Encode()
//	sig, err := util.SignPKCS1v15WithPemKey([]byte(content), []byte(key), crypto.SHA256)
//	if err != nil {
//		return "", err
//	}
//	s64 := base64.StdEncoding.EncodeToString(sig)
//	return s64, nil
//}
//
//func verifySign(vals url.Values, sig, key string) (bool, error) {
//	content := vals.Encode()
//	err := util.VerifyPKCS1v15WithDerKey([]byte(content), []byte(sig), []byte(key), crypto.SHA256)
//	if err != nil {
//		return false, err
//	}
//	return true, nil
//}
