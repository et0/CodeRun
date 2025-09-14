package main // https://coderun.yandex.ru/problem/pedigree-counting-levels

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	var children, parent string
	tree := make(map[string]string, n-1)
	for range n {
		fmt.Fscan(in, &children, &parent)
		tree[children] = parent
	}

	var root string
	data := make([]string, 0, n)
	for k := range tree {
		if root == "" {
			root = getRoot(k, &tree) + " 0"
		}
		data = append(data, k+" "+strconv.Itoa(getLvl(k, 0, &tree)))
	}
	data = append(data, root)

	sort.Strings(data)

	for _, v := range data {
		fmt.Fprintln(out, v)
	}
}

func getRoot(children string, tree *map[string]string) string {
	parent, ok := (*tree)[children]
	if !ok {
		return children
	}

	return getRoot(parent, tree)
}

func getLvl(children string, lvl int, tree *map[string]string) int {
	parent, ok := (*tree)[children]
	if !ok {
		return lvl
	}

	return getLvl(parent, lvl+1, tree)
}
