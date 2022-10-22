package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatalf("Got %d args, requires at least 2 args", flag.NArg())
	}

	fmt.Printf("First arg is %q, second arg is %q\n", flag.Arg(0), flag.Arg(1))
	fmt.Printf("Adding first and second arg gives %q\n", flag.Arg(0)+flag.Arg(1))

	// Join all args, slowly.
	s := ""
	for _, a := range flag.Args() {
		s = s + a
	}
	fmt.Printf("Joined slowly gives %q\n", s)

	// Join all args, quicker.
	var b strings.Builder
	for _, a := range flag.Args() {
		b.WriteString(a)
	}
	fmt.Printf("Joined quicker gives %q\n", b.String())

	// Join quickest.
	joined := strings.Join(flag.Args(), "")
	fmt.Printf("Joined quickest gives %q\n", joined)
	fmt.Printf("... ToUpper gives %q\n", strings.ToUpper(joined))
	fmt.Printf("... ToLower gives %q\n", strings.ToLower(joined))
	fmt.Printf("... Split on `.` gives %#v\n", strings.Split(joined, "."))
}
