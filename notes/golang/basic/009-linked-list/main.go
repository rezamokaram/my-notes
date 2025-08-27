package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(value int) {
	newNode := &Node{data: value}
	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (list *LinkedList) Display() {
	if list.head == nil {
		fmt.Println("The list is empty.")
		return
	}
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (list *LinkedList) Delete(value int) {
	if list.head == nil {
		fmt.Println("The list is empty.")
		return
	}

	if list.head.data == value {
		list.head = list.head.next
		return
	}

	current := list.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}

	if current.next == nil {
		fmt.Println("Value not found in the list.")
		return
	}

	current.next = current.next.next
}

func (list *LinkedList) Search(value int) bool {
	current := list.head
	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}

func main() {
	list := &LinkedList{}

	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)

	fmt.Print("Initial List: ")
	list.Display()

	fmt.Println("Searching for 30:", list.Search(30))   // Output: true
	fmt.Println("Searching for 100:", list.Search(100)) // Output: false

	fmt.Println("Deleting 20...")
	list.Delete(20)
	list.Display()

	fmt.Println("Deleting 100...")
	list.Delete(100)

	fmt.Print("Final List: ")
	list.Display()
}
