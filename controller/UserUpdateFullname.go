package controller

import (
	"fmt"
	"impl-api/pear"
	"impl-api/service"
	"net/http"

	"github.com/gorilla/context"
)

// @Summary userUpdateFullname
// @Id 8
// @Tags userUpdate
// @description 使用者fullname更新
// @Param Authorization header string true "Bearer JWT"
// @param fullname formData string true "fullname"
// @version 2.0
// @produce text/json
// @Success 200 string string
// @Router /api/userUpdateFullname [patch]
func UserUpdateFullname(w http.ResponseWriter, r *http.Request) {
	val, _ := context.GetOk(r, "acct")
	acct := fmt.Sprintf("%v", val)
	fullname := pear.InputCheck(r.FormValue("fullname"))
	var response ApiResponse
	if fullname == "" {
		response = ApiResponse{ResultCode: 2, ResultMessage: ErrorMessage["error2"]}
		service.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	} else {
		ResultCode, result := service.UserUpdateFullname(fullname, acct)
		response = ApiResponse{ResultCode: ResultCode, ResultMessage: result}
		switch ResultCode {
		case 0:
			service.ResponseWithJson(w, http.StatusOK, response) //回傳
		case 5:
			service.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
		}
	}
}
