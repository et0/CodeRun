package main // https://coderun.yandex.ru/problem/nearest-number

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	data := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &data[i])
	}

	var x int
	fmt.Fscan(in, &x)

	if len(data) == 0 {
		fmt.Fprintln(out, -1)
		return
	}

	minIdx := 0

	// -10 5 3 11 -20
	for i := 1; i < n; i++ {
		if abs(x-data[i]) < abs(x-data[minIdx]) {
			minIdx = i
		}
	}
	fmt.Println(data[minIdx])
}
