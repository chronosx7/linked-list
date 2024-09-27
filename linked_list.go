package linkedlist

type ListNode struct {
	Prev  *ListNode
	Next  *ListNode
	Value int
}

type LinkedList struct {
	Nodes []*ListNode
}

func (l *LinkedList) Length() int {
	return len(l.Nodes)
}

func (l *LinkedList) Get(index int) *ListNode {
	return l.Nodes[index]
}

func (l *LinkedList) GetVal(index int) int {
	return l.Get(index).Value
}

func (l *LinkedList) Head() int {
	return l.Nodes[0].Value
}

func (l *LinkedList) HeadNode() *ListNode {
	return l.Nodes[0]
}

func (l *LinkedList) Tail() int {
	return l.Nodes[len(l.Nodes)-1].Value
}

func (l *LinkedList) TailNode() *ListNode {
	return l.Nodes[len(l.Nodes)-1]
}

func (l *LinkedList) Add(new_value int) {
	new_node := ListNode{Value: new_value, Prev: l.Nodes[len(l.Nodes)-1], Next: nil}
	l.Nodes = append(l.Nodes, &new_node)
	l.Nodes[len(l.Nodes)-2].Next = &new_node
}

func (l *LinkedList) Insert(value, position int) {
	new_node := ListNode{Value: value}

	if position < 0 || position > l.Length() {
		return
	}

	if position == 0 { // Beginning of list
		l.Nodes = append([]*ListNode{&new_node}, l.Nodes[:]...)
		l.Nodes[0].Next = l.Nodes[1]
		l.Nodes[1].Prev = l.Nodes[0]
	} else if position == l.Length() { // End of list
		l.Nodes = append(l.Nodes, &new_node)
		l.Nodes[position].Prev = l.Nodes[position-1]
		l.Nodes[position-1].Next = l.Nodes[position]
	} else { // Somewhere in between
		l.Nodes = append(l.Nodes[:position], append([]*ListNode{&new_node}, l.Nodes[position:]...)...)
		l.Nodes[position-1].Next = l.Nodes[position]
		l.Nodes[position].Prev = l.Nodes[position-1]
		l.Nodes[position].Next = l.Nodes[position+1]
		l.Nodes[position+1].Prev = l.Nodes[position]
	}
}

func (l *LinkedList) Pop() {
	l.Nodes = l.Nodes[:len(l.Nodes)-1]
	l.Nodes[len(l.Nodes)-1].Next = nil
}

func NewLinkedList(init_value int, circular bool) LinkedList {
	first_node := ListNode{Value: init_value, Prev: nil, Next: nil}
	list := LinkedList{}
	list.Nodes = []*ListNode{&first_node}

	return list
}
