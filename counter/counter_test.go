package counter_test

import (
	"testing"

	"github.com/walkersumida/exporttest-sample/counter"
)

func TestReset(t *testing.T) {
	c := counter.Counter{}
	c.ExportSetN(1)

	want := 0
	counter.ExportReset(&c)

	got := c.ExportN()
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
