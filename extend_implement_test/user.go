package main

type people interface {
	getName() string
	getAge() int
}

type user struct {
	name string
	age  int
}

func (u *user) getName() string {
	return "需要重写方法才能拿到name"
}

func (u *user) getAge() int {
	return -1
}

type man struct {
	user
	sex string
}

func (m *man) getName() string {
	return m.name
}

func getPeopleName(p people) string {
	return p.getName()
}
