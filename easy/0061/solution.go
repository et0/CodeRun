package main // https://coderun.yandex.ru/problem/intersection-sets

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	first := getSliceFromLine(in)
	second := getSliceFromLine(in)

	notUniq := make([]int, 0, len(*first)+len(*second))
	for one, two := 0, 0; one < len(*first) && two < len(*second); {
		if (*first)[one] == (*second)[two] {
			notUniq = append(notUniq, (*first)[one])

			one++
			two++

			continue
		}

		if (*first)[one] > (*second)[two] {
			two++
		} else {
			one++
		}
	}

	for _, v := range notUniq {
		fmt.Fprint(out, v, " ")
	}
}

func getSliceFromLine(in *bufio.Reader) *[]int {
	line, _ := in.ReadString('\n')
	setString := strings.Split(strings.TrimSpace(line), " ")

	data := make([]int, len(setString))
	conversToSortInt(&setString, &data)

	return &data
}

func conversToSortInt(setString *[]string, data *[]int) {
	for i, v := range *setString {
		number, _ := strconv.Atoi(v)
		(*data)[i] = number
	}
	sort.Ints(*data)
}
