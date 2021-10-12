package binarygap

import (
	"strconv"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	// write your code in Go 1.4
	n := int64(N)
	max := 0
	curr := 0
	cal := false
	binaryString := strconv.FormatInt(n, 2)
	for _, char := range binaryString {
		if string(char) == "1" {
			if !cal {
				cal = true
			} else {
				if curr > max {
					max = curr
				}
				curr = 0
			}

		} else {
			if cal {
				curr++
			}
		}
	}
	return max
}
