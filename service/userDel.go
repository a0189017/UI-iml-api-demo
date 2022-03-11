package service

type user struct{}

func (v user) TableName() string {
	return "users"
}

func UserDel(acct string) (ResultCode int, result interface{}) {
	db := DbConnect("ui")
	//
	user := user{}
	dbResult := db.Where("acct = ?", acct).Take(&user)

	//close connect
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if dbResult.Error != nil {
		return 5, "acct not exist."
	} else {
		db.Where("acct = ?", acct).Delete(&user)
		return 0, "success"
	}
}
