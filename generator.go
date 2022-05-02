package payment

type Generator interface {
	BeforeGenerate() error
	Generate() error
	AfterGenerate() error
}
