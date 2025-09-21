package main // https://coderun.yandex.ru/problem/number-words-text

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var word string
	words := make(map[string]struct{})
	for {
		_, err := fmt.Fscan(in, &word)
		if errors.Is(err, io.EOF) {
			break
		}

		words[word] = struct{}{}
	}

	fmt.Fprint(out, len(words))
}
