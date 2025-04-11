
# Fundamentals

### **Part 1: The Basics of Go**
Go is a statically typed, compiled language designed for simplicity, performance, and concurrency. It was created by Google in 2009 by Robert Griesemer, Rob Pike, and Ken Thompson.

#### **1.1 Installation**
- Download Go from [golang.org](https://golang.org/dl/).
- Install it (e.g., on macOS/Linux: extract the tarball to `/usr/local/go`; on Windows: use the MSI installer).
- Set up your `GOPATH` (a workspace directory, e.g., `~/go`) and add Go’s `bin` directory to your `PATH`.
- Verify with: `go version` in your terminal.

#### **1.2 First Program**
Create a file called `hello.go`:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
- Run it: `go run hello.go`
- Build it: `go build hello.go` (creates an executable).

**Key Concepts:**
- `package main`: Every Go program belongs to a package. `main` is the entry point.
- `func main()`: The starting point of execution.
- `import`: Brings in libraries (like `fmt` for formatting and printing).

#### **1.3 Variables and Types**
Go is strongly typed but concise:
```go
package main

import "fmt"

func main() {
    var x int = 10        // Explicit declaration
    y := 20               // Short declaration (type inferred)
    name := "Grok"        // String
    isCool := true        // Boolean
    pi := 3.14            // Float

    fmt.Println(x, y, name, isCool, pi)
}
```
- Basic types: `int`, `float64`, `string`, `bool`.
- Zero values: Variables declared without initialization get a default (e.g., `0` for `int`, `""` for `string`).

#### **1.4 Control Structures**
Go keeps it simple:
```go
package main

import "fmt"

func main() {
    x := 10
    if x > 5 {
        fmt.Println("x is big")
    } else {
        fmt.Println("x is small")
    }

    for i := 0; i < 3; i++ {
        fmt.Println(i)
    }
}
```
- No parentheses needed in `if` or `for`.
- Only `for` is used for loops (no `while`).

#### **1.5 Functions**
Functions are straightforward:
```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(3, 4)
    fmt.Println(result) // 7
}
```
- Multiple return values are supported:
```go
func swap(x, y string) (string, string) {
    return y, x
}
```

---

### **Part 2: Core Features**
#### **2.1 Arrays, Slices, and Maps**
- **Arrays**: Fixed size.
```go
var arr [3]int = [3]int{1, 2, 3}
```
- **Slices**: Dynamic, built on arrays.
```go
slice := []int{1, 2, 3}
slice = append(slice, 4) // Grows dynamically
```
- **Maps**: Key-value pairs.
```go
m := map[string]int{"age": 30, "score": 100}
fmt.Println(m["age"]) // 30
```

#### **2.2 Structs**
Go uses structs instead of classes:
```go
type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Grok", Age: 1}
    fmt.Println(p.Name) // Grok
}
```

#### **2.3 Pointers**
Go has pointers but no pointer arithmetic:
```go
x := 10
p := &x       // p is a pointer to x
*p = 20       // Changes x to 20
fmt.Println(x) // 20
```

#### **2.4 Error Handling**
No exceptions—errors are values:
```go
import "errors"

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

---

### **Part 3: Concurrency**
Go’s concurrency is a killer feature, built around **goroutines** and **channels**.

#### **3.1 Goroutines**
Lightweight threads:
```go
func sayHello() {
    fmt.Println("Hello from goroutine")
}

func main() {
    go sayHello() // Runs concurrently
    fmt.Println("Main function")
    time.Sleep(time.Second) // Wait for goroutine
}
```

#### **3.2 Channels**
Communication between goroutines:
```go
func main() {
    ch := make(chan string)
    go func() {
        ch <- "Hello from goroutine"
    }()
    msg := <-ch // Receive from channel
    fmt.Println(msg)
}
```

---

### **Part 4: Internals of Rosé**
Now, let’s peek under the hood of Go.

#### **4.1 Compilation**
- Go is compiled to machine code (no VM like Java).
- The `go build` command uses the Go compiler (`gc`) to produce a single, statically linked binary.
- Fast compilation due to a simple dependency model (no header files).

#### **4.2 Memory Management**
- Go has a garbage collector (GC):
  - Mark-and-sweep algorithm.
  - Low-latency, concurrent GC (runs alongside your program).
- Memory is managed automatically—no manual `malloc`/`free`.

#### **4.3 Goroutines**
- Not OS threads—managed by the Go runtime.
- Thousands or millions can run due to low overhead (~2 KB per goroutine vs. ~1 MB per OS thread).
- Go’s scheduler (part of the runtime) multiplexes goroutines onto OS threads.

#### **4.4 Runtime**
- Go embeds a small runtime in every binary:
  - Handles scheduling, GC, stack management.
  - Written in Go and assembly.

#### **4.5 Why So Fast?**
- Simple syntax → faster parsing and compilation.
- Static linking → no dynamic library overhead.
- Efficient concurrency model.

---

### **Part 5: Practice**
Try this:
1. Write a program that spawns 10 goroutines, each printing its number.
2. Use a channel to collect results from them.
3. Handle potential errors.

Here’s a solution to study:
```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    ch := make(chan int)
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            ch <- n
        }(i)
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    for num := range ch {
        fmt.Println("Goroutine:", num)
    }
}
```

---

### **Next Steps**
- Explore the standard library (e.g., `net/http` for web servers).
- Learn about interfaces (Go’s way of polymorphism).
- Check out modules (`go mod`) for dependency management.
- Read [Effective Go](https://golang.org/doc/effective_go) and the [spec](https://golang.org/ref/spec).