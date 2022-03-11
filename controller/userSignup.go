package controller

import (
	"impl-api/pear"
	service "impl-api/service"
	"net/http"
)

// @Summary userSignup
// @Id 5
// @Tags Signin/Signup
// @description 使用者註冊
// @param acct formData string true "acct"
// @param pwd formData string true "pwd"
// @param fullname formData string true "fullname"
// @version 2.0
// @produce text/json
// @Success 201 string string
// @Router /api/userSignup [post]
func UserSignup(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	acct, pwd, fullname := pear.InputCheck(r.FormValue("acct")), r.FormValue("pwd"), pear.InputCheck(r.FormValue("fullname"))
	signupInfo := service.UserSignup{Acct: acct, Pwd: pwd, Fullname: fullname}
	var response ApiResponse

	if len(signupInfo.Acct) == 0 || len(signupInfo.Pwd) == 0 || len(signupInfo.Fullname) == 0 {
		response = ApiResponse{ResultCode: 2, ResultMessage: ErrorMessage["error2"]}
		service.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	} else {
		ResultCode, result := service.Signup(signupInfo)
		response = ApiResponse{ResultCode: ResultCode, ResultMessage: result}
		switch ResultCode {
		case 0:
			service.ResponseWithJson(w, http.StatusCreated, response) //回傳
		case 5:
			service.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
		}
	}
}
