package service

import "time"

type updateFullname struct {
	Fullname   string
	Updated_at time.Time
}

func (v updateFullname) TableName() string {
	return "users"
}
func UserUpdateFullname(fullname string, acct string) (ResultCode int, result interface{}) {
	db := DbConnect("ui")
	update_at := time.Now()
	updateInfo := updateFullname{fullname, update_at}
	db.Model(&updateFullname{}).Where("acct = ?", acct).Updates(updateInfo)

	return 0, "success"
}
