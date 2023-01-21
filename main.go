package main

import (
	"fmt"

	"github.com/walkersumida/exporttest-sample/calc"
	"github.com/walkersumida/exporttest-sample/counter"
)

func main() {
	cnter := counter.Counter{}
	cnter.Count()
	fmt.Println(cnter.Get())

	ave := calc.Ave(10, 2)
	fmt.Println(ave)
}
