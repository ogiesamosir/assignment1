package repositories

import (
	"latihan/models"
	"latihan/params"
	"strings"
)

func InsertUser(req *params.CreateUser) models.User {
	var new_user models.User

	new_user.AttNo = req.AttNo
	new_user.Name = req.Name
	new_user.Job = req.Job
	new_user.Reason = req.Reason

	return new_user
}

func SearchUser(req *params.LookUser) models.User {
	var users = req.Users
	var found models.User

	//test isi user
	// fmt.Println("Isi User : ")
	// fmt.Println(users)

	for i, user := range users {
		if strings.EqualFold(user.AttNo, req.AttNo) {
			found = users[i]
			break
		}
	}
	return found
}
