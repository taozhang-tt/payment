package order

type Order struct{}

func (o *Order) Deliver(orderId string) error {

	return nil
}

func (o *Order) Refund(orderId string) error {

	return nil
}
