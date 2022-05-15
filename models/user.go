package models

import (
	"fmt"
)

type User struct {
	AttNo  string
	Name   string
	Job    string
	Reason string
}

func (u *User) Print() {
	fmt.Println("Attendance Number : ", u.AttNo)
	fmt.Println("Name : ", u.Name)
	fmt.Println("Job : ", u.Job)
	fmt.Println("Reason of Joining Golang Class : ", u.Reason)
}
