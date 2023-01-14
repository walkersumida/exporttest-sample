package counter

type Counter struct {
	n int
}

func (c *Counter) Count() {
	c.n++
}
func (c *Counter) reset() {
	c.n = 0
}
