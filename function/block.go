package main

import "fmt"

var (
	a int = 10;
	b int = 20;
)

func main (){

	p := 18

	if p >=18 {
		sum := 30;

		fmt.Println("I have National ID Card", p);

		result := a + sum;

		fmt.Println("Result is:", result);
		
		if p >3 {
			gTotal :=  sum + b;

			fmt.Println("Global total is:", gTotal);

			fmt.Println("You are eligible for voting");
		}
	}


}
