package utlitis

import (
	"errors"
	"fmt"
	"github.com/iitheo/theohubuc/pkg/app/models/usermodel"
	"github.com/iitheo/theohubuc/pkg/repositories/userrepo"
)

func ValidateUser(userData *usermodel.Users) (myErr error) {
	listOfUsers := userrepo.UserRepo()
	//check for duplicate usernames
	if _, ok := listOfUsers[userData.Email]; ok {
		return errors.New(fmt.Sprintf("duplicate user - %s", userData.Email))
	}

	return nil
}
