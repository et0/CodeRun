package main // https://coderun.yandex.ru/problem/tree-height

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var value int
	root := &Node{}
	maxLvl := 0
	for {
		fmt.Fscan(in, &value)
		if value == 0 {
			break
		}

		lvl := 1
		for cur := root; ; {
			if cur.Value == 0 {
				cur.Value = value
				break
			}

			// увеличиваем высоту
			lvl++

			if value > cur.Value {
				if cur.Right != nil {
					cur = cur.Right
					continue
				}

				cur.Right = &Node{Value: value}
				break
			} else if value < cur.Value {
				if cur.Left != nil {
					cur = cur.Left
					continue
				}

				cur.Left = &Node{Value: value}
				break
			} else {
				lvl = 0
				break
			}
		}

		if lvl > maxLvl {
			maxLvl = lvl
		}
	}

	fmt.Fprintln(out, maxLvl)
}
