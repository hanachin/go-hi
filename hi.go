package main

import (
	"fmt"
	"math/rand"
	"time"
)
import hi "fmt"

func main() {
	fmt.Printf("Hi\n")
	hi.Println("The time is", time.Now())
	hi.Println("My favorite number is", rand.Intn(10))
}
