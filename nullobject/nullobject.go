package nullobject

type customer interface {
	getName() string
	isNil() bool
}

type customerImpl struct {
	name string
}

func (c *customerImpl) getName() string {
	return c.name
}

type realCustomer struct {
	*customerImpl
}

func (r *realCustomer) isNil() bool {
	return false
}

type nullCustomer struct {
	*customerImpl
}

func (n *nullCustomer) isNil() bool {
	return true
}
