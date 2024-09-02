package models

import (
	"api/db"
	"api/utils"
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