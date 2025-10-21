package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		capacity: capacity,
		data:     make([]int, 0, capacity),
		size:     0,
	}
}

func (m *MultiSet) insert(value int) {
	if len(m.data) == 0 {
		m.data = append(m.data, value)
		m.size++
		return
	}

	inserted := false
	var new_data []int

	for i := range m.data {
		if !inserted && value <= m.data[i] {
			new_data = append(new_data, value)
			inserted = true
		}
		new_data = append(new_data, m.data[i])
	}

	if !inserted {
		new_data = append(new_data, value)
	}

	m.data = new_data
	m.size++
}

func (m *MultiSet) contains(value int) bool {
	for i := range m.data {
		if m.data[i] == value {
			return true
		}
	}
	return false
}

func (m *MultiSet) erase(value int) {
	var aux []int
	v := value
	for i := range m.data {
		if v == m.data[i] {
			m.size--
			v *= -v
			continue
		}
		aux = append(aux, m.data[i])
	}
	if v > 0 {
		fmt.Println("value not found")
	}
	m.data = aux
}

func (m *MultiSet) count(value int) int {
	cont := 0
	for i := range m.data {
		if value == m.data[i] {
			cont++
		}
	}
	return cont
}

func (m *MultiSet) unique() int {
	cont := 1
	if len(m.data) == 0 {
		return 0
	}
	for i := range len(m.data) - 1 {
		if m.data[i] != m.data[i+1] {
			cont++
		}
	}
	return cont
}

func (m *MultiSet) clear() {
	m.data = nil
	m.size = 0
}

func (m *MultiSet) show() string {
	saida := "["
	for i := range m.data {
		if i != 0 {
			saida += ", " + fmt.Sprint(m.data[i])
		} else {
			saida += fmt.Sprint(m.data[i])
		}
	}
	saida += "]"
	return saida
}

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

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.insert(value)
			}
		case "show":
			fmt.Println(ms.show())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			ms.erase(value)
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.count(value))
		case "unique":
			fmt.Println(ms.unique())
		case "clear":
			ms.clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
