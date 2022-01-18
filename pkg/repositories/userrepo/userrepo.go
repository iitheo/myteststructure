package userrepo

import "github.com/iitheo/theohubuc/pkg/app/models/usermodel"

var mapOfUsers = map[string]usermodel.Users{}

func init() {
	mapOfUsers["theo@yahoo.com"] = usermodel.Users{
		Username: "theo",
		Email:    "theo@yahoo.com",
		Password: "hello",
	}

	mapOfUsers["john@gmail.com"] = usermodel.Users{
		Username: "john",
		Email:    "john@gmail.com",
		Password: "goodbye",
	}

	mapOfUsers["susan@gmail.com"] = usermodel.Users{
		Username: "susan",
		Email:    "susan@gmail.com",
		Password: "myPassword",
	}
}
func InsertUser(newUser usermodel.Users) {
	mapOfUsers[newUser.Email] = newUser
}

func UserRepo() map[string]usermodel.Users {
	return mapOfUsers
}
