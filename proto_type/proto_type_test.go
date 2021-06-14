package proto_type

import "testing"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Clone() Cloneable {
	res := *p
	return &res
}

func TestProtoType(t *testing.T) {
	manager := NewManager()

	obj := &Person{Name: "tom", Age: 33}

	manager.Set("person1", obj)

	res := manager.Get("person1").(*Person)

	if obj == res {
		t.Errorf("obj: %p,res %p\n", obj, res)
	}

	if obj.Name != res.Name {
		t.Errorf("got: %s,want %s\n", res.Name, obj.Name)
	}
	if obj.Age != res.Age {
		t.Errorf("got: %d,want %d\n", res.Age, obj.Age)
	}
}
