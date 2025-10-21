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

type Vector struct{
	size int
	capacity int
	data []int
}

func NewVector(capacity int) *Vector{
	return &Vector{
		size: 0,
		capacity: capacity,
		data: make([]int, capacity),
	}
}

func (v *Vector) Status() string{
	return "size:" + fmt.Sprint(v.size) + " capacity:" + fmt.Sprint(v.capacity)
}

func (v Vector) At(i int) (int, error) {
    if i < 0 || i >= len(v.data) {
        return 0, fmt.Errorf("index out of range")
    }
    return v.data[i], nil
}

func (v *Vector) Set(index, value int) error {
	if index < 0 || index >= len(v.data) {
		return fmt.Errorf("index out of range")
	}
	v.data[index] = value
	return nil
}

func (v *Vector) PopBack() error{
	if v.size == 0{
		return fmt.Errorf("vector is empty")
	}
	v.data = v.data[:v.size-1]
	v.size--
	return nil
}

func (v *Vector) Insert(index, value int) error{
	if index < 0 || index > v.size-1{
		return fmt.Errorf("index out of range")
	}
	if v.size == v.capacity{
		v.Reserve(max(1, 2 * v.capacity))
	}
	novo := make([]int, v.size+1)
	for i := range v.size{
		if i<index{
			novo[i] = v.data[i]
		}
		if i>=index{
			novo[i+1] = v.data[i]
		}
	}
	novo[index] = value
	v.data = novo
	v.size++
	return nil
}

func (v *Vector) Erase(index int) error{
	if index < 0 || index > v.size-1{
		return fmt.Errorf("index out of range")
	}
	v.data = slices.Delete(v.data, index, index+1)
	v.size--
	return nil
}

func (v *Vector) Reserve(capacity int){
	if capacity < v.size {
		return
	}
	novo := make([]int, capacity)
	for i := range v.size{
		novo[i] = v.data[i]
	}
	v.data = novo
	v.capacity = capacity
}

func (v *Vector) String() string{
	final := "["
	for i := range v.size{
		if i != 0{
			final += ", "
		}
		valor := v.data[i]
		final += fmt.Sprint(valor)
	}
	final += "]"
	return final
}

func (v *Vector) PushBack(value int){
	if v.size == v.capacity{
		v.Reserve(max(1, 2 * v.capacity))
	}
	v.data[v.size] = value
	v.size += 1
}

func (v *Vector) Clear() {
	v.data = []int{}
	v.size = 0
}

func (v *Vector) Contains(value int) bool{
	ok := slices.Contains(v.data, value)
	return ok
}

func (v *Vector) IndexOf(value int) int{
	ok := v.Contains(value)
	if !ok{
		return -1
	}
	return slices.Index(v.data, value)
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
			 	value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			 err := v.PopBack()
			 if err != nil {
			 	fmt.Println(err)
			 }
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
			fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
			fmt.Println("true")
			} else {
			fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			 index, _ := strconv.Atoi(parts[1])
			 value, err := v.At(index)
			 if err != nil {
			 	fmt.Println(err)
			 } else {
			 	fmt.Println(value)
			 }
		case "set":
			 index, _ := strconv.Atoi(parts[1])
			 value, _ := strconv.Atoi(parts[2])
			 err := v.Set(index, value)
			 if err != nil {
			 	fmt.Println(err)
			 }
			 
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
