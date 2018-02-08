package tree

import "fmt"

var root *TreeNode

type TreeNode struct {
	data      int
	leftNode  *TreeNode
	rightNode *TreeNode
}

type TreeController struct {
}

// 初始化node节点
func (t *TreeController) Init(value int) *TreeController {
	root = &TreeNode{
		data:      value,
		leftNode:  nil,
		rightNode: nil,
	}
	return t
}

func (t *TreeController) Insert(value int) *TreeController {
	newNode := &TreeNode{
		data:      value,
		leftNode:  nil,
		rightNode: nil,
	}
	TraverseAndInsert(root, newNode)
	fmt.Println(root, root.leftNode, root.rightNode)
	return t
}

func TraverseAndInsert(list, newNode *TreeNode) {
	if list.data > newNode.data {
		for {
			if list.leftNode == nil {
				list.leftNode = newNode
				break
			}
			list = list.leftNode
		}
	}

	if list.data < newNode.data {
		for {
			if list.rightNode == nil {
				list.rightNode = newNode
				break
			}
			list = list.rightNode
		}
	}

	root = list
}
