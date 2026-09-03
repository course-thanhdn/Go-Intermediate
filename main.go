package main
import ("bufio"; "fmt"; "os"; "strconv")

func safeDivide(a, b int) (q int, err error) {
    if b == 0 {
        return 0, fmt.Errorf("divide by zero")
    }
    
    return a / b, nil
}

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan(); a, _ := strconv.Atoi(sc.Text())
    sc.Scan(); b, _ := strconv.Atoi(sc.Text())
    q, err := safeDivide(a, b)
    if err != nil {
        fmt.Printf("error: %s\n", err)
    } else {
        fmt.Printf("result: %d\n", q)
    }
}
