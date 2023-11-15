package db

import "time"

type Result struct {
	ResultId     int       `xorm:"pk not null autoincr"`
	ResultOrder  int       `xorm:"not null"`
	ResultResult int       `xorm:"not null"` // 0 for undone while 1 for done
	ResultTime   time.Time `xorm:"not null"`
	ResultReason string    `xorm:"null"`     // Null for expiring
	UserId       int       `xorm:"not null"` // Primary key of user
	DoingId      int       `xorm:"not null"` // Primary key of doing item
}

func (res *Result) TableName() string {
	return TableResult
}

func QueryResultById(id int) (*Result, error) {
	r, err := queryById(id, new(Result))
	return r.(*Result), err
}
