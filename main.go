package main
import ("bufio"; "fmt"; "os"; "strconv"; "strings"; "sync")
func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan(); n, _ := strconv.Atoi(sc.Text())
    sc.Scan(); fields := strings.Fields(sc.Text())
    nums := make([]int, n)
    for i, f := range fields { nums[i], _ = strconv.Atoi(f) }
    chunk := (n+3)/4
    var wg sync.WaitGroup
    var total int
    var mu sync.Mutex

    for i := 0; i < 4; i++ {
        start, end := i * chunk, (i + 1) * chunk

        if start > n {
            start = n
        }

        if end > n {
            end = n
        }

        wg.Add(1)
        go func(start, end int) {
            defer wg.Done()
            partial := 0

            for _, v := range nums[start:end] {
                partial += v
            }

            mu.Lock()
            total += partial
            mu.Unlock()

        }(start, end)
    }

    wg.Wait()
    
    fmt.Println(total)
}
