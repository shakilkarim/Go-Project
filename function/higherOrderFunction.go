package main 

import "fmt"

// Higher Order Function 

/* 
	Following 3 type role
	1.Parameter hishebe function ney
	2.Return hishebe function dey
	3.Both ney and dey
*/
func processOperation (a int , b int , op func  (p int , q int) int) {
	op(a,b)
	
}

func add (a int , b int)  {
	result := a + b

	fmt.Println("Addition is : ", result)
	return result;
}

// var {
// 	a = 10
// 	b = 20
// }

func main (){
	processOperation (10, 20, add)
}