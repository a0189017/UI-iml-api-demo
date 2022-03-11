package controller

import (
	"impl-api/pear"
	service "impl-api/service"
	"net/http"
)

// @Summary userSignin
// @Id 1
// @Tags Signin/Signup
// @description 使用者登入
// @param acct formData string true "acct"
// @param pwd formData string true "pwd"
// @version 2.0
// @produce text/json
// @Success 201 string string
// @Router /api/userSignin [post]
func UserSignin(w http.ResponseWriter, r *http.Request) {
	acct, pwd := pear.InputCheck(r.FormValue("acct")), r.FormValue("pwd")
	loginInfo := service.UserLogin{Acct: acct, Pwd: pwd}

	defer r.Body.Close()
	var response ApiResponse

	if len(loginInfo.Acct) == 0 || len(loginInfo.Pwd) == 0 {
		response = ApiResponse{ResultCode: 2, ResultMessage: ErrorMessage["error2"]}
		service.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	} else {
		ResultCode, result := service.Signin(loginInfo)
		response = ApiResponse{ResultCode: ResultCode, ResultMessage: result}
		switch ResultCode {
		case 0:
			service.ResponseWithJson(w, http.StatusCreated, response) //回傳
		case 5:
			service.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
		}
	}
}
