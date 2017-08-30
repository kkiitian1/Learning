package main

import "fmt"

func main() {
	elements:= map[string]map[string]string{
		"H": map[string]string {
			"name":"Hydrogen",
			"state":"gas",
		},
		"He": map [string]string {
			"name" : "Helium",
			"state" : "gas",
		},
	}

	fmt.Println(elements)
	fmt.Println(elements["H"])
	fmt.Println(elements["H"]["state"])
	fmt.Println(elements["H"]["name"])
}

