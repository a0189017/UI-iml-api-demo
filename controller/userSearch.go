package controller

import (
	"impl-api/pear"
	"impl-api/service"
	"net/http"

	"github.com/gorilla/mux"
)

// @Summary userSearch
// @Id 4
// @Tags UserInfo
// @description 搜尋使用者
// @Param Authorization header string true "Bearer JWT"
// @param key path string true "key"
// @version 2.0
// @produce text/json
// @Success 200 string string
// @Router /api/userSearch/{key} [get]
func UserSearch(w http.ResponseWriter, r *http.Request) {
	key := pear.InputCheck(mux.Vars(r)["key"])

	ResultCode, result := service.SearchFullname(key)
	response := ApiResponse{ResultCode: ResultCode, ResultMessage: result}
	switch ResultCode {
	case 0:
		service.ResponseWithJson(w, http.StatusOK, response) //回傳
	case 5:
		service.ResponseWithJson(w, http.StatusBadRequest, response) //回傳
	}
}
