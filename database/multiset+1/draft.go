package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data:     make([]int, 0, capacity),
		capacity: capacity,
	}
}

func (m *MultiSet) insert(values []int) {
	m.data = append(m.data, values...)
}

func (m *MultiSet) show() {
	fmt.Print("[")
	for i := range m.data {
		if i != 0 {
			fmt.Print(", ")
		}
		fmt.Print(m.data[i])
	}
	fmt.Println("]")
}

func (m *MultiSet) contains(value int) {
	index := sort.Search(len(m.data), func(i int) bool {
		return m.data[i] >= value
	})
	if index < len(m.data) && m.data[index] == value {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func (m *MultiSet) erase(value int) {
	index := sort.Search(len(m.data), func(i int) bool {
		return m.data[i] >= value
	})
	if index < len(m.data) && m.data[index] == value {
		m.data = append(m.data[:index], m.data[index+1:]...)
	} else {
		fmt.Println("value not found")
	}
}

func (m *MultiSet) count(value int) {
	lo := sort.Search(len(m.data), func(i int) bool {
		return m.data[i] >= value
	})
	hi := sort.Search(len(m.data), func(i int) bool {
		return m.data[i] > value
	})
	fmt.Println(hi - lo)
}

func (m *MultiSet) unique() {
	if len(m.data) == 0 {
		fmt.Println(0)
		return
	}
	count := 1
	for i := 1; i < len(m.data); i++ {
		if m.data[i] != m.data[i-1] {
			count++
		}
	}
	fmt.Println(count)
}

func (m *MultiSet) clear() {
	m.data = []int{}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ms *MultiSet

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

		switch args[0] {
		case "init":
			capacity, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(capacity)
		case "insert":
			n, _ := strconv.Atoi(args[1])
			values := make([]int, 0, n)
			for _, s := range args[2:] {
				val, _ := strconv.Atoi(s)
				values = append(values, val)
			}
			ms.insert(values)
		case "show":
			ms.show()
		case "contains":
			val, _ := strconv.Atoi(args[1])
			ms.contains(val)
		case "erase":
			val, _ := strconv.Atoi(args[1])
			ms.erase(val)
		case "count":
			val, _ := strconv.Atoi(args[1])
			ms.count(val)
		case "unique":
			ms.unique()
		case "clear":
			ms.clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
