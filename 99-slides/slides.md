# Intro to Go Language

by Henry Strickland

`strick@yak.net`

https://github.com/strickyak/golang-demo/

## Broad Outline

* Why I choose Go
* Install Go and my Demos
* Syntax overview
* Command-line demos (e.g. Hello World)
* Webserver demo
* Threaded Music-Synthesizer demo


## Why Go?

I know and use lots of programming languages:
C, C++, Python, Tcl, Java, Forth, Lisp, APL, Haskell,
Bourne Shell, Basic09.

But recently I often pick Go for hobby projects
and recreational programming.  I've used Go for

* My own 6809 emulator (doing_os9/gomar/emu)
* My own APL interpreter (livy-apl)
* My own Tcl interpreter (chirp-lang)
* My own Python-to-Go translater (rye)
* MIDI Player (trig-synth)
* Ham Radio spectrum decoder (for QRSS)
* Ham Radio tone generation (for QRSS)
* File utilities (e.g. decoding and encoding formats)

## Why Go?

* Go is fun!  (Rob Pike said he wanted it to be "fun to use"
like Python.  It works for me!)

* I still like my Go code, even when I come back to it
years later.

## Why Go?

* Fast: Compiles to native code (like C, C++, Rust).
* Garbage collected (like Python, Java, C#), mostly in a separate thread!
* Designed for threading (like Java, C#). (There's so many
CPU cores in your machines today -- use them!)
* Quite "safe" (not as safe as Rust, but much much easier to use).
* Go's lightweight "goroutines" and "channels" are influencing other languages.
* Modules and methods and interfaces and clear public/private distinctions provide strong encapsulation.
* Interfaces and dynamic dispatch and reflection provide flexability.
* Very little magic in the "white space", so it's easy to understand other people's code (including _your own_ code, years later).
 
## Go is conservative.

* Slow to add more features.
* Very serious about backwards compatibility.
* Stable, well-designed standard libraries.
* Repeatable builds with strong module version dependency graph.
* (or disable it with `GO111MODULE=off`)
* The core language required unanimous approval from the three original designers.

## Go team learned from experience

* Based on decades of experience at Bell Labs (e.g. Plan9) and Google's warehouse-scale computing clusters
* Including specific frustrations with Google3 code in C++ (e.g. web indexing and web search serving) and other Google3 code in Java.

## Go is non-orthodox

* Lacks many checkboxes on the list of features that
so many "othodox" langauges are all converging to.

* You Ain't Gonna Need Them!

No OO inheritance. No overloaded functions or methods.
No overloaded operators.  No classes.  No decorators.
No promotion.  No assert.  No conditional expression.
Not a functional language.

No try/catch (but panic/recover are similar).

No Generics/Templates (until the latest release!)

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

* It might help to install "git".
* Or you can download my raw Go source files from the github web interface.

Web Browser: https://github.com/strickyak/golang-demos

Git: `git clone https://github.com/strickyak/golang-demos.git`

## [syntax] Always Left-to-Right

The thing being defined always is always named *before* its definition.

Brief examples (we'll visit them in more detail):

`const Limit = 999`

`var Marker byte = 255`

`type Bucket *map[int][]*Fruit`

`func Squared(x int) (int, error) { return x*x, nil }`

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

Use `const` for *Compile Time* constants
(like `constexpr` in C++, not C++ `const`).

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

Technically, `type` creates *new types*, not aliases
to existing types.

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

## [syntax] `func` for methods

The receiver variable (like `this` or `self)
is in parentheses before the method name.
Usually it will be a pointer to a struct type.

```
func (a *Apple) String() {
  return fmt.Sprintf("A %s apple with %d worms.",
                     a.Color, a.NumWorms)
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

## Types: Atomic

* bool

* int, uint, int8, byte=uint8, int16, uint16, int32, rune=uint32,
int64, uint64, uintptr

* float32, float64, complex64, complex128

* string

## Types: Composed

* pointers: `*Apple`

* structs: `struct { Name string, Zip uint }`

* slice: `[]string`

* map: `map[string]*Apple`

* channel:  `chan float64`

Those last three (slices, maps, channels) pass by reference.
Everything else in Go passes by value.

## Types: Interfaces

Interfaces are magic references to some underlying non-interface type.

Interfaces act like pointers.

* `type Stringer interface { String() string }`

* Frequently used:  `error`, `io.Reader`, `io.Writer`, `fmt.Stringer`

* `interface {}` matches ANY TYPE IN GO.

There are ways to find out the underlying type
(e.g. type assertions, type switches, reflection,
calling a method and see where it goes).


## Special constants

* true

* false

* nil


## Statements: quick and obvious

* `return foo`

* `return foo, bar`

* `break` and `continue`

## Statements: assignment

```
var a int  // Creates variable a
a = b + c
```

```
a := b + c  // Creates variable a
```

```
a += b + c
a++
a--
```

## Statements: if

```
if x < 2 {
  result = 1
} else {
  result = x * factorial(x-1)
}
```

```
if p := something(); p != nil {
  p.frobnicate()
}
```

## Statements: for

```
for {         // Loops forever.
  step()
}
```

```
line := GetLine()
for line != "" {  // Loops until false.
  line = GetLine()
}
```

```
for i := 1; i <= 10; i++ {
  fmt.Printf("counting %d\n", i)
}
```

```
for key, value := range aSliceOrMap {
  fmt.Printf("Key is %v and value is %v\n", key, value)
}
```

## Statements: `defer` and `go`

```
defer fileDescriptor.Close()

go RunSomething(i, data)
```

## Statements: Slice type

```
var lines []string
for {
  line := GetLine()
  if line == "" {
    break
  }
  lines = append(lines, line)
}

for _, s := range lines {
  fmt.Printf("Got line %q\n", s)
}
```

## Statements: Map type

```
d := make(map[string]int)
for {
  line := GetLine()
  if line == "" {
    break
  }
  x, ok := d[line]
  if ok {
    d[line] += x
  } else {
    d[line] = x
  }
}

for k, v := range d {
  fmt.Printf("Line %q occured %d times\n", k, v)
}
```



## Statements: Channel type

```
pipe := make(chan float64, 1000)  // Buffers 1000 items.

pipe <- 3.14

close(pipe)

x, ok := <-pipe   // ok gets false if pipe is closed.
```

## Statements: Printf & friends

```
fmt.Printf("Hello %s, you are visitor %d\n", name, nth)

fmt.Fprintf(w, "Hello %s, you are visitor %d\n", name, nth)

message := fmt.Sprintf("%d: %s", nth, name)

log.Printf("Funny syntax in %q: %v", filename, error)

log.Panicf("Cannot open %q: %v", filename, error)

log.Fatalf("Cannot create %q: %v", filename, error)
```

