package main

// go build main.go -> compile
// go run main.go ->  compile and run

// package main (must have main func) -> executable package because "main" is special

// package calcualtor -> reusable package which can be used as adepednecy

// package apple
// go run main.go -> go run: cannot run non-main package
// go build -> nothing compile

import "fmt"

func main() {
	fmt.Println("Hello there!")
}
