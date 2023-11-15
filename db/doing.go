package db

import "time"

type Doing struct {
	DoingId      int       `xorm:"pk not null unique autoincr"`
	DoingTitle   string    `xorm:"not null"`
	DoingContent string    `xorm:"null"`
	DoingComment string    `xorm:"null"`
	DoingType    int       `xorm:"not null"`
	UserId       int       `xorm:"not null"` // Foreign key
	CreateTime   time.Time `xorm:"not null"`
	DeadlineTime time.Time `xorm:"null"` // Deadline will be the next day of created day if null
}

func (d *Doing) TableName() string {
	return TableDoing
}

func QueryDoingById(id int) (*Doing, error) {
	d, err := queryById(id, new(Doing))
	return d.(*Doing), err
}
