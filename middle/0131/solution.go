package main // https://coderun.yandex.ru/problem/beautiful-line

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
		k int
		s string
	)
	fmt.Fscan(in, &k)
	fmt.Fscan(in, &s)

	size := len(s)
	if size < k {
		fmt.Fprintln(out, len(s))

		return
	}

	var letter byte = 'a'
	maxRepeat := 0
	for ; letter <= 'z'; letter++ {
		swap := 0
		for left, right := 0, 0; right < size; right++ {
			if s[right] != letter {
				swap++
			}

			for ; left < right && swap > k; left++ {
				if s[left] != letter {
					swap--
				}
			}

			if right-left > maxRepeat {
				maxRepeat = right - left
			}
		}
	}

	fmt.Fprintln(out, maxRepeat+1)
}
