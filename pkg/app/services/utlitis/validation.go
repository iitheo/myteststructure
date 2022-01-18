package utlitis

import (
	"errors"
	"fmt"
	"github.com/iitheo/theohubuc/pkg/app/models/usermodel"
	"github.com/iitheo/theohubuc/pkg/repositories/userrepo"
)

func ValidateUser(userData *usermodel.Users) (myErr error) {
	if (userData.Email == "") || (userData.Username == "") || !(len(userData.Password) >= 6 && len(userData.Password) <= 120) {
		return errors.New("Password must be between 6 and 120. Email and Username are required")
	}
	listOfUsers := userrepo.UserRepo()
	//check for duplicate usernames
	if _, ok := listOfUsers[userData.Email]; ok {
		return errors.New(fmt.Sprintf("duplicate user - %s", userData.Email))
	}

	return nil
}
