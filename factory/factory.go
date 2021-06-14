package factory

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(a int) {
	o.b = a
}

type OperatorPlus struct {
	*OperatorBase
}

func (o *OperatorPlus) Result() int {
	return o.a + o.b
}

type OperatorMinus struct {
	*OperatorBase
}

func (o *OperatorMinus) Result() int {
	return o.a - o.b
}

type OperatorPlusFactory struct {
}

func (o *OperatorPlusFactory) Create() Operator {
	return &OperatorPlus{&OperatorBase{}}
}

type OperatorMinusFactory struct {
}

func (o *OperatorMinusFactory) Create() Operator {
	return &OperatorMinus{&OperatorBase{}}
}
