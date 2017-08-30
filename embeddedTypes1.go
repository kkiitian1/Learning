package main

import "fmt"

type person struct {
	name string
}

func (p *person) talk () {
	fmt.Println("Hello my name is", p.name)
}

type android struct {
	p person
	model string
}

func main() {
	a:= person{"Kiran"}
	b:= new(person)
	b.name = "Sheetal"
    fmt.Println(a.name, b.name)
	a.talk()
	b.talk()

	t:= new(android)
	t.p = *b
	t.model = "Young"

	t.p.talk()
	fmt.Println(t.model, t.p.name)

	//t.talk()

	c:= person {"Nirbheek"}
	k:= android{c, "Chutiya"}
	k.p.talk()
}
