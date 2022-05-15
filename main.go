package main

import (
	"latihan/models"
	"latihan/params"
	"latihan/repositories"
	"os"
)

func main() {
	// receiving string argument
	args := os.Args
	args1 := args[1]

	var users []models.User

	// creating and inserting user
	req1 := params.NewCreateUser("11319006", "Julius", "Developer", "Probabtion")
	req2 := params.NewCreateUser("11319007", "Martogi", "Designer", "Same like Hamonangan")
	req3 := params.NewCreateUser("11319008", "Hamonangan", "Researcher", "Just like Julius do")
	req4 := params.NewCreateUser("11319009", "Samosir", "Lecturer", "Martogi is my role model")

	users = append(users, repositories.InsertUser(req1))
	users = append(users, repositories.InsertUser(req2))
	users = append(users, repositories.InsertUser(req3))
	users = append(users, repositories.InsertUser(req4))

	// printing list of users
	// for _, user := range users {
	// 	user.Print()
	// 	println()
	// }

	// looking for user from the list by argument added(id)
	req5 := params.FindUser(args1, users)

	find := repositories.SearchUser(req5)

	// printing the detail of users found
	find.Print()

}
