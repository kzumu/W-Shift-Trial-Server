package wsct

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbTest() error {

	db, err := sql.Open("mysql", "kazu:password@tcp(localhost:3306)/wshift_trial?parseTime=true&collation=utf8_general_ci&interpolateParams=true")
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`CREATE TABLE HOGE (
	 id INTEGER,
	 hoge VARCHAR(255))`,
	)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`INSERT INTO HOGE (id, hoge)
		VALUES (?, ?) `,
		1,
		"hogege",
	)
	if err != nil {
		return err
	}

	return nil
}

func New() {

}
