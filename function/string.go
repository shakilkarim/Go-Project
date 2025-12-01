package main 

import "fmt"

func printSomeString(s string){  
	fmt.Println("This is my name , ",s)
}

func main(){
	name := "Shakeel Ahmed"
	printSomeString("shakil")
	printSomeString(name)
}