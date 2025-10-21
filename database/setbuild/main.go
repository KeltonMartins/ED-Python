package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		size:     0,
		capacity: capacity,
		data:     make([]int, capacity),
	}
}
func (v *Set) Contains(value int){
	ok := slices.Contains(v.data, value)
	if ok{
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
}

func (v *Set) Erase(value int) {
	index := -1
	for i := range v.size {
		if v.data[i] == value {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("value not found")
		return
	}
	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}
	v.size--
}

func (v *Set) Insert(value int) {
	if v.size >= len(v.data) {
		v.data = append(v.data, 0)
	}
	for i := range v.size {
		if v.data[i] == value {
			return
		}
	}
	pos := 0
	for pos < v.size && v.data[pos] < value {
		pos++
	}
	for i := v.size; i > pos; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[pos] = value
	v.size++
}

func (v *Set) String() string {
	final := "["
	for i := range v.size {
		if i != 0 {
			final += ", "
		}
		valor := v.data[i]
		final += fmt.Sprint(valor)
	}
	final += "]"
	return final
}

func (v *Set) Clear() {
	v.size = 0
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
			value, _ := strconv.Atoi(part)
			v.Insert(value)
			}
		case "show":
			fmt.Println(v)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			v.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			v.Contains(value)
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
