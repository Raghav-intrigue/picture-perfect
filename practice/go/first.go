// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")

	f := fib()
	//Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f())
}
