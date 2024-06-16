package demo

import (
	"fmt"
)

// 創建一個typw struct -> 類似於OOP的class
type Person struct {
	Name   string
	height int
}

// 也可以串葡萄
type Group struct {
	groupName string
	Person
}

func SetNewPerson() *Person {
	// new物件
	result := new(Person)
	result.Name = "Test"
	result.height = 178
	return result
}

func SetGroupPerson() *Group {
	reult := new(Group)
	reult.groupName = "group test"
	reult.Person = *SetNewPerson()
	return reult
}

func (recv Person) PersonSayHi() {
	fmt.Printf("Hi, my name is %s!\n", recv.Name)
}

func (p *Person) GetHeight() int {
	return p.height
}

func (p *Person) SetHeight(newHeight int) {
	p.height = newHeight
}

type Engine interface {
	Run()
	Stop()
}

type Bus struct {
	Engine
}

func (c *Bus) Working() {
	c.Run()
	c.Stop()
}

func (c *Bus) Run() {
	fmt.Println("Bus is running")
}

func (c *Bus) Stop() {
	fmt.Println("Bus is stop")
}
