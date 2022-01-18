package userscontroller

import (
	"fmt"
	"github.com/iitheo/theohubuc/pkg/api/libs/httplibr"
	"github.com/iitheo/theohubuc/pkg/app/config/httpresponses"
	"github.com/iitheo/theohubuc/pkg/app/models/usermodel"
	"github.com/iitheo/theohubuc/pkg/app/services/utlitis"
	"github.com/iitheo/theohubuc/pkg/repositories/userrepo"
	"net/http"
)

func CreateUser(res http.ResponseWriter, req *http.Request) {
	c := httplibr.C{Req: req, Res: res}
	var (
		resp    httpresponses.HttpResponse
		newUser usermodel.Users
	)
	err := c.BindJSON(&newUser)
	if err != nil {
		resp.Success = false
		resp.Message = fmt.Sprintf("Error creating user - %s", err.Error())
		resp.Error = fmt.Sprintf("Error creating user - %s", err.Error())
		httplibr.Response400(res, resp)
		return
	}

	err = utlitis.ValidateUser(&newUser)
	if err != nil {
		resp.Success = false
		resp.Message = fmt.Sprintf("Error creating user - %s", err.Error())
		resp.Error = fmt.Sprintf("Error creating user - %s", err.Error())
		httplibr.Response400(res, resp)
		return
	}

	userrepo.InsertUser(newUser)

	resp.Success = true
	resp.Data = newUser
	resp.Message = fmt.Sprintf("user successfully created")
	httplibr.Response201(res, resp)
}

/*
username  //required
email //required, email
password //required 6-120 characters
*/
