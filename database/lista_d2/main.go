package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
type Node[T comparable] struct{
	value T
	next *Node[T]
	prev *Node[T]
	root *Node[T]
}

type List[T comparable] struct {
	root *Node[T]
	size int
}

func NewList[T comparable]() *List[T] {
    root := &Node[T]{}
    root.next = root
    root.prev = root
    root.root = root
    return &List[T]{
        root: root,
        size: 0,
    }
}

func (n *Node[T]) Next() * Node[T]{
	if n.next == n.root{
		return nil
	}
	return n.next
}

func (n *Node[T]) Prev() *Node[T]{
	if n.prev == n.root{
		return nil
	}
	return n.prev
}

func (l *List[T]) Front() *Node[T]{
	if l.root.next == l.root{
		return nil
	}
	return l.root.next
}

func (l *List[T]) Back() *Node[T]{
	if l.root.prev == l.root{
		return nil
	}
	return l.root.prev
}

func (l *List[T]) Size() int{
	return l.size
}

func (l *List[T]) PushBack(value T) {
    node := &Node[T]{
        value: value,
        next:  l.root,
        prev:  l.root.prev,
        root:  l.root,
    }
    l.root.prev.next = node
    l.root.prev = node
    l.size++
}

func (l *List[T]) PushFront(value T) {
    node := &Node[T]{
        value: value,
        next:  l.root.next,
        prev:  l.root,
        root:  l.root,
    }
    l.root.next.prev = node
    l.root.next = node
    l.size++
}

func (l *List[T]) Clear(){
	for node := l.Front(); node != nil;{
		next := node.Next()
		node.next = nil
		node.prev = nil
		node.root = nil
		node = next
	}
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func (l *List[T]) Search(value T) *Node[T] {
	for node := l.Front(); node != nil; node = node.Next() {
		if node.value == value {
			return node
		}
	}
	return nil
}

func (l *List[T]) Insert(node *Node[T], value T){
	novo := &Node[T]{
		value: value,
		next: node,
		prev: node.prev,
		root: node.root,
	}
	node.prev.next = novo
	node.prev = novo
	l.size++
}

func (l *List[T]) Remove(node *Node[T]) *Node[T]{
	no := node
	node.next.prev = node.prev
	node.prev.next = node.next
	node.next = nil
	node.prev = nil
	node.root = nil
	return no
}

func (l *List[T]) String() string{
	final := "["
	for node := l.Front(); node != nil; node = node.Next(){
		final += fmt.Sprint(node.value)
		if node.Next() != nil{
			final += ", "
		}
	}
	final += "]"
	return final

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewList[int]()

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
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}