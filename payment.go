package payment

import (
	"errors"
)

var (
	ErrOrderGeneratorNil    = errors.New("order generator is nil")
	ErrSDKParamGeneratorNil = errors.New("sdk param generator is nil")
	ErrCallbackNil          = errors.New("callback is nil")
)

type Payment struct {
	orderGenerator    OrderGenerator
	sdkParamGenerator SDKParamGenerator
	callback          Callback
}

type Option func(payment *Payment)

func WithOrderGenerator(v OrderGenerator) Option {
	return func(payment *Payment) {
		payment.orderGenerator = v
	}
}

func WithSDKParamGenerator(v SDKParamGenerator) Option {
	return func(payment *Payment) {
		payment.sdkParamGenerator = v
	}
}

func WithCallback(v Callback) Option {
	return func(payment *Payment) {
		payment.callback = v
	}
}

func NewPayment(opts ...Option) *Payment {
	p := &Payment{}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (p *Payment) GenerateOrder() error {
	if p.orderGenerator == nil {
		return ErrOrderGeneratorNil
	}

	if err := p.orderGenerator.BeforeGenerate(); err != nil {
		return err
	}
	if err := p.orderGenerator.Generate(); err != nil {
		return err
	}
	if err := p.orderGenerator.AfterGenerate(); err != nil {
		return err
	}

	return nil
}

func (p *Payment) GenerateSDKParam() error {
	if p.sdkParamGenerator == nil {
		return ErrSDKParamGeneratorNil
	}

	if err := p.sdkParamGenerator.BeforeGenerate(); err != nil {
		return err
	}
	if err := p.sdkParamGenerator.Generate(); err != nil {
		return err
	}
	if err := p.sdkParamGenerator.AfterGenerate(); err != nil {
		return err
	}

	return nil
}

func (p *Payment) Callback() error {
	if p.callback == nil {
		return ErrCallbackNil
	}

	if err := p.callback.Callback(); err != nil {
		return err
	}

	return nil
}

func (p *Payment) Query() error {
	panic("not implement")
}
