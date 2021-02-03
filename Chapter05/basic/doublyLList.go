package main

import "fmt"

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

var root = new(Node)

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{
			Val:  v,
			Prev: nil,
			Next: nil,
		}
		root = t
		return 0
	}

	if t.Val == v {
		fmt.Println("Node already exists:", v)
		return -1
	}

	if t.Next == nil {
		temp := t
		t.Next = &Node{
			Val:  v,
			Prev: temp,
			Next: nil,
		}
		return -2
	}

	return addNode(t.Next, v)
}

func traverse(t *Node) {

	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Val)
		t = t.Next
	}
	fmt.Println()

}

func reverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}

	for temp.Prev != nil {
		fmt.Printf("%d -> ", temp.Val)
		temp = temp.Prev
	}

	fmt.Printf("%d -> ", temp.Val)
	fmt.Println()
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}

	return i
}

func lookupNode(t *Node, v int) bool {
	if root == nil {
		return false
	}

	if v == t.Val {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, 1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 0)
	addNode(root, 0)
	traverse(root)
	addNode(root, 100)
	fmt.Println("Size:", size(root))
	traverse(root)
	reverse(root)
}
