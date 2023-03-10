package calc_test

import (
	"testing"

	"github.com/walkersumida/exporttest-sample/calc"
)

func TestSum(t *testing.T) {
	x := 1
	y := 2
	want := 3
	if got := calc.ExportSum(x, y); got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
