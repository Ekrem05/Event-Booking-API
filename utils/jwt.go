package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey="i328hr9quadn10921nd19002smnbg4n891"

func Generate(email string, userId int64) (string,error) {
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
		"userId":userId,
		"exp":time.Now().Add(time.Hour*2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func Verify(token string) (int64,error) {

	parsedToken,err:=jwt.Parse(token,func(t *jwt.Token) (interface{}, error) {

		_,ok:=t.Method.(*jwt.SigningMethodHMAC);

		if !ok{
			return nil,errors.New("Unexpected signing method")
		}
		return []byte(secretKey),nil;
	})

	if err!=nil{
		return 0,errors.New("Coud not parse token")
	}

	isValid:=parsedToken.Valid

	if !isValid{
		return 0,errors.New("Invalid token")
	}

	claims:=parsedToken.Claims.(jwt.MapClaims);
	userId:=int64(claims["userId"].(float64))

	return userId,nil
}