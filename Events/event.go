package events

import (
	"api/db"
	"fmt"
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
	INSERT INTO events (name,description,location,datetime,user_id)(
		VALUES ($1, $2, $3, $4, $5)
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

func GetAll() ([]Event, error){
	var query string= `
	SELECT * FROM events
	` 
	rows,err:=db.DB.Query(query)

	if err!=nil{
		return nil,err
	}

	defer rows.Close()

	var events []Event;

	for rows.Next(){
		var event Event
		fmt.Print(rows)
		rows.Scan(&event.Id,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserId);
		events=append(events, event)
	}
	return events,nil
}