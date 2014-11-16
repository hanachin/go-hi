package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)
import hi "fmt"

func main() {
	fmt.Printf("Hi\n")
	hi.Println("The time is", time.Now())
	hi.Println("My favorite number is", rand.Intn(10))

	// same
	hi.Printf("Now you have %g problems.\n", math.Nextafter(2, 3))
	hi.Printf("Now you have %g problems.\n", math.Nextafter(2, 100))

	fmt.Println(a(2, 5))
}

func a(x, y int) int {
	return x + y
}
