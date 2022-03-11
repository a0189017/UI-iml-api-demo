package controller

import (
	"fmt"
	"impl-api/pear"
	service "impl-api/service"
	"net/http"

	"github.com/gorilla/context"
)

// @Summary userUpdate
// @Id 7
// @Tags userUpdate
// @description 使用者資訊更新
// @Param Authorization header string true "Bearer JWT"
// @param pwd formData string true "pwd"
// @param fullname formData string true "fullname"
// @version 2.0
// @produce text/json
// @Success 200 string string
// @Router /api/userUpdate [put]
func UserUpdate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	val, _ := context.GetOk(r, "acct")
	acct := fmt.Sprintf("%v", val)
	fullname, pwd := pear.InputCheck(r.FormValue("fullname")), r.FormValue("pwd")
	var response ApiResponse
	updateInfo := service.UserUpdateInfo{Fullname: fullname, Pwd: pwd}
	if updateInfo.Fullname == "" || updateInfo.Pwd == "" {
		response = ApiResponse{ResultCode: 2, ResultMessage: ErrorMessage["error2"]}
		service.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	} else {
		ResultCode, result := service.UserUpdate(updateInfo, acct)
		response = ApiResponse{ResultCode: ResultCode, ResultMessage: result}
		switch ResultCode {
		case 0:
			service.ResponseWithJson(w, http.StatusOK, response) //回傳
		case 5:
			service.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
		}
	}
}
