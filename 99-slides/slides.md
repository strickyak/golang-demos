# Intro to Go Language

by Henry Strickland

`strick@yak.net`

https://github.com/strickyak/golang-demo/

## Install Go

Download from https://go.dev/dl/
for Windows, MacOS (ARM64 or x86-64), or Linux (amd64).

Debian/Ubuntu/Pi: `sudo apt install golang`

Then try this:

```
$ go version
go version go1.17.5 linux/amd64
```

## Fetch today's Demo programs

Web Browser: https://github.com/strickyak/golang-demos

Linux: `git clone https://github.com/strickyak/golang-demos.git`

## [syntax] Global Sections

The first word in each global section is one of these six keywords:

* `package`
* `import`
* `const`
* `type`
* `var`
* `func`

followed immediately by the word being defined,
followed by the details of the definition.

## [syntax] `package`

Today we will only use

```
package main
```

This must be the first section in your program
(besides comments).

## [syntax] `import`

```
import (
  "flag"
  "fmt"
  "net/http"
  "math"
  apl "github.com/strickyak/livy-apl"
)
```

The `import` section is the second section in your program.
It comes after the `package` section.

## [syntax] `const`

```
const MaxUnsignedInt64 = 0xFFFFFFFFFFFFFFFF
const Greeting = "Hello World!"
const PageTemplate = `
  <h1>{{.Title}}</h1>
  <div>{{.Content}}</div> `

const (
  E   = 2.718281828459
  Tau  = 2.0 * math.Pi
)
```

## [syntax] `type`

```
type Time uint64

type Apple struct {
  Color    string
  NumWorms int
  Birthday Time
}

type AppleEater func (*Apple) error
```

## [syntax] `func`

```
func Twice(x int) int {
  return x + x
}
func Double(x string) string {
  return x + x
}
func Hypotenuse(x, y float64) float64 {
  return math.Sqrt(x*x + y*y)
}
```

## [syntax] Comments

```
/* Just like in C */

// Or just like
// in C++, Java, C#, Javascript, etc.
```

## `go doc`

```
go doc net/http

go doc net/http.Header

go doc net/http.Header.Get
```

The problem is usually knowing where to start, like `net/http`.
Search the internet to find out.

## Running a Go program

```
go run hello.go -- args...
```

Go compiles so quickly, you can usually compile every time
you run.  The `go run` command does this for you.
You can separate your arguments from the `go run` arguments
with `--` if you need to.


