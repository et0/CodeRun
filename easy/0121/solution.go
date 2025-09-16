package main // https://coderun.yandex.ru/problem/depth-added-elements

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	root := &Node{}

	var value int
	for {
		fmt.Fscan(in, &value)

		if value == 0 {
			break
		}

		lvl := 1
		for cur := root; ; {
			if cur.Val == 0 {
				cur.Val = value
				fmt.Fprintf(out, "%d ", lvl)
				break
			}

			lvl++

			if value > cur.Val {
				if cur.Right == nil {
					cur.Right = &Node{Val: value}
					fmt.Fprintf(out, "%d ", lvl)
					break
				}

				cur = cur.Right
			} else if value < cur.Val {
				if cur.Left == nil {
					cur.Left = &Node{Val: value}
					fmt.Fprintf(out, "%d ", lvl)
					break
				}

				cur = cur.Left
			} else {
				break
			}
		}
	}
}
