package Doublelist

import (
	"fmt"
)

type Node struct {
	prev *Node
	data int
	next *Node
}

type List struct {
	head   *Node
	tail   *Node
	lenght int
}

func Create() *List {
	new_list := &List{head: nil, tail: nil, lenght: 0}
	return new_list
}

func Append(list *List, value int) {
	new_node := &Node{prev: nil, data: value, next: nil}

	if list.head == nil {
		list.lenght++
		list.head = new_node
		list.tail = new_node
	} else {
		current := list.tail
		list.tail.next = new_node
		list.tail = new_node
		new_node.prev = current
		list.lenght++

	}

}

func Display(list *List) {
	current := list.head
	fmt.Printf("[")
	for current != nil {
		if current == list.tail {
			fmt.Printf("%d", current.data)
			break
		}
		fmt.Printf("%d,", current.data)
		current = current.next
	}
	fmt.Println("]")
	fmt.Printf("Lenght -> %d\n", list.lenght)
}

func ReverseDisplay(list *List) {
	current := list.tail
	fmt.Printf("[")
	for current != nil {
		if current == list.head {
			fmt.Printf("%d", current.data)
			break
		}
		fmt.Printf("%d,", current.data)
		current = current.prev
	}
	fmt.Println("]")
}

func Get(list *List, index int) int {
	if list.head == nil {
		fmt.Println("Empty list")
		return -1
	}
	count := 0
	for current := list.head; current != nil; current = current.next {
		if count == index {
			return current.data
		}
		count++
	}
	fmt.Println("ERRO INVALID INDEX")
	return -1
}

func Remove(list *List, index int) int {
	if list.head == nil {
		fmt.Println("Erro empty list")
		return -1
	}
	if index == 0 {
		current := list.head
		list.head = current.next
		list.lenght--
		return current.data
	}

	if index == list.lenght-1 {
		value := list.tail.data
		list.tail = list.tail.prev
		list.tail.next = nil
		list.lenght--
		return value
	}

	count := 0
	current := list.head
	for current != nil {
		if count == index {
			current.prev.next = current.next
			current.next.prev = current.prev
			list.lenght--
			return current.data
		}

		count++
		current = current.next
	}
	fmt.Println("INVALID INDEX")
	return -1
}

func Set(list *List, index int, value int) {
	if list.head == nil {
		fmt.Println("ERRO EMPTY LIST")
		return
	}

	if index > list.lenght/2 && index <= list.lenght-1 {
		count := list.lenght - 1
		current := list.tail
		for count > list.lenght/2 {
			if count == index {
				current.data = value
				return
			}
			count--
			current = current.prev

		}
	} else if index <= list.lenght/2 && index >= 0 {
		count := 0
		current := list.head
		for count <= list.lenght/2 {
			if count == index {
				current.data = value
				return
			}
			count++
			current = current.next
		}
	} else {
		fmt.Println("Invalid index")
		return
	}

}

func Insert(list *List, index int, value int) {
	new_node := &Node{prev: nil, data: value, next: nil}
	if list.lenght == 0 || list.head == nil {
		list.lenght++
		list.head = new_node
		list.tail = new_node
		return
	}

	if index == 0 {
		list.head.prev = new_node
		new_node.next = list.head
		list.head = new_node
		list.lenght++
		return
	}

	count := 0
	current := list.head
	for current != nil {
		if count == index {
			new_node.prev = current.prev
			new_node.next = current
			current.prev.next = new_node
			current.prev = new_node
			list.lenght++
			return
		}
		count++
		current = current.next
	}

	fmt.Println("Invalid index")

}
