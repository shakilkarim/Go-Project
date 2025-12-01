package main
import "fmt"

var a int = 10;

var b int = 20;

func add(x, y int) int {
	result := x + y
	fmt.Println("Sum is:", result)
	return result
}

func main(){
	p := 2
	q := 3

	first := add(p,a);

	fmt.Println("First Sum is:", first);

	add(b,q);
}