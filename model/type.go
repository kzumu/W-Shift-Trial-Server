package model

import "time"

type User struct {
	ID        int64
	Name      *string
	Uuid      string
	PartnerId *int
	Created   time.Time
	Updated   time.Time
}

type Shift struct {
	ID      int64
	UserId  int64
	Name    *string
	Color   *string
	Date    time.Time
	Created time.Time
	Updated time.Time
}

type TempKey struct {
	ID      int64
	UserID  int64
	TempKey int64
	Created time.Time
	Updated time.Time
}
