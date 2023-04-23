package collections

type node struct {
	prev *node
	next *node
	data interface{}
}

type LinkedList struct {
	head *node
	tail *node
	size int
}

func (dll *LinkedList) InsertAtBeginning(data interface{}) {
	newnode := &node{data: data}
	if dll.head == nil {
		dll.head = newnode
		dll.tail = newnode
	} else {
		dll.head.prev = newnode
		newnode.next = dll.head
		dll.head = newnode
	}
	dll.size++
}

func (dll *LinkedList) InsertAtEnd(data interface{}) {
	newnode := &node{data: data}
	if dll.head == nil {
		dll.head = newnode
		dll.tail = newnode
	} else {
		dll.tail.next = newnode
		newnode.prev = dll.tail
		dll.tail = newnode
	}
	dll.size++
}

func (dll *LinkedList) InsertAtIndex(data interface{}, index int) error {
	if index < 0 || index > dll.size {
		panic("Index out of bounds")
	}
	if index == 0 {
		dll.InsertAtBeginning(data)
	} else if index == dll.size {
		dll.InsertAtEnd(data)
	} else {
		newnode := &node{data: data}
		currnode := dll.head
		for i := 0; i < index-1; i++ {
			currnode = currnode.next
		}
		newnode.prev = currnode
		newnode.next = currnode.next
		currnode.next.prev = newnode
		currnode.next = newnode
		dll.size++
	}
	return nil
}

func (dll *LinkedList) RemoveAtBeginning() error {
	if dll.head == nil {
		panic("List is empty")
	}
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.head = dll.head.next
		dll.head.prev = nil
	}
	dll.size--
	return nil
}

func (dll *LinkedList) RemoveAtEnd() error {
	if dll.head == nil {
		panic("List is empty")
	}
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	}
	dll.size--
	return nil
}

func (dll *LinkedList) RemoveAtIndex(index int) error {
	if index < 0 || index >= dll.size {
		panic("Index out of bounds")
	}
	if index == 0 {
		dll.RemoveAtBeginning()
	} else if index == dll.size-1 {
		dll.RemoveAtEnd()
	} else {
		currnode := dll.head
		for i := 0; i < index; i++ {
			currnode = currnode.next
		}
		currnode.prev.next = currnode.next
		currnode.next.prev = currnode.prev
		dll.size--
	}
	return nil
}

func (dll *LinkedList) GetFirst() interface{} {
	return dll.head.data
}

func (dll *LinkedList) GetLast() interface{} {
	return dll.tail.data
}

func (dll *LinkedList) GetAtIndex(index int) interface{} {
	if index < 0 || index >= dll.size {
		return nil
	}
	currnode := dll.head
	for i := 0; i < index; i++ {
		currnode = currnode.next
	}
	return currnode.data
}
