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

	// Make a slice of bool, pre-allocating n slots.
	// They are automatically initialized to the "zero value" for bool, which is false.
	sieve := make([]bool, n)
	// For integers i starting at 2, mark all the multiples of i, starting at 2*i, as composite.
	for i := 2; i < n; i++ {
		for j := 2 * i; j < n; j += i {
			sieve[j] = true // j is composite.
		}
	}

	// Create an empty slice of string to hold our reports.
	var reports []string
	// For integers i starting at 2, report that i is prime, if sieve[i] is false.
	for i := 2; i < n; i++ {
		if !sieve[i] {
			// Append a string to the reports.
			reports = append(reports, fmt.Sprintf("%d is prime.", i))
		}
	}
	// Iterate over reports, printing each one.
	for _, report := range reports {
		fmt.Printf("%s\n", report)
	}
}
