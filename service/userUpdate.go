package service

import (
	pear "impl-api/pear"
	"time"
)

type UserUpdateInfo struct {
	Pwd        string
	Fullname   string
	Updated_at time.Time
}

func (v UserUpdateInfo) TableName() string {
	return "users"
}
func UserUpdate(updateInfo UserUpdateInfo, acct string) (ResultCode int, result interface{}) {
	update_at := time.Now()
	updateInfo.Updated_at = update_at
	updateInfo.Pwd = pear.PasswordMd5(updateInfo.Pwd)
	db := DbConnect("ui")
	db.Model(&UserUpdateInfo{}).Where("acct = ?", acct).Updates(updateInfo)
	//close connect
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	return 0, "success"

}
