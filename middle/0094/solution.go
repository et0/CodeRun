package main // https://coderun.yandex.ru/problem/beauty-above-all

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

	var n, k int
	fmt.Fscan(in, &n, &k)

	trees := make([]int, n)
	for i := range trees {
		fmt.Fscan(in, &trees[i])
	}

	minLength := n              // минимальная длина должна быть во весь слайс
	outLeft, outRight := 0, n-1 // важно задать границы от начала и до конца строки
	kinds := make(map[int]int)
	for left, right := 0, 0; right < n; right++ {
		kinds[trees[right]]++

		for ; left <= right && len(kinds) == k; left++ {
			if right-left+1 < minLength {
				minLength = right - left + 1
				outLeft, outRight = left, right
			}

			kinds[trees[left]]--
			if kinds[trees[left]] == 0 {
				delete(kinds, trees[left])
			}
		}
	}

	fmt.Fprintf(out, "%d %d", outLeft+1, outRight+1)
}
