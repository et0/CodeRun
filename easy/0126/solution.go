package main

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

	var value int
	var root *Node
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
			if value > cur.Val {
				if cur.Right == nil {
					cur.Right = &Node{Val: value}
					break
				}

				cur = cur.Right
			} else if value < cur.Val {
				if cur.Left == nil {
					cur.Left = &Node{Val: value}
					break
				}

				cur = cur.Left
			} else {
				break
			}
		}
	}

	stack := make([]int, 0, 1000)
	search(root, &stack)

	for _, v := range stack {
		fmt.Fprintln(out, v)
	}

}

func search(root *Node, stack *[]int) {
	if root == nil {
		return
	}

	// cначала уходим влево для поиска самого маленького числа
	search(root.Left, stack)

	if (root.Left != nil && root.Right == nil) || (root.Left == nil && root.Right != nil) {
		*stack = append(*stack, root.Val)
	}

	// потом вправо, для поиска большего числа
	search(root.Right, stack)
}
