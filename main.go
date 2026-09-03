package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Stack struct {
    items []int
}

func (s *Stack) Push(x int) {
    s.items = append(s.items, x)
}

func (s *Stack) Pop() (int, bool) {
    if len(s.items) == 0 {
        return 0, false
    }

    lastIndex := len(s.items) - 1
    value := s.items[lastIndex]

    s.items = s.items[:lastIndex]

    return value, true
}

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    parts := strings.Fields(sc.Text())
    var s Stack
    for _, p := range parts {
        n, _ := strconv.Atoi(p)
        s.Push(n)
    }
    for {
        x, ok := s.Pop()
        if !ok { break }
        fmt.Println(x)
    }
    // _ = fmt.Print
}