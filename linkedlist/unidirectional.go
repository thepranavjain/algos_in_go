package linkedlist

type UnidirectionalLLNode struct {
	data string
	next *UnidirectionalLLNode
}

type UnidirectionalLL struct {
	head *UnidirectionalLLNode
	tail *UnidirectionalLLNode
}

func NewUnidirectionalLL(initData string) *UnidirectionalLL {
	firstNode := &UnidirectionalLLNode{
		data: initData,
		next: nil,
	}
	return &UnidirectionalLL{
		head: firstNode,
		tail: firstNode,
	}
}

func (li *UnidirectionalLL) Push(data string) {
	newNode := &UnidirectionalLLNode{
		data: data,
		next: nil,
	}

	if li.head == nil {
		li.head = newNode
	} else {
		li.tail.next = newNode
	}
	li.tail = newNode
}

func (li *UnidirectionalLL) Insert(index int, data string) {
	newNode := &UnidirectionalLLNode{
		data: data,
	}
	if index == 0 {
		newNode.next = li.head
		li.head = newNode
		if newNode.next == nil {
			li.tail = newNode
		}
	} else {
		p := li.head
		j := 1
		for p != nil {
			if p.next == nil && index > j {
				panic("Index out of range")
			}
			if index == j {
				newNode.next = p.next
				p.next = newNode
				break
			} else {
				p = p.next
				j++
			}
		}
	}
}
