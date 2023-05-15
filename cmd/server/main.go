package main

import "fmt"

// Run - is goin to be responsible for the instantiation
// and startup of go application
func Run() error {
	fmt.Println("Starting up application")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}