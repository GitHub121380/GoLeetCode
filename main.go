package main

import "fmt"

func main() {

	s := "adb"
	fmt.Printf("%T", s[0])
}

type Book struct {
	name  string
	price uint
}

func (this *Book) setName(name string) {
	this.name = name
}

type StudyBook struct {
	Book //写父类名，表示继承父类
	time uint
}

func (this *StudyBook) setName(name string) {
	this.name = name + "ttt"
}
