package main // https://coderun.yandex.ru/problem/leaf-conclusion/description

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

	var value, counter int
	for {
		fmt.Fscan(in, &value)
		if value == 0 {
			break
		}

		counter++

		if root.Val == 0 {
			root.Val = value
			continue
		}

		for cur := root; ; {
			if value < cur.Val {
				if cur.Left == nil {
					cur.Left = &Node{Val: value}
					break
				}

				cur = cur.Left
			} else if value > cur.Val {
				if cur.Right == nil {
					cur.Right = &Node{Val: value}
					break
				}

				cur = cur.Right
			} else {
				counter--
				break
			}
		}
	}

	stack := make([]*Node, 0, counter)
	for {
		if root.Left != nil {
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			root = root.Left
		} else if root.Right != nil {
			root = root.Right
		} else {
			fmt.Fprintln(out, root.Val)

			size := len(stack)
			if size == 0 {
				break
			}
			root = stack[size-1]
			stack = stack[:size-1]
		}
	}
}
