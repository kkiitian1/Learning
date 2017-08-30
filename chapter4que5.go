package main

import "fmt"

func main() {
	var c,f float64
	fmt.Println("Enter the temperature :")
	fmt.Scanf("%f" , &f)
	c= (f-32)*5/9
	fmt.Println("Temperature in celsius is :")
	fmt.Println(c)
}
