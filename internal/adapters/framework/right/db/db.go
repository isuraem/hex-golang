package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

type Adapter struct {
	DB *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failure %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("DB ping failure %v", err)
	}

	return &Adapter{DB: db}, nil
}

func (da *Adapter) CloseDbConnection() {
	err := da.DB.Close()
	if err != nil {
		log.Fatalf("DB close failure: %v", err)
	}
}

func (da *Adapter) AddToHistory(answer int32, operation string) error {
	queryString, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").
		Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		return err
	}

	_, err = da.DB.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}
