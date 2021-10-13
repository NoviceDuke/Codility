package cyclicrotation

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K int) []int {
	if K == len(A) {
		return A
	}
	tmp := make([]int, len(A))
	copy(tmp, A)
	shift := K
	for i := 0; i < len(A); i++ {
		A[(i+shift)%len(A)] = tmp[i]
	}
	return A
}
