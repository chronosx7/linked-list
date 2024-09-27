package linkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	// create list with single item
	// - length equals 1
	// - head is the value passed on construction
	t.Run("create list with single item", func(t *testing.T) {
		start_value := 7
		list := NewLinkedList(start_value)

		AssertListLength(t, list.Length(), 1)
		AssertNodeValue(t, list.Head(), start_value)
	})

	t.Run("list supports both INTs and STRINGs", func(t *testing.T) {
		list1 := NewLinkedList(1)
		list2 := NewLinkedList("a")

		list1.Add(9)
		list1.Add(8)
		list1.Add(7)

		list2.Add("x")
		list2.Add("y")
		list2.Add("z")

		AssertListLength(t, list1.Length(), 4)
		AssertListsAreEqual(t, ExtractListValues(list1), []int{1, 9, 8, 7})

		AssertListLength(t, list2.Length(), 4)
		AssertListsAreEqual(t, ExtractListValues(list2), []string{"a", "x", "y", "z"})

	})

	// add items to list
	// - length equals initialization + number of add operations
	// - tail is the value passed in the last add operation
	// - next and prev values match insert order
	t.Run("add items to list", func(t *testing.T) {
		list := NewLinkedList(7)
		list.Add(3)
		list.Add(5)
		AssertListLength(t, list.Length(), 3)
		AssertNodeValue(t, list.Head(), 7)
		AssertNodeValue(t, list.Tail(), 5)

		AssertNodeValue(t, list.Get(0).Next.Value, 3)
		AssertNodeValue(t, list.Get(1).Prev.Value, 7)
		AssertNodeValue(t, list.Get(1).Next.Value, 5)
		AssertNodeValue(t, list.Get(2).Prev.Value, 3)
		AssertNextIsNil(t, list.Get(2))
	})

	// pop item from list
	// - length equals total inserts - number of pop operations
	// - tail value equals the second-to-last item before pop operation
	// - next of new tail is nil
	t.Run("pop item from list", func(t *testing.T) {
		list := NewLinkedList(7)
		list.Add(3)
		list.Add(5)
		list.Pop()

		AssertListLength(t, list.Length(), 2)
		AssertNodeValue(t, list.Head(), 7)
		AssertNodeValue(t, list.Tail(), 3)
		AssertNextIsNil(t, list.Get(1))
	})

	// insert item between existing items
	// - length equals total add and insert operations
	// - inserted position has correct value
	// - prev and next of new item match existing items
	// - next of item on the left is the new one
	// - prev of item on the right is the new one
	t.Run("insert item between existing items", func(t *testing.T) {
		target_pos := 1
		new_val := 0
		list := NewLinkedList(7)
		list.Add(3)
		list.Add(5)
		list.Insert(new_val, target_pos)

		AssertListLength(t, list.Length(), 4)
		AssertListsAreEqual(t, ExtractListValues(list), []int{7, 0, 3, 5})
		AssertNodeValue(t, list.Get(target_pos).Value, new_val)
		AssertNodeValue(t, list.Get(target_pos).Prev.Value, 7)
		AssertNodeValue(t, list.Get(target_pos).Next.Value, 3)
		AssertNodeValue(t, list.Get(target_pos-1).Next.Value, new_val)
		AssertNodeValue(t, list.Get(target_pos+1).Prev.Value, new_val)
	})

	// insert item on the last position
	// - length equals total add and insert operations
	// - last item is "pushed to the right"
	// - next of new matches previous tail
	// - prev of previous tail matches new item
	t.Run("insert item on the last position", func(t *testing.T) {
		target_pos := 2
		new_val := 0
		list := NewLinkedList(7)
		list.Add(3)
		list.Add(5)
		list.Insert(new_val, target_pos)

		AssertListLength(t, list.Length(), 4)
		AssertNodeValue(t, list.GetVal(target_pos), new_val)
		AssertNodeValue(t, list.Get(target_pos-1).Next.Value, new_val)
		AssertNodeValue(t, list.Get(target_pos+1).Prev.Value, new_val)
		AssertNodeValue(t, list.Get(target_pos).Next.Value, 5)
		AssertNodeValue(t, list.Get(target_pos).Prev.Value, 3)
		AssertListsAreEqual(t, ExtractListValues(list), []int{7, 3, 0, 5})
	})

	// insert item at the end of the list
	// - length equals total add and insert operations
	// - prev new item matches last tail
	// - next new item is nil
	// - next of item on the left is the new one
	// - adding past the end does nothing
	t.Run("insert item at the end of the list", func(t *testing.T) {
		target_pos := 3
		new_val := 0
		list := NewLinkedList(7)
		list.Add(3)
		list.Add(5)
		list.Insert(new_val, target_pos)

		AssertListLength(t, list.Length(), 4)
		AssertListsAreEqual(t, ExtractListValues(list), []int{7, 3, 5, 0})
		AssertNodeValue(t, list.Get(target_pos).Value, new_val)
		AssertNodeValue(t, list.Get(target_pos).Prev.Value, 5)
		AssertNextIsNil(t, list.Get(target_pos))
		AssertNodeValue(t, list.Get(target_pos-1).Next.Value, new_val)

		list.Insert(new_val, target_pos+2)
		AssertListsAreEqual(t, ExtractListValues(list), []int{7, 3, 5, 0})
	})

	// insert item at the beginning of the list
	// - length equals total add and insert operations
	// - prev new item matches last tail
	// - next new item is nil
	// - next of item on the left is the new one
	t.Run("insert item at the beginning of the list", func(t *testing.T) {
		target_pos := 0
		new_val := 0
		list := NewLinkedList(7)
		list.Add(3)
		list.Add(5)
		list.Insert(new_val, target_pos)

		AssertListLength(t, list.Length(), 4)
		AssertListsAreEqual(t, ExtractListValues(list), []int{0, 7, 3, 5})
		AssertNodeValue(t, list.Get(target_pos).Value, new_val)
		AssertNodeValue(t, list.Get(target_pos).Next.Value, 7)
		AssertNodeIsNil(t, list.Get(target_pos).Prev)
	})
}

func AssertListLength(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %d items, got %d", want, got)
	}
}

func AssertNodeValue(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("value should be %d, got %d", want, got)
	}
}

func AssertNextIsNil[T any](t *testing.T, node *ListNode[T]) {
	t.Helper()

	if node.Next != nil {
		t.Errorf("item should be nil, got %v", node.Next)
	}
}

func AssertNodeIsNil[T any](t *testing.T, node *ListNode[T]) {
	t.Helper()

	if node != nil {
		t.Errorf("item should be nil, got %v", node.Next)
	}
}

func ExtractListValues[T any](list LinkedList[T]) []T {
	var values []T

	for index := range list.Length() {
		values = append(values, list.GetVal(index))
	}

	return values
}

func AssertListsAreEqual[T comparable](t *testing.T, got, want []T) {
	t.Helper()
	for index := range got {
		if got[index] != want[index] {
			t.Errorf("value mismatch at pos %d, got %v, want %v", index, got[index], want[index])
		}
	}
}
