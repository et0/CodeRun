package main // https://coderun.yandex.ru/problem/substring

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
		n, k int
		data string
	)
	fmt.Fscan(in, &n, &k)
	fmt.Fscan(in, &data)

	// максимальная длина подстроки
	maxCounter := 0
	// индекс начала максимальной подстроки
	indexStart := 0

	hash := make(map[byte]int)
	for right, left := 0, 0; right < n; right++ {
		hash[data[right]]++

		for ; hash[data[right]] > k && left < right; left++ {
			hash[data[left]]--
		}

		if maxCounter < right-left+1 {
			maxCounter = right - left + 1
			indexStart = left
		}
	}

	fmt.Fprintln(out, maxCounter, indexStart+1)
}
