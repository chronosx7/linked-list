package linkedlist

type ListNode[T any] struct {
	Prev  *ListNode[T]
	Next  *ListNode[T]
	Value T
}

type LinkedList[T any] struct {
	Nodes []*ListNode[T]
}

func (l *LinkedList[T]) Length() int {
	return len(l.Nodes)
}

func (l *LinkedList[T]) Get(index int) *ListNode[T] {
	return l.Nodes[index]
}

func (l *LinkedList[T]) GetVal(index int) T {
	return l.Get(index).Value
}

func (l *LinkedList[T]) Head() T {
	return l.Nodes[0].Value
}

func (l *LinkedList[T]) HeadNode() *ListNode[T] {
	return l.Nodes[0]
}

func (l *LinkedList[T]) Tail() T {
	return l.Nodes[len(l.Nodes)-1].Value
}

func (l *LinkedList[T]) TailNode() *ListNode[T] {
	return l.Nodes[len(l.Nodes)-1]
}

func (l *LinkedList[T]) Add(new_value T) {
	new_node := ListNode[T]{Value: new_value, Prev: l.Nodes[len(l.Nodes)-1], Next: nil}
	l.Nodes = append(l.Nodes, &new_node)
	l.Nodes[len(l.Nodes)-2].Next = &new_node
}

func (l *LinkedList[T]) Insert(value T, position int) {
	new_node := ListNode[T]{Value: value}

	if position < 0 || position > l.Length() {
		return
	}

	if position == 0 { // Beginning of list
		l.Nodes = append([]*ListNode[T]{&new_node}, l.Nodes[:]...)
		l.Nodes[0].Next = l.Nodes[1]
		l.Nodes[1].Prev = l.Nodes[0]
	} else if position == l.Length() { // End of list
		l.Nodes = append(l.Nodes, &new_node)
		l.Nodes[position].Prev = l.Nodes[position-1]
		l.Nodes[position-1].Next = l.Nodes[position]
	} else { // Somewhere in between
		l.Nodes = append(l.Nodes[:position], append([]*ListNode[T]{&new_node}, l.Nodes[position:]...)...)
		l.Nodes[position-1].Next = l.Nodes[position]
		l.Nodes[position].Prev = l.Nodes[position-1]
		l.Nodes[position].Next = l.Nodes[position+1]
		l.Nodes[position+1].Prev = l.Nodes[position]
	}
}

func (l *LinkedList[T]) Pop() {
	l.Nodes = l.Nodes[:len(l.Nodes)-1]
	l.Nodes[len(l.Nodes)-1].Next = nil
}

func NewLinkedList[T any](init_value T) LinkedList[T] {
	first_node := ListNode[T]{Value: init_value, Prev: nil, Next: nil}
	list := LinkedList[T]{}
	list.Nodes = []*ListNode[T]{&first_node}

	return list
}
