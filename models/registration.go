package models

import "api/db"

type Registration struct {
	Id       int64
	User_id  int64 `binding:"required"`
	Event_id int64 `binding:"required"`
}

func (registration *Registration) New() error {
	
	query:=`INSERT INTO registrations(user_id,event_id)(
		VALUES($1,$2)
	)`
		
	_,err:=db.DB.Exec(query,registration.User_id,registration.Event_id)
	if err!=nil{
		return err
	}

	return nil;
}

func (registration *Registration) Remove() error {
	
	query:=`DELETE FROM registrations
	WHERE user_id=$1 and event_id=$2`
		
	_,err:=db.DB.Exec(query,registration.User_id,registration.Event_id)
	if err!=nil{
		return err
	}

	return nil;
}

func DoesExist(registration *Registration) error {
	
	query:=`SELECT * FROM registrations
	WHERE user_id=$1 and event_id=$2`
		
	row:=db.DB.QueryRow(query,registration.User_id,registration.Event_id)
	
	row.Err()

	return nil;
}