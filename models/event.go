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

func GetById(id int64) (*Event,error){

	var query string=`SELECT *
	FROM events
	WHERE id = $1`

	sqlStm,err:=db.DB.Prepare(query)
	
	if err!=nil{
		return nil,err;
	}
	defer sqlStm.Close();

	row:=sqlStm.QueryRow(id)


	var event Event;

	scanErr:=row.Scan(&event.Id,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserId);

	if scanErr!=nil{
		return nil,scanErr;
	}
	
	return &event,nil
}

func UpdateEvent(event *Event) error {
	var query string = `UPDATE events
	SET name = $1, description = $2, location = $3, datetime = $4
	WHERE id = $5
	`

	sqlStm,err:=db.DB.Prepare(query)
	
	if err!=nil{
		return err;
	}
	defer sqlStm.Close();

	_,err=sqlStm.Exec(event.Name,event.Description,event.Location,event.DateTime,event.Id);
	if err!=nil{
		return err;
	}
	return nil;
} 

func DeleteEvent(id int64) error {
	var query string = `DELETE FROM events WHERE id = $1`;
	sqlStm,err:=db.DB.Prepare(query)
	
	if err!=nil{
		return err;
	}
	defer sqlStm.Close();

	_,err=sqlStm.Exec(id);
	if err!=nil{
		return err;
	}
	return nil
}