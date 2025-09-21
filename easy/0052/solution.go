package main // https://coderun.yandex.ru/problem/dictionary-synonyms

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	var word1, word2 string
	words := make(map[string]string, n)
	for range n {
		fmt.Fscan(in, &word1, &word2)

		words[word1] = word2
		words[word2] = word1
	}

	var find string
	fmt.Fscan(in, &find)

	fmt.Fprint(out, words[find])
}
