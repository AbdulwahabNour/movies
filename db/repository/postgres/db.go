package postgres

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
    maxOpenDbConn =10
    maxIdleConn = 5
    maxDbLifetime = 5 * time.Minute
)

 
type Database struct {
    Client *sqlx.DB
}


func ConnectSql(dsn string)(*Database, error){
    db, err := NewDatabase(dsn)
    if err != nil{
        panic(err)
    }

    db.SetMaxOpenConns(maxOpenDbConn)
    db.SetMaxIdleConns(maxIdleConn)
    db.SetConnMaxLifetime(maxDbLifetime)

    if err:= db.Ping(); err != nil{
        return nil, err
    }

    return &Database{Client: db}, nil
}



func NewDatabase(dataConn string) (*sqlx.DB, error) {
    conn, err := sqlx.Connect("postgres", dataConn)

    if err != nil {
        return nil, errors.New("could not connect to database ", err)
    }
    
    return conn, nil
}