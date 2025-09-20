package main // https://coderun.yandex.ru/problem/avl-balance

import (
	"bufio"
	"fmt"
	"os"
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

	if isAVL(root) {
		fmt.Fprint(out, "YES")
	} else {
		fmt.Fprint(out, "NO")
	}
}

func isAVL(root *Node) bool {
	return getHeight(root) != -1
}

func getHeight(node *Node) int {
	if node == nil {
		return 0
	}

	left := getHeight(node.Left)
	if left == -1 {
		return -1
	}

	right := getHeight(node.Right)
	if right == -1 {
		return -1
	}

	if abs(left-right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func abs(number int) int {
	if number < 0 {
		return -number
	}

	return number
}
