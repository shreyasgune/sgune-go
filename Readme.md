# Go
Golang, also known as Go, is a statically typed, compiled programming language designed by Google. It was created by Robert Griesemer, Rob Pike, and Ken Thompson in 2007 and released publicly in 2009. Go is particularly popular for its simplicity, performance, and suitability for modern software development needs like concurrency, microservices, and cloud-native applications.

## Features
- Go’s lightweight threads (goroutines) and channels simplify concurrent programming.

```
package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printMessage("Hello, Goroutines!") // Runs concurrently
	printMessage("Main Function")
}
```

- Go avoids unnecessary complexity. Functions and code blocks are concise and easy to read.

- Compiled directly to machine code with minimal overhead, offering high speed. Here’s a small example of Go’s efficient looping and memory management.

```
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	fmt.Println("Sum:", sum)
}
```
- Standard libaries makes life easy.

```
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

```
- Statically Typed (enforces type safety)

```
package main

import "fmt"

func multiply(a int, b int) int {
	return a * b
}

func main() {
	result := multiply(4, 5) // Type-safe operation
	fmt.Println("Product:", result)
}
```

- Cross platform binaries compilation

```
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Running on:", runtime.GOOS, "with", runtime.GOARCH, "architecture.")
}
```
