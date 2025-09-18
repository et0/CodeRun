package main // https://coderun.yandex.ru/problem/fork-conclusion

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Node struct {
	Val         int
	Left, Right *Node
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var (
		value int
		root  *Node
	)
	for {
		fmt.Fscan(in, &value)

		if value == 0 {
			break
		}

		if root == nil {
			root = &Node{Val: value}
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
				break
			}
		}
	}

	stack := make([]int, 0, 1000)
	search(root, &stack)

	sort.Ints(stack)

	for _, v := range stack {
		fmt.Fprintln(out, v)
	}
}

func search(root *Node, stack *[]int) {
	if root == nil {
		return
	}

	if root.Left != nil && root.Right != nil {
		*stack = append(*stack, root.Val)
	}

	search(root.Left, stack)
	search(root.Right, stack)
}
