package service

type userListStruct struct {
	Acct     string `json:"acct"`
	Fullname string `json:"fullname"`
}

func (v userListStruct) TableName() string {
	return "users"
}
func UserList(page int) (ResultCode int, result interface{}) {
	page = (page - 1) * 10
	db := DbConnect("ui")
	data := []userListStruct{}
	db.Order("created_at asc").Limit(10).Offset(page).Find(&data)
	//close connect
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	return 0, data
}
