package main

import (
	"flag"
	"fmt"
	"log"
)

func LogAnArg(label string, i int, s string) {
	log.Printf("%s: Arg %d: %q", label, i, s)
	println("***",i,s)
}

// LogOnce logs the args using a C-style for loop.
func LogOnce() {
	n := flag.NArg()
	for i := 0; i < n; i++ {
		a := flag.Arg(i)
		LogAnArg("once", i, a)
	}
}

// LogAgain logs the args using a range-style for loop.
func LogAgain() {
	for i, a := range flag.Args() {
		LogAnArg("again", i, a)
	}
}

func main() {
	println("first")
	fmt.Printf("second\n")
	flag.Parse()
	LogOnce()
	LogAgain()
}
