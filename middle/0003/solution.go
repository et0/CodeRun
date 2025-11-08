package main // https://coderun.yandex.ru/problem/print-the-route-of-the-maximum-cost

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		n, m  int
		value int
	)
	fmt.Fscan(in, &n, &m)

	route := make([][]int, n)
	for i := range n {
		route[i] = make([]int, m)
		for j := range m {
			if i != 0 && j != 0 {
				route[i][j] = max(route[i-1][j], route[i][j-1])
			} else if i != 0 {
				route[i][j] = route[i-1][j]
			} else if j != 0 {
				route[i][j] = route[i][j-1]
			}

			fmt.Fscan(in, &value)
			route[i][j] += value
		}
	}

	fmt.Fprintln(out, route[n-1][m-1])

	move := make([]byte, 0, n*m)
	for n, m = n-1, m-1; n >= 0 && m >= 0; {
		if n != 0 && m != 0 {
			if route[n-1][m] > route[n][m-1] {
				move = append(move, 'D')
				n--
			} else {
				move = append(move, 'R')
				m--
			}
		} else if n != 0 {
			move = append(move, 'D')
			n--
		} else if m != 0 {
			move = append(move, 'R')
			m--
		} else {
			break
		}
	}

	for i := len(move) - 1; i >= 0; i-- {
		fmt.Fprint(out, string(move[i]), " ")
	}
}
