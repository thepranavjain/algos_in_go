package linkedlist

import (
	"testing"
)

func TestUnidirectionalLLPush(t *testing.T) {
	li := NewUnidirectionalLL("lorem")
	li.Push("ipsum")
	if li.head.data != "lorem" || li.head.next == nil {
		t.Error("Expected head data: 'lorem', next: non nil. Received head:", li.head)
	}
	if li.head.next != li.tail {
		t.Error("Expected head next to be same as tail. Received Tail:", li.tail, "li.head.next", li.head.next, "are not equal")
	}
	if li.tail.data != "ipsum" || li.tail.next != nil {
		t.Error("Expected tail data: ipsum, next: nil. Received tail:", li.tail)
	}

	remaining := []string{"dolor", "sit", "amet"}
	all := append([]string{"lorem", "ipsum"}, remaining...)
	for _, v := range remaining {
		li.Push(v)
	}

	pointer := li.head
	index := 0
	for pointer != nil {
		if pointer.data != all[index] {
			t.Errorf("At index %d. Expected %s. Received %s", index, all[index], pointer.data)
		}
		index++
		pointer = pointer.next
	}
}

func TestUnidirectionalLLInsert(t *testing.T) {
	ar := []string{"lorem", "ipsum", "dolor"}
	li := &UnidirectionalLL{}
	for _, v := range ar {
		li.Push(v)
	}

	ar = []string{"lorem", "ipsum", "johndoe", "dolor"}
	li.Insert(2, "johndoe")
	p := li.head
	i := 0
	for p != nil {
		if p.data != ar[i] {
			t.Errorf("At index %d, expected %s, received %s", i, ar[i], p.data)
		}
		p = p.next
		i++
	}

	ar = []string{"lorem", "ipsum", "johndoe", "dolor", "sit"}
	li.Insert(4, "sit")
	p = li.head
	i = 0
	for p != nil {
		if p.data != ar[i] {
			t.Errorf("At index %d, expected %s, received %s", i, ar[i], p.data)
		}
		p = p.next
		i++
	}

	ar = []string{"abc", "lorem", "ipsum", "johndoe", "dolor", "sit"}
	li.Insert(0, "abc")
	p = li.head
	i = 0
	for p != nil {
		if p.data != ar[i] {
			t.Errorf("At index %d, expected %s, received %s", i, ar[i], p.data)
		}
		p = p.next
		i++
	}

	li2 := &UnidirectionalLL{}
	li2.Insert(0, "abc")
	if li2.head != li2.tail || li2.head.data != "abc" || li2.head.next != nil {
		t.Error("Error in inserting at 0th index of empty Uni-D LL")
	}
}
