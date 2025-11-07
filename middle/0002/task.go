package main // https://coderun.yandex.ru/problem/cheapest-way/description

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

	var (
		n, m  int
		value int
	)
	fmt.Fscan(in, &n, &m)

	route := make([][]int, n)
	for i := 0; i < n; i++ {
		route[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i != 0 && j != 0 {
				route[i][j] = min(route[i-1][j], route[i][j-1])
			} else if j != 0 {
				route[i][j] = route[i][j-1]
			} else if i != 0 {
				route[i][j] = route[i-1][j]
			}

			fmt.Fscan(in, &value)
			route[i][j] += value
		}
	}

	fmt.Fprintln(out, route[n-1][m-1])
}
