package main

import "fmt"

type Node struct {
	number int
	prev   *Node
	next   *Node
}
type temp struct {
	head  *Node
	left  *Node
	right *Node
}

func (temp *temp) AddFront(number int) {
	newNode := &Node{number: number}
	if temp.left == nil {
		temp.left = newNode
		temp.right = newNode
	} else {
		newNode.next = temp.left
		temp.left.prev = newNode
		temp.left = newNode
	}
}

func (temp *temp) AddBack(number int) {
	newNode := &Node{number: number}
	if temp.right == nil {
		temp.left = newNode
		temp.right = newNode
	} else {
		newNode.prev = temp.right
		temp.right.next = newNode
		temp.right = newNode
	}
}

func (temp *temp) PopFront() {
	fmt.Println(temp.left.number, "was popped")
	temp.left.next.prev = nil
	temp.left = temp.left.next
}
func (temp *temp) PopBack() {
	fmt.Println(temp.right.number, "was popped")
	temp.right.prev.next = nil
	temp.right = temp.right.prev
}

func (temp *temp) IsExist(number int) {
	now := temp.left
	flag := false
	for now != nil {
		if now.number == number {
			flag = true
			break
		}
		now = now.next
	}
	if flag {
		fmt.Println(number, "is exist")
	} else {
		fmt.Println(number, "is not exist")
	}
}

func (temp *temp) PrintuyEmae() {
	now := temp.left
	for now != nil {
		fmt.Print(now.number, " ")
		now = now.next
	}
	fmt.Println()
}

func main() {
	temp := temp{}
	temp.AddBack(5)
	temp.AddBack(10)
	temp.PrintuyEmae()
	temp.AddBack(15)
	temp.AddFront(52)
	temp.PrintuyEmae()
	temp.PopBack()
	temp.PrintuyEmae()
	temp.PopFront()
	temp.PrintuyEmae()
}
