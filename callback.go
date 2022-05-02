package payment

type Callback interface {
	Callback() error
}
