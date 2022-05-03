package ali

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"payment/charge"
	"payment/model"
	"payment/shared"
	"sort"
	"strings"
	"time"
)

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
	charge.Register("ali", func(conf *model.ChargeConfig) (charge.Client, error) {
		extra := new(Extra)
		if err := json.Unmarshal([]byte(conf.Extra), extra); err != nil {
			return nil, err
		}
		return &Ali{
			AppId:            conf.AppId,
			AliPubKey:        conf.PubKey,
			RsaPriKey:        conf.PriKey,
			NotifyUrl:        conf.NotifyUrl,
			SignType:         extra.SignType,
			IsSandbox:        extra.IsSandbox,
			SpecifiedChannel: extra.SpecifiedChannel,
		}, nil
	})
}

func (a *Ali) AppCharge(param *charge.AppChargeParam) (interface{}, error) {
	vals, err := a.generateAppChargeCommonParam(param)
	if err != nil {
		return nil, err
	}

	sign, err := calSign(vals, a.RsaPriKey)
	if err != nil {
		return nil, err
	}

	ecodedParam := vals.Encode()

	data := appendSign(ecodedParam, sign)

	return data, nil
}

func appendSign(param, sign string) string {
	return param + "&sign=" + sign
}

func (a *Ali) generateAppChargeCommonParam(param *charge.AppChargeParam) (url.Values, error) {
	bizContent, err := a.generateBizContent(param)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("app_id", a.AppId)
	params.Add("method", "alipay.trade.app.pay")
	params.Add("format", "json")
	params.Add("charset", "utf-8")
	params.Add("sign_type", "RSA2")
	params.Add("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	params.Add("version", "1.0")
	params.Add("notify_url", a.NotifyUrl)
	params.Add("biz_content", bizContent)

	return params, nil
}

func (a *Ali) generateBizContent(param *charge.AppChargeParam) (string, error) {
	bizContentMap := map[string]string{
		"out_trade_no": param.OrderId,
		"total_amount": param.Amount,
		"subject":      "alipay 充值",
		"product_code": "QUICK_MSECURITY_PAY",
	}
	if a.SpecifiedChannel != "" {
		bizContentMap["specified_channel"] = a.SpecifiedChannel
	}
	bs, err := json.Marshal(bizContentMap)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

// 支付回调
func (a *Ali) Handle(r *http.Request, order charge.Order) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	vals := r.Form

	sign := vals.Get("sign")
	vals.Del("sign")
	vals.Del("sign_type")

	err := verifySign(vals, sign, a.AliPubKey)
	if err != nil {
		return err
	}

	orderId := vals.Get("trade_no")
	transId := vals.Get("out_trade_no")
	tradStatus := vals.Get("trade_status")

	switch tradStatus {
	case "TRADE_SUCCESS":
		return order.Deliver(orderId, transId)
	default:
	}
	return nil
}

// 交易查询
func (a *Ali) Query(orderId string) (shared.OrderStatus, error) {
	vals := url.Values{
		"app_id":      {a.AppId},
		"method":      {"alipay.trade.query"},
		"charset":     {"utf-8"},
		"sign_type":   {"RSA2"},
		"timestamp":   {time.Now().Format("2006-01-02 15:04:05")},
		"version":     {"1.0"},
		"biz_content": {fmt.Sprintf(`{"out_trade_no":"%v"}`, orderId)},
	}
	sign, err := calSign(vals, a.RsaPriKey)
	vals.Add("sign", sign)

	url := fmt.Sprintf("https://openapi.alipay.com/gateway.do?%s", vals.Encode())
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	bs, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	resp := new(QueryAlipayTradeResp)
	if err = json.Unmarshal(bs, resp); err != nil {
		return 0, err
	}

	if resp.Response.Code != "10000" {
		err = fmt.Errorf("query alipay trade failed, msg: %s, sub_msg: %s", resp.Response.Msg, resp.Response.SubMsg)
		return 0, err
	}
	var status shared.OrderStatus
	switch resp.Response.TradeStatus {
	case "WAIT_BUYER_PAY":
		status = shared.OrderStatusInit
	case "TRADE_CLOSED":
		status = shared.OrderStatusFail
	case "TRADE_SUCCESS", "TRADE_FINISHED":
		status = shared.OrderStatusSuccess
	}
	return status, nil
}

func calSign(vals url.Values, pri string) (string, error) {
	return "", nil
}

// convertVals2String convert the values into string (without url encode)
// ("bar=baz&foo=quux") sorted by key.
func convertVals2String(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	return buf.String()
}

func verifySign(vals url.Values, sig, pub string) error {
	return nil
}
