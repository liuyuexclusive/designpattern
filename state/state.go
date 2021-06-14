package state

type WeekDay interface {
	Val() int
	Action(*Context)
}

type Context struct {
	Today WeekDay
}

type Sunday struct{}

func (m *Sunday) Action(c *Context) {
	c.Today = &Sunday{}
}

func (m *Sunday) Val() int {
	return 0
}

type Monday struct{}

func (m *Monday) Action(c *Context) {
	c.Today = &Monday{}
}

func (m *Monday) Val() int {
	return 1
}

type Tuesday struct{}

func (t *Tuesday) Action(c *Context) {
	c.Today = &Tuesday{}
}

func (m *Tuesday) Val() int {
	return 2
}

type Wednesday struct{}

func (t *Wednesday) Action(c *Context) {
	c.Today = &Wednesday{}
}

func (m *Wednesday) Val() int {
	return 3
}

type Thursday struct{}

func (t *Thursday) Action(c *Context) {
	c.Today = &Thursday{}
}

func (m *Thursday) Val() int {
	return 4
}

type Friday struct{}

func (t *Friday) Action(c *Context) {
	c.Today = &Friday{}
}

func (m *Friday) Val() int {
	return 5
}

type Saturday struct{}

func (t *Saturday) Action(c *Context) {
	c.Today = &Saturday{}
}

func (m *Saturday) Val() int {
	return 6
}
