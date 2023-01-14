package counter

var ExportReset = (*Counter).reset

func (c *Counter) ExportN() int {
	return c.n
}
func (c *Counter) ExportSetN(n int) {
	c.n = n
}
