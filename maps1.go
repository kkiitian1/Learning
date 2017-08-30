package main

import "fmt"

func main() {
	elements:= make (map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"


	fmt.Println(elements)
	fmt.Println(elements["He"])

	if name,ok := elements["Un"] ; ok {
		fmt.Println( name,ok)
	} else {
		fmt.Println("Successful")
	}

}
