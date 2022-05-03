package charge

import (
	"net/http"
	"payment/model"
	"payment/shared"
)

type Client interface {
	AppCharge(param *AppChargeParam) (interface{}, error) // 创建第三方支付单参数
	Handle(r *http.Request, order Order) error            // 处理第三方回调
	Query(orderId string) (shared.OrderStatus, error)     // 查询
}

var creators = make(map[string]Creator)

type Creator func(*model.ChargeConfig) (Client, error)

func Register(typ string, creator Creator) {
	creators[typ] = creator
}

func GetClient(conf *model.ChargeConfig) (Client, error) {
	creator, ok := creators[conf.ChargeType]
	if !ok {
		panic("charge type does not register")
	}
	return creator(conf)
}

type AppChargeParam struct {
	OrderId          string `binding:"required"`
	Amount           string `binding:"required"`
	ProductName      string
	ProductId        string
	ChannelProductId string
}
