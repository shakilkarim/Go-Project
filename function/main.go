package main

import "fmt"

func add(a int, b int) int{
	return  a + b;
}


func main() {

	a := 10
	b := 20

	sum := add(a,b);


	fmt.Println("Sum:", sum)
}