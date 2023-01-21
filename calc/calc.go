package calc

func sum(x, y int) int {
	return x + y
}

func Ave(x, y int) int {
	return (x + y) / 2
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
