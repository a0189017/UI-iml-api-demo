package controller

import (
	"impl-api/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary userList
// @Id 2
// @Tags UserInfo
// @description 取得使用者列表
// @Param Authorization header string true "Bearer JWT"
// @param page path int true "page"
// @version 2.0
// @produce text/json
// @Success 200 string string
// @Router /api/userList/{page} [get]
func UserList(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["page"]
	page, _ := strconv.Atoi(param)
	ResultCode, result := service.UserList(page)
	response := ApiResponse{ResultCode: ResultCode, ResultMessage: result}
	switch ResultCode {
	case 0:
		service.ResponseWithJson(w, http.StatusOK, response) //回傳
	case 5:
		service.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}

}
