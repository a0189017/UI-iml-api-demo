package service

import (
	pear "impl-api/pear"
)

type UserSignup struct {
	Acct     string
	Pwd      string
	Fullname string
}

func (v UserSignup) TableName() string {
	return "users"
}
func Signup(userInfo UserSignup) (ResultCode int, result interface{}) {
	db := DbConnect("ui")
	pwd := pear.PasswordMd5(userInfo.Pwd)
	userInfo.Pwd = pwd
	dbResult := db.Select("acct", "pwd", "fullname").Create(&userInfo)
	//close connect
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	if dbResult.Error != nil {
		return 5, "acct/pwd repeat or error."
	} else {
		return 0, "success"
	}
}
