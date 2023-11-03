package utils

import (
	"fmt"
)

type Printer struct{}

func (p *Printer) Print(s string) {
	SlowPrint(s)
}

func (p *Printer) Scan(a ...interface{}) (int, error) {
	n, err := fmt.Scanln(a...)
	return n, err
}
