package main // https://coderun.yandex.ru/problem/cubes

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	anya := make([]int, n)
	boris := make([]int, m)

	for i := range n {
		fmt.Fscan(in, &anya[i])
	}

	for i := range m {
		fmt.Fscan(in, &boris[i])
	}

	sort.Ints(anya)
	sort.Ints(boris)

	notUniq := make([]int, 0, m+n)
	uniqAnya := make([]int, 0, n)
	uniqBoris := make([]int, 0, m)

	// 4 3
	// anaya = [0 1 10 9]
	// boris = [1 3 0]

	a, b := 0, 0
	for a < n && b < m {
		if anya[a] == boris[b] {
			notUniq = append(notUniq, anya[a])
			a++
			b++

			continue
		}

		if anya[a] > boris[b] {
			uniqBoris = append(uniqBoris, boris[b])
			b++
		} else {
			uniqAnya = append(uniqAnya, anya[a])
			a++
		}
	}
	uniqAnya = append(uniqAnya, anya[a:]...)
	uniqBoris = append(uniqBoris, boris[b:]...)

	fmt.Fprintln(out, len(notUniq))
	for _, v := range notUniq {
		fmt.Fprint(out, v, " ")
	}
	fmt.Fprint(out, "\n")

	fmt.Fprintln(out, len(uniqAnya))
	for _, v := range uniqAnya {
		fmt.Fprint(out, v, " ")
	}
	fmt.Fprint(out, "\n")

	fmt.Fprintln(out, len(uniqBoris))
	for _, v := range uniqBoris {
		fmt.Fprint(out, v, " ")
	}
}
