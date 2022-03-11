package service

type userSearch struct {
	Acct     string `json:"acct"`
	Fullname string `json:"fullname"`
}

func (v userSearch) TableName() string {
	return "users"
}

func SearchFullname(key string) (ResultCode int, result interface{}) {
	db := DbConnect("ui")
	key = "%" + key + "%"
	data := []userSearch{}
	db.Where("fullname like ? ", key).Find(&data)
	//close connect
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	return 0, data
}
