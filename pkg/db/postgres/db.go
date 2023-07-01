package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AbdulwahabNour/movies/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
    maxOpenDbConn =25
    maxIdleConn = 25

    maxDbIdleLifetime = 15 * time.Minute
)

 
type PsqlDB struct {
    Client *sqlx.DB
}


func ConnectSql(config *config.Config)(*PsqlDB, error){
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    db, err := NewDatabase(config)
    if err != nil{
        panic(err)
    }
 
    db.SetMaxOpenConns(maxOpenDbConn)
    db.SetMaxIdleConns(maxIdleConn)
    db.SetConnMaxIdleTime(maxDbIdleLifetime)
 
 
    if err:= db.PingContext(ctx); err != nil{
        return nil, err
    }

    return &PsqlDB{Client: db}, nil
}



func NewDatabase(config *config.Config) (*sqlx.DB, error) {
    dataConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", 
                            config.Postgres.PostgresqlHost,
                            config.Postgres.PostgresqlPort,
                            config.Postgres.PostgresqlUser,
                            config.Postgres.PostgresqlDbname,
                            config.Postgres.PostgresqlPassword)
    ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
    defer cancel()
    conn, err := sqlx.ConnectContext(ctx, "postgres", dataConn)
  
    if err != nil {
        return nil, errors.New("could not connect to database")
    }
    
    return conn, nil
}