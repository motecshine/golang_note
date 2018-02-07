package tree

var root *TreeNode

type TreeNode struct {
	data      int
	leftNode  *TreeNode
	rightNode *TreeNode
}

type TreeController struct {
}

// 初始化node节点
func (t *TreeController) init(value int) *TreeController {
	root = &TreeNode{
		data:      value,
		leftNode:  nil,
		rightNode: nil,
	}
	return t
}

func (t *TreeController) insert(value int) *TreeController {
	newNode := &TreeNode{
		data:      value,
		leftNode:  nil,
		rightNode: nil,
	}
	for {
		// 第一个叶子
		if root.leftNode == nil && root.rightNode == nil {
			if root.data < value {
				root.leftNode = newNode
			} else {
				root.rightNode = newNode
			}
		} else {
			if root.data < value {
				traverseAndInsert(root.leftNode, newNode)
			} else {
				traverseAndInsert(root.rightNode, newNode)
			}
		}
	}
}

func traverseAndInsert(list, newNode *TreeNode) {
	for {
		if list.data > newNode.data {
			if list.leftNode == nil {
				list.leftNode = newNode
				break
			}
			list = list.leftNode
		}

		if list.data < newNode.data {
			if list.rightNode == nil {
				list.rightNode = newNode
				break
			}
			list = list.rightNode
		}
	}
}
