package calc

func sum(x, y int) int {
	return x + y
}

func join(x, y string) string {
	return x + y
}

type Counter struct {
	n int
}

func (c *Counter) Count() {
	c.n++
}
func (c *Counter) reset() {
	c.n = 0
}
