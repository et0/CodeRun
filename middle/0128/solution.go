package main // https://coderun.yandex.ru/problem/pedigree-number-of-descendants

// 1. сохраняем в мапу всех детей, в другую мапу с ключом родитель и значением в виде слайса (все дети)
// 2. находим родителя (корень дерева), которого нет в мапе всех детей
// 3. начиная с корня проходим по всем детям сохраняя количество потомков
// 4. сортируем

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
	nodes := make(map[string]int, n)
	relations := make(map[string][]string, n-1)
	for range n - 1 {
		fmt.Fscan(in, &children, &parent)

		nodes[children] = 0

		relations[parent] = append(relations[parent], children)
	}

	// ищем корень дерева - родитель, которого нет в общем списке (nodes)
	var root string
	for p := range relations {
		if _, ok := nodes[p]; !ok {
			root = p
			break
		}
	}

	setCountChildren(root, &nodes, &relations)

	result := make([]string, 0, n)
	for k, v := range nodes {
		result = append(result, k+" "+strconv.Itoa(v))
	}

	sort.Strings(result)

	for _, v := range result {
		fmt.Fprintln(out, v)
	}
}

func setCountChildren(parent string, nodes *map[string]int, relations *map[string][]string) int {
	counter := 0

	for _, children := range (*relations)[parent] {
		counter += 1 + setCountChildren(children, nodes, relations)
	}

	(*nodes)[parent] = counter

	return counter
}
