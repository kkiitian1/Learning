package main

import (
	"fmt"
	"strings"
)

func main() {
	var t,r int

	dictionary:=map[int]string{
		1:"apple",
		2:"ball",
		3:"cat",
		4:"dog",
		5:"elephant",
		6:"ant",
		7:"aflame",
		8:"eve",
		9:"ale",
		10:"absorb",
		11:"abide",
		12:"beer",
		13:"bed",
	}

	fmt.Println("Enter first character: ")
	var input string
	fmt.Scanln(&input)
	for i,value:=range dictionary{
		flag:=strings.HasPrefix(dictionary[i],input)
		if flag==true{
			fmt.Println(i,value)
			r++

		}

	}
	if r==0 {
		fmt.Println("Words starting with",input, "are not present in dictionary")
	} else {

	fmt.Println("Enter index of the word from the list")
	fmt.Scanln(&t)
	fmt.Println(dictionary[t])
	}

}
