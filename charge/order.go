package charge

type Order interface {
	Deliver(orderId, transId string) error // 发货
	Refund(orderId string) error           // 退款
}
