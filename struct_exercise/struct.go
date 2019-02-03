package main

type treeNode struct {
	value       int
	left, right *treeNode
}

func main() {
	//创建根节点
	root := treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{}

}
