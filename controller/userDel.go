package controller

import (
	"fmt"
	"impl-api/service"
	"net/http"

	"github.com/gorilla/context"
)

// @Summary userDel
// @Id 6
// @Tags userDel
// @description 使用者刪除
// @Param Authorization header string true "Bearer JWT"
// @version 2.0
// @produce text/json
// @Success 200 string string
// @Router /api/userDel [delete]
func UserDel(w http.ResponseWriter, r *http.Request) {
	val, _ := context.GetOk(r, "acct")
	acct := fmt.Sprintf("%v", val)
	ResultCode, result := service.UserDel(acct)
	response := ApiResponse{ResultCode: ResultCode, ResultMessage: result}
	switch ResultCode {
	case 0:
		service.ResponseWithJson(w, http.StatusOK, response) //回傳
	case 5:
		service.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}
}
