package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	if l.Head == nil {
		l.Head.Data = data
	} else {
		currentNode := l.Head

		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}

		currentNode.Next.Data = data
	}
}