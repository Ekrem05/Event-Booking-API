package models

import (
	"api/db"
	"api/utils"
	"errors"
)

type User struct {
	Id       int64
	Email    string `binding: "required"`
	Password string `binding: "required"`
}

func (user User) Save() error {

	var query string = `INSERT INTO users(email,password) VALUES ($1,$2)`
	hashedPassword,err:=utils.HashPassword(user.Password);
	if err!=nil{
		return err;
	}
	_,err=db.DB.Exec(query,user.Email,hashedPassword);

	if err!=nil{
		return err;
	}
	return nil
}

func (user User) Validate() (string,error) {

	validEmailQuery:=`SELECT * FROM users 
	WHERE email=$1`

	row:=db.DB.QueryRow(validEmailQuery,user.Email)
	
	var userFound User 
	err:=row.Scan(&userFound.Id,&userFound.Email,&userFound.Password)

	if err!=nil{
		return "",err
	}

	isValid:=utils.CheckPassword(user.Password,userFound.Password)

	if !isValid{
		return "",errors.New("Invalid password")
	}

	token,err:=utils.Generate(userFound.Email,userFound.Id)
	if err!=nil{
		return "",err
	}

	return token,nil
}