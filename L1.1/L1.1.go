package main

import "log"

type Human struct {
	name string
	age  int8
}

func (man *Human) SetName(name string) {
	man.name = name
}

func (man *Human) SetAge(age int8) {
	man.age = age
}

func (man *Human) GetName() string {
	return man.name
}

func (man *Human) GetAge() int8 {
	return man.age
}

// Наследование методов в языке golang реализуется с помощью композиции
type Action struct {
	Human
}

func main() {
	var man Action
	// у структуры Action присутствуют методы структуру Human
	man.SetName("qwe")
	log.Println(man.GetName())
}
