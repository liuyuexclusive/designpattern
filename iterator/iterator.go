package iterator

type Iterator interface {
	CanNext() bool
	Next() interface{}
}

type Iterable interface {
	Iter() Iterator
}

type Collection []string

type CollectionIter struct {
	Index      int
	Collection Collection
}

func (c *CollectionIter) CanNext() bool {
	return c.Index < len(c.Collection)
}
func (c *CollectionIter) Next() interface{} {
	res := c.Collection[c.Index]
	c.Index++
	return res
}

func (c *Collection) Iter() Iterator {
	return &CollectionIter{Collection: *c}
}
