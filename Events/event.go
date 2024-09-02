package events

import (
	"api/db"
	"time"
)

type Event struct {
	Id          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserId int
}

var events = []Event{}

func (event Event) Save() error{
	var query string= `
	INSERT INTO events (name,description,location,dateTime,user_Id)(
		VALUES (?,?,?,?,?)
	)
	`
	sqlStm,err:=db.DB.Prepare(query)
	
	if err!=nil{
		return err;
	}
	defer sqlStm.Close()
	result,err:=sqlStm.Exec(event.Name,event.Description,event.Location,event.DateTime,event.UserId)

	if err!=nil{
		return err;
	}
	id,err:=result.LastInsertId()

	if err!=nil{
		return err;
	}
	event.Id=id;
	return nil
}

func GetAll() error{
	var query string= `
	SELECT * FROM events
	` 
	rows,err:=db.DB.Query(query)

	if err!=nil{
		return err
	}

	defer rows.Close()

	var events []Event;

	for rows.Next(){
		var event Event
		rows.Scan(&event.Id,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserId);
		events=append(events, event)
	}
	return nil
}