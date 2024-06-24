package data_structures_examples

import "fmt"

// Node представляет узел в однонаправленном списке
type Node struct {
	value string
	next  *Node
}

// LinkedList представляет однонаправленный список
type LinkedList struct {
	head *Node
}

// Add добавляет новый узел в конец списка
func (list *LinkedList) Add(value string) {
	newNode := &Node{value: value}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// Remove удаляет узел с заданным значением
func (list *LinkedList) Remove(value string) bool {
	if list.head == nil {
		return false
	}

	if list.head.value == value {
		list.head = list.head.next
		return true
	}

	current := list.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}

	if current.next == nil {
		return false
	}

	current.next = current.next.next
	return true
}

// Find ищет узел с заданным значением
func (list *LinkedList) Find(value string) *Node {
	current := list.head
	for current != nil {
		if current.value == value {
			return current
		}
		current = current.next
	}
	return nil
}

// Print выводит все узлы списка
func (list *LinkedList) Print() {
	current := list.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
