package main

import (
    "fmt"
    "math/big"
    "time"
)

func main() {
    for {
        f0 := big.NewInt(0)
        f1 := big.NewInt(1)
        var n int
        
        fmt.Println("Enter an integer, or 9999 to quit:")
        
        _, err := fmt.Scanf("%d", &n)
        if err != nil {
            fmt.Println(err)
        }
        if n == 9999 {
            break
        }
        // repeat fib(n) 100 times
        start := time.Now()
        i := 100
        for i > 0 {   // while i > 0
            for n > 0 {
                f0, f1, n = f1, f0.Add(f0, f1), n-1
            }
            i--
        }
        // and divide elapsed time by 100
        duration := time.Since(start) / 100.0
        fmt.Println(duration)
    }
}