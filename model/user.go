package model

import "database/sql"

func CreateUserByUUID(db *sql.DB, uuid string) (*int64, error) {
	q := `insert into users
			(uuid) values (?)`
	res, err := db.Exec(q, uuid)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &id, nil
}
