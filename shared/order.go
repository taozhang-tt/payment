package shared

type OrderStatus int

const (
	OrderStatusInit OrderStatus = iota + 1
	OrderStatusFail
	OrderStatusSuccess
)
