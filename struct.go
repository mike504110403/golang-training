package main

// 創建一個typw struct -> 類似於OOP的class
type Person struct {
	name   string
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
	result.name = "Test"
	result.height = 178
	return result
}

func SetGroupPerson() *Group {
	reult := new(Group)
	reult.groupName = "group test"
	reult.Person = *SetNewPerson()
	return reult
}
