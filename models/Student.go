package models

import (
	"fmt"
)

type Student struct {
	Name    string
	Id      int
	ClassId int
}

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I am %s", person, stu.Name)
}
