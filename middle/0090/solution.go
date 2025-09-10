package main // https://coderun.yandex.ru/problem/stylish-clothes

import (
	"bufio"
	"fmt"
	"os"
)

func abs(number int) int {
	if number < 0 {
		return number * -1
	}

	return number
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n)
	tshirts := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &tshirts[i])
	}

	fmt.Fscan(in, &m)
	pants := make([]int, m)
	for i := range m {
		fmt.Fscan(in, &pants[i])
	}

	counter := 10_000_001
	outT, outP := 0, 0
	for t, p := 0, 0; t < n && p < m; {
		if tshirts[t] == pants[p] {
			outT, outP = t, p
			break
		}

		diff := abs(tshirts[t] - pants[p])
		if counter > diff {
			outT, outP = t, p
			counter = diff
		}

		if tshirts[t] > pants[p] {
			p++
		} else {
			t++
		}
	}

	fmt.Fprintf(out, "%d %d\n", tshirts[outT], pants[outP])
}
