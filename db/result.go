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
