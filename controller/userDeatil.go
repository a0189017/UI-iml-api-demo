package controller

import (
	"fmt"
	"impl-api/service"
	"net/http"

	"github.com/gorilla/context"
)

// @Summary UserDetail
// @Id 3
// @Tags UserInfo
// @description 取得使用者資訊
// @Param Authorization header string true "Bearer JWT"
// @version 2.0
// @produce text/json
// @Success 200 string string
// @Router /api/userDetail [get]
func UserDetail(w http.ResponseWriter, r *http.Request) {
	val, _ := context.GetOk(r, "acct")
	acct := fmt.Sprintf("%v", val)
	ResultCode, result := service.UserDetail(acct)
	response := ApiResponse{ResultCode: ResultCode, ResultMessage: result}
	switch ResultCode {
	case 0:
		service.ResponseWithJson(w, http.StatusOK, response) //回傳
	case 5:
		service.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}
}
