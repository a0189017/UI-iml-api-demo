package service

import (
	"encoding/json"
	"errors"
	"fmt"
	pear "impl-api/pear"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,PATCH")
	w.Header().Set("Access-Control-Expose-Headers", "Cache-Control,Content-Language,Content-Type, Expires, Last-Modified")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//dbConnect
func DbConnect(Dbname string) (db *gorm.DB) {
	//配置MySQL
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	//host := "postgres"
	host := "localhost"
	port := 5432
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", host, username, password, Dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("連線失敗, error=" + err.Error())
	}

	return db
}
func JWTcheck(authToken string) (acct string, err error) {
	authToken = strings.Replace(authToken, "Bearer ", "", -1)
	secretKey := pear.GetBytes("SECRET_KEY")
	// parse and verify the token string
	tokenClaims, err := jwt.ParseWithClaims(authToken, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return secretKey, nil
	})
	// detail for jwt token err message
	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
		err = errors.New(message)
		return "", err
	}

	claims, _ := tokenClaims.Claims.(*Claims)
	// prevent userId from empty string
	if claims.Acct == "" {
		return "", errors.New("token is improper")
	}
	return claims.Acct, nil
}
