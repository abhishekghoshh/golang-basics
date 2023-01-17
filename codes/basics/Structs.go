package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) welcome() {
	fmt.Println("Welcome", p.name)
}
func (p *Person) increase(age int) {
	p.age += age
}
func (p Person) showAge() {
	fmt.Println("age of", p.name, "is", p.age)
}
func main() {
	abhishek := Person{"Abhishek", 25}
	kushal := Person{name: "Kushal", age: 25}
	fmt.Println(abhishek, kushal)
	abhishek.welcome()
	abhishek.increase(2)
	abhishek.showAge()
}
