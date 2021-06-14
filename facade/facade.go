package facade

import "fmt"

type API interface {
	Test() string
}

type AModuleAPI interface {
	TestA() string
}
type BModuleAPI interface {
	TestB() string
}

type AModuleAPIImpl struct{}
type BModuleAPIImpl struct{}

func (a *AModuleAPIImpl) TestA() string {
	return "a module running"
}
func (b *BModuleAPIImpl) TestB() string {
	return "b module running"
}

type APIImpl struct {
	A AModuleAPI
	B BModuleAPI
}

func (a *APIImpl) Test() string {
	return fmt.Sprintf("%s,%s", a.A.TestA(), a.B.TestB())
}

func NewAPI(a AModuleAPI, b BModuleAPI) API {
	return &APIImpl{A: a, B: b}
}
