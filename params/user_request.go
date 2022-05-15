package params

import "latihan/models"

type CreateUser struct {
	AttNo  string
	Name   string
	Job    string
	Reason string
}

type LookUser struct {
	AttNo  string
	Name   string
	Job    string
	Reason string
	Users  []models.User
}

func NewCreateUser(attNo string, name string, job string, reason string) *CreateUser {
	return &CreateUser{
		AttNo:  attNo,
		Name:   name,
		Job:    job,
		Reason: reason,
	}
}

func FindUser(attNo string, users []models.User) *LookUser {
	return &LookUser{
		AttNo: attNo,
		Users: users,
	}
}
