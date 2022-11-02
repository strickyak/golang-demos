# Intro to Go Language

by Henry Strickland

`strick@yak.net`

https://github.com/strickyak/golang-demos/

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

[GitHub Git Guides Install Git on Windows] -> "Git for Windows" ->
I chose default or recommended options for everything except
I chose to use "main" instead of "master".  
I use the included "Git Bash" (mingw64 terminal) for my shell.

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
func (a *Apple) String() string {
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

Interfaces act like pointers, but you don't need the `*`.

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

Empty slices are equal to nil.

Empty strings are NOT equal to nil.


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

fmt.Printf("There are %d lines\n", len(lines))
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

fmt.Printf("There are %d items\n", len(d))
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

## Demo 01: hello.go

```
$ cat 01-hello/hello.go
     1	package main
     2
     3	import "fmt"
     4
     5	func main() {
     6		fmt.Println("Hello World!")
     7	}
$ go run 01-hello/hello.go
Hello World!
$
```

## Demo 02: func.go

```
$ go run 02-func/func.go one two three
2022/10/27 15:59:41 You provided 3 args.
2022/10/27 15:59:41 once: Arg 0: "one"
2022/10/27 15:59:41 once: Arg 1: "two"
2022/10/27 15:59:41 once: Arg 2: "three"
2022/10/27 15:59:41 again: Arg 0: "one"
2022/10/27 15:59:41 again: Arg 1: "two"
2022/10/27 15:59:41 again: Arg 2: "three"
$
```

## Demo 03: string.go

```
$ go run 03-string/string.go Yak.Net taz.txt
First arg is "Yak.Net", second arg is "taz.txt"
Adding first and second arg gives "Yak.Nettaz.txt"
Joined slowly gives "Yak.Nettaz.txt"
Joined quicker gives "Yak.Nettaz.txt"
Joined quickest gives "Yak.Nettaz.txt"
... ToUpper gives "YAK.NETTAZ.TXT"
... ToLower gives "yak.nettaz.txt"
... Split on `.` gives []string{"Yak", "Nettaz", "txt"}
$
```

## Demo 11: slice.go

```
$ go run 11-slice/slice.go
2 is prime.
3 is prime.
5 is prime.
7 is prime.
11 is prime.
13 is prime.
17 is prime.
19 is prime.
23 is prime.
29 is prime.
31 is prime.
37 is prime.
41 is prime.
43 is prime.
47 is prime.
53 is prime.
59 is prime.
61 is prime.
67 is prime.
71 is prime.
73 is prime.
79 is prime.
83 is prime.
89 is prime.
97 is prime.
$
```

## Demo 21: webserver.go

```
    11	func HelloWeb(w http.ResponseWriter, r *http.Request) {
    12		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    13	}
    14
    15	func main() {
    16		flag.Parse()
    17
    18		http.HandleFunc("/", HelloWeb)
    19
    20		log.Print("Serving.  Go to `http://localhost:8080/foo/bar` in your web browser.")
    21		log.Fatal(http.ListenAndServe("localhost:8080", nil))
    22	}
$ go run 21-webserver/webserver.go
2022/10/27 16:05:42 Serving.  Go to `http://localhost:8080/foo/bar` in your web browser.
^C
signal: interrupt
```

## Demo 22: fileserver.go

```
     9	var PubDir = flag.String("pubdir", ".", "Directory to publish")
    10
    11	func main() {
    12		flag.Parse()
    13
    14		http.Handle("/", http.FileServer(http.Dir(*PubDir)))
    15
    16		log.Print("Serving.  Go to `http://localhost:8080/` in your web browser.")
    17		log.Fatal(http.ListenAndServe("localhost:8080", nil))
    18	}
$ go run 22-fileserver/fileserver.go
2022/10/27 16:06:58 Serving.  Go to `http://localhost:8080/` in your web browser.
^C
signal: interrupt
```

## Demo 31: synth.go

```
$ go run 31-synth/synth.go c5,d5,e5,f5,g5
$ file output.wav
output.wav: RIFF (little-endian) data, WAVE audio, Microsoft PCM, 16 bit, mono 8000 Hz
$
```

VLC plays the file on Windows or Linux.

mplayer will play it on Linux.

If you omit the WAV header (the -n option),
on Linux you can use `aplay -f S16_LE output.wav`

## Asking go programs for --help

If a Go program uses `flag.Parse()`, which I recommend,
you can ask it for `--help`:

```
$ go run 31-synth/synth.go --help
Usage of /tmp/go-build1639908412/b001/exe/synth:
  -n	omit WAV header; just raw S16_LE
  -o string
    	wav file to write (default "output.wav")
  -r int
    	sample rate (in samples per second) (default 8000)
  -t float
    	transpose this many steps up or down
$
```

## END
