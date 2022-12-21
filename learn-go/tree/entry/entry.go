package main

import (
	"fmt"
	"learn-go/tree"
)

type myTreeNode struct {
	*tree.Node // Embedding
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("This method is shadow.")
}

func main() {
	//var root tree.Node
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateTreeNode(2)

	root.Right.Left.SetValue(4)
	//root.right.left.print()
	//fmt.Println(root)
	////root.print()
	//
	//root.print()
	//root.setValue(100)

	root.Node.Traverse()
	fmt.Println()
	//myNode := myTreeNode{&root}
	myNode := root
	myNode.postOrder()
	fmt.Println()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count: ", nodeCount)
}
