package utils

import (
	"fmt"
	"strings"
	"time"
)

const (
	delay = 30 * time.Millisecond
)

func SlowPrint(text string) {
	arr := strings.Split(text, "")
	for _, char := range arr {
		fmt.Print(char)
		time.Sleep(delay)
	}
	fmt.Println()
}
