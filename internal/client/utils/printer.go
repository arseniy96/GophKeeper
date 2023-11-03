package utils

import (
	"fmt"
	"strings"
	"time"
)

const (
	delay = 30 * time.Millisecond
)

type Printer struct{}

func (p *Printer) Print(s string) {
	slowPrint(s)
}

func (p *Printer) Scan(a ...interface{}) (int, error) {
	n, err := fmt.Scanln(a...)
	return n, err
}

func slowPrint(text string) {
	arr := strings.Split(text, "")
	for _, char := range arr {
		fmt.Print(char)
		time.Sleep(delay)
	}
	fmt.Println()
}
