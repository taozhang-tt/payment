package shared

type OrderStatus int

const (
	OrderStatusInit OrderStatus = iota
	OrderStatusPay
	OrderStatusDeliver
	OrderStatusSuccess
	OrderStatusRefund
	OrderStatusFail
)
