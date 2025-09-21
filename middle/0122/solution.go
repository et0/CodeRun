package main // https://coderun.yandex.ru/problem/second-maximum

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

	secondMaxNode := findSecondMaxNode(root)

	fmt.Fprintln(out, secondMaxNode.Val)
}

func findSecondMaxNode(root *Node) *Node {
	// Находим максимальный элемент и его родителя
	maxNode, prevMaxNode := findNextMax(root)

	// Если у максимального элемента нет левой ветки, то вторым максимальным будет его родитель
	if maxNode.Left == nil {
		return prevMaxNode
	}

	// Если у максимального элемента есть левая ветка, значит в ней находится второй максимум
	// Находим максимальный элемент у левой ветки
	secondMaxNode, _ := findNextMax(maxNode.Left)

	return secondMaxNode
}

// Функция возвращает максимальный элемент и его родителя
func findNextMax(root *Node) (*Node, *Node) {
	prev := root

	for {
		if root.Right == nil {
			break
		}

		prev = root
		root = root.Right
	}

	return root, prev
}
