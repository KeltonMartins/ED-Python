package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
type Node struct{
	value int
	next *Node
	prev *Node
	root *Node
}

type LList struct{
	root *Node
}
func NewLList() *LList{
	root := &Node{}
	root.next = root
	root.prev = root
	root.root = root
	return &LList{root: root}
}

func (ll *LList) PushBack(value int){
	
	node := &Node{
		value: value,
		next: ll.root,
		prev: ll.root.prev,
		root: ll.root,
	}

	ll.root.prev.next = node
	ll.root.prev = node

}

func (ll *LList) String() string{
	var final []string
	for node := ll.root.next; node != ll.root; node = node.next {
		final = append(final, strconv.Itoa(node.value))
	}
	return "[" + strings.Join(final, ", ") + "]"
}

func (ll *LList) Size() int {
	size := 0
	for node := ll.root.next; node != ll.root; node = node.next {
		size++
	}
	return size
}

func (ll *LList) PushFront(value int) {
	node := &Node{
		value: value,
		next:  ll.root.next,
		prev:  ll.root,
		root:  ll.root,
	}
	ll.root.next.prev = node
	ll.root.next = node
}

func (ll *LList) PopBack() {
	if ll.root.prev == ll.root {
		return
	}
	ll.root.prev.prev.next = ll.root
	ll.root.prev = ll.root.prev.prev
}

func (ll *LList) PopFront() {
	if ll.root.next == ll.root {
		return
	}
	ll.root.next.next.prev = ll.root
	ll.root.next = ll.root.next.next
}
func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
    ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			 fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushBack(num)
			}
		case "push_front":
			 for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushFront(num)
			 }
		case "pop_back":
			 ll.PopBack()
		case "pop_front":
			 ll.PopFront()
		case "clear":
			 ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
