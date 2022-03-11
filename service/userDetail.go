package service

import (
	"time"
)

type userDetail struct {
	Acct       string    `json:"acct"`
	Fullname   string    `json:"fullname"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func (v userDetail) TableName() string {
	return "users"
}

func UserDetail(acct string) (ResultCode int, result interface{}) {
	db := DbConnect("ui")
	data := userDetail{}
	dbResult := db.Where("acct = ?", acct).Take(&data)
	//close connect
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if dbResult.Error != nil {
		return 5, "acct not exist."
	} else {
		return 0, data
	}
}
