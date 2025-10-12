package main

import "fmt"

func main() {
    // var name string
    // var age int
    // var height float64

    // name = "John"
    // age = 30
    // height = 5.9

    // fmt.Println("Name:", name)
    // fmt.Println("Age:", age)
    // fmt.Println("Height:", height)

    var a int = 10;
    var b int = 20;

    var total int = a + b;

    var newPayment float32 = 100000.4477777777;
    fmt.Println("Payment:", newPayment);
    var bkashPayment float64 = 70000000000000.7777777777777777777777777;

    var Shakil string = "shakil new";

    fmt.Println("Name: ", Shakil);
    fmt.Println("Bkash Payment:", bkashPayment);
    fmt.Println("Total:", total);

    for i:=0; i<5; i++{
        fmt.Println("Hello World", i);
    }
    
    var isTrue bool = false;

    if isTrue {
        fmt.Println("This is true");
    } else {
        fmt.Println("This is false");
    }

    for i:= 5; i>0; i--{
        fmt.Println("Hello Bangladesh", i);
    }
}