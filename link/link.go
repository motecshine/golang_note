package link

import "log"

type Node struct {
	data        interface{}
	next        *Node
	isFisrtNode bool
}

var Link *Node

func Pop(link *Node) *Node {
	lastNode := getLastNode(link)
	lastNode.next = nil
	return link
}

func Push(data interface{}, link *Node) *Node {
	oldFisrtNode := getFisrtNode(link)
	oldFisrtNode.isFisrtNode = false
	newNode := &Node{
		data:        data,
		next:        oldFisrtNode,
		isFisrtNode: true,
	}
	return newNode
}

func getLastNode(link *Node) *Node {
	tmp := link
	for {
		if tmp.next == nil {
			return tmp
		}
		tmp = tmp.next
	}
}

func getFisrtNode(link *Node) *Node {
	tmp := link
	for {
		if tmp.isFisrtNode == true {
			return tmp
		}
		tmp = tmp.next
	}
}

func InitStack() *Node {
	Link = &Node{
		data:        nil,
		next:        nil,
		isFisrtNode: true,
	}

	return Link
}

func PrintLink(link *Node) {
	tmp := link
	for {
		if tmp.next != nil {
			log.Println("print: ", tmp)
			tmp = tmp.next
		} else {
			break
		}
	}
}
