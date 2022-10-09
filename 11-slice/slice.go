// Sieve of Eratosthenes
package main

import (
    "flag"
    "fmt"
)

var nFlag = flag.Int("n", 100, "biggest number to consider")

func main() {
    flag.Parse()
    n := *nFlag

    sieve := make([]bool, n)
    for i := 2; i < n; i++ {
        for j := 2*i; j < n; j += i {
            sieve[j] = true
        }
    }

    for i := 2; i < n; i++ {
        if !sieve[i] {
            fmt.Printf(" %d", i)
        }
    }
    fmt.Printf("\n")
}
