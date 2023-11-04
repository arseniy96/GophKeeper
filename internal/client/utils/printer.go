package utils

import (
	"fmt"
	"strings"
	"time"
)

const (
	delay = 30 * time.Millisecond
)

// Printer – структура для работы с текстом.
type Printer struct{}

// Print – метод вывода текстовых данных на экран.
func (p *Printer) Print(s string) {
	slowPrint(s)
}

// Scan – обёртка для fmt.Scan.
func (p *Printer) Scan(a ...interface{}) (int, error) {
	n, err := fmt.Scanln(a...)
	return n, err
}

func slowPrint(text string) {
	arr := strings.Split(text, "")
	for _, char := range arr {
		fmt.Print(char)
		<-time.After(delay)
	}
	fmt.Println()
}
