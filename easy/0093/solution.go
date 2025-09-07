package main // https://coderun.yandex.ru/problem/city-of-che

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, r int
	fmt.Fscan(in, &n, &r)

	dist := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &dist[i])
	}

	counter := 0
	for left, right := 0, 0; left < n; left++ {
		for ; right < n && dist[right]-dist[left] <= r; right++ {

		}

		counter += n - right
	}

	fmt.Fprintln(out, counter)
}
