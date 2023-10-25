package utils

import (
	"fmt"
	"strings"
	"time"
)

func SlowPrint(text string) {
	arr := strings.Split(text, "")
	for _, char := range arr {
		fmt.Print(char)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println()
}
