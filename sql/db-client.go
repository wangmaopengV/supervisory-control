package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBClient struct {
	dbClient *sql.DB
}

func NewDBClient(endpoint string, keySpace string, userName string, password string) (*DBClient, error) {

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s", userName, password, endpoint, keySpace)
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	client := &DBClient{
		dbClient: db,
	}
	return client, nil
}
