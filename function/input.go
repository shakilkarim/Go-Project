package main 

import "fmt"

func welcomeMessage(){
	fmt.Println("Welcome to Go Programming!")
}

func printName() string {
	fmt.Println("Enter your name: ")
	var name string 
	fmt.Scanln(&name)
	return name
}

func getTwoNumbers() (int,int){
	var num1 int
	var num2 int

	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)
	
	fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)

	return num1, num2

}

func display(name string, sum int){
	//Display name 
	fmt.Println("Hello,", name)

	//Display sum
	fmt.Println("The sum of the two numbers is:", sum)
}

func goodbyeMessage(){
	fmt.Println("Thank you for using the program. Goodbye!")
}

func arrayExampleAndLoop(arr [5]int) int {
	fmt.Println("Array elements are:")

	sum := 0
	for i := 0; i< len(arr); i++{
		sum += arr[i]
	}

	range based loop 
	for  value := range arr {
		fmt.Printf("Element at index %d is %d\n", value)
	}

	//Any other loop exit in Go
	for _, value := range arr {
		fmt.Printf("Element is %d\n", value)
		sum += value
	}

	return sum
}

func main(){
	welcomeMessage()	

	userName := printName()


	//Two number received from user
	n1, n2 := getTwoNumbers() 


	sum := n1 + n2

	display( userName , sum)


	goodbyeMessage()

	fmt.Println("Program has ended , Thank you! , I love Go programming")

	// Array example and loop another function
	arr := [5]int{10,20,30,40,50}

	sumNumbers:= arrayExampleAndLoop(arr)

	fmt.Println("Sum of array elements is:", sumNumbers)
}