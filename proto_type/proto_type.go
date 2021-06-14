package proto_type

type Cloneable interface {
	Clone() Cloneable
}

type ProtoManager struct {
	Data map[string]Cloneable
}

func (p *ProtoManager) Set(name string, obj Cloneable) {
	p.Data[name] = obj
}
func (p *ProtoManager) Get(name string) Cloneable {
	return p.Data[name].Clone()
}

func NewManager() *ProtoManager {
	return &ProtoManager{Data: make(map[string]Cloneable)}
}
