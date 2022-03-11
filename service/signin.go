package service

import (
	"time"

	pear "impl-api/pear"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
)

type UserLogin struct {
	Acct string
	Pwd  string
}
type Claims struct {
	Acct string `json:"acct"`
	jwt.StandardClaims
}
type token struct {
	JWTtoken string `json:"JWTtoken"`
}

func (v UserLogin) TableName() string {
	return "users"
}
func Signin(userInfo UserLogin) (ResultCode int, result interface{}) {
	//Chcek acct
	data := UserLogin{}
	db := DbConnect("ui")
	pwd := pear.PasswordMd5(userInfo.Pwd)
	db.Where("acct = ? and pwd = ?", userInfo.Acct, pwd).Find(&data)
	//close connect
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if (UserLogin{} == data) {
		return 5, "acct/pwd error."
	} else {
		if JWT, error := getJWT(userInfo.Acct); error != nil {
			return 5, error
		} else {
			return 0, token{JWTtoken: JWT}
		}
	}

}

func getJWT(acct string) (string, error) {
	secretKey := pear.GetBytes("SECRET_KEY")
	claims := Claims{
		Acct: acct,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 6).Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(secretKey)

}
