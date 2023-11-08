package hard

import (
	"fmt"
	"math"
)

type BNode struct {
	data  int
	left  *BNode
	right *BNode
}

type BinaryTree struct {
	root *BNode
}

func (bt *BinaryTree) Add(data int) {

	newNode := &BNode{data: data, left: nil, right: nil}

	if bt.root == nil {
		bt.root = newNode
		return
	}

	queue := []*BNode{bt.root}
	for len(queue) > 0 {
		current := queue[0]

		if current.left == nil {
			current.left = newNode
			return
		}
		if current.right == nil {
			current.right = newNode
			return
		}

		queue = queue[1:]
		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}

	}

}

func (bt *BinaryTree) Count() int {
	count := 1

	visited := []*BNode{bt.root}
	for len(visited) > 0 {
		current := visited[0]

		visited = visited[1:]
		if current.left != nil {
			visited = append(visited, current.left)
			count++
		}
		if current.right != nil {
			visited = append(visited, current.right)
			count++
		}
	}
	return count
}

func calculateLevels(nodes int) int {
	if nodes <= 0 {
		return 0
	}

	// height := int(math.Log2(float64(nodes+1))) - 1
	levels := int(math.Log2(float64(nodes)) + 1)
	return levels
}

func (bt *BinaryTree) GetNode(node int) *BNode {

	nodeFound := false
	queue := []*BNode{bt.root}
	var current *BNode = nil

	for len(queue) > 0 && !nodeFound {
		current = queue[0]
		if current.data == node {
			current = queue[0]
			break
		}

		queue = queue[1:]
		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}

	}
	return current
}

func (bt *BinaryTree) Swap(a, b int) {

	node1 := bt.GetNode(a)
	node2 := bt.GetNode(b)

	if node1 == nil || node2 == nil {
		fmt.Println("Invalid nodes")
		return
	}

	node1.data, node2.data = node2.data, node1.data
	node1.left, node2.left = node2.left, node1.left
	node1.right, node2.right = node2.right, node1.right

	levels := calculateLevels(bt.Count())
	bt.Print(levels)

}

func (bt *BinaryTree) Print(levels int) {

	q := []*BNode{bt.root}
	data := []int{}
	if bt.root == nil {
		fmt.Println("No binary tree found!")
		return
	}
	for len(q) > 0 {
		current := q[0]

		data = append(data, current.data)

		q = q[1:]
		if current.left != nil {
			q = append(q, current.left)
		}

		if current.right != nil {
			q = append(q, current.right)
		}
	}

	for i := 1; i <= levels; i++ {
		nodes := int(math.Pow(2, float64(i-1)))
		bt.PrintNodes(data, nodes)
	}

}

func (bt *BinaryTree) PrintNodes(data []int, nodes int) {
	limit := (nodes * 2) - 1
	if limit > len(data)-1 {
		limit = len(data)
	}
	slice := data[nodes-1 : limit]

	for i, v := range slice {
		fmt.Printf("%v ", v)

		if i == len(slice)-1 {
			fmt.Println()
		}
	}
}

func SwapNodes() {

	// Challenge 9: Swap Nodes in a Binary Tree Using Pointers
	// Description: Write a Go function that swaps two nodes in a binary tree using pointers.

	// Input:
	// 	    1
	//     / \
	//    2   3
	//   / \
	//  4   5
	// Swap nodes 2 and 3.

	// Output:
	// 	  1
	//   / \
	//  3   2
	// 	   / \
	// 	  4   5

	// Incorrect Output:
	// 	   1
	//    / \
	//   2   3
	//  / \
	// 4   5

	binaryTree := BinaryTree{}
	binaryTree.Add(1)
	binaryTree.Add(2)
	binaryTree.Add(3)
	binaryTree.Add(4)
	binaryTree.Add(5)
	binaryTree.Add(6)
	binaryTree.Add(7)
	binaryTree.Add(8)
	binaryTree.Add(9)
	binaryTree.Add(10)
	binaryTree.Add(11)
	binaryTree.Add(12)
	binaryTree.Add(13)
	binaryTree.Add(14)
	binaryTree.Add(15)
	binaryTree.Add(16)

	levels := calculateLevels(binaryTree.Count())
	binaryTree.Print(levels)

	binaryTree.Swap(2, 3)

}
