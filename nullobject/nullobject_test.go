package nullobject

import "testing"

func TestNullObject(t *testing.T) {
	n := &nullCustomer{}

	if n.isNil() == false {
		t.Errorf("wrong")
	}
}
