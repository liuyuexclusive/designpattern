package flyweight

import "fmt"

type FlyweightImage struct {
	Data string
}

type FlyweightImageFactory struct {
	M map[string]*FlyweightImage
}

func NewFactory() *FlyweightImageFactory {
	return &FlyweightImageFactory{M: make(map[string]*FlyweightImage)}
}

func NewFlyweightImage(name string) *FlyweightImage {
	data := fmt.Sprintf("data from %s", name)
	return &FlyweightImage{Data: data}
}

func (f *FlyweightImageFactory) Get(name string) *FlyweightImage {
	if v, ok := f.M[name]; ok {
		return v
	}
	f.M[name] = NewFlyweightImage(name)
	return f.M[name]
}
