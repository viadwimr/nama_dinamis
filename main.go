package main

import (
	"fmt"
	"os"
)

var name = os.Getenv("nama")

func main()  {
	fmt.Println("Hello", name)
}