package postgres

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

func(psqldb *PsqlDB)MigrateDB() error{
    
     driver, err := postgres.WithInstance(psqldb.Client.DB, &postgres.Config{})
     if err != nil{
         return fmt.Errorf("could not create the postgres driver: %w ", err)
     }

     m, err := migrate.NewWithDatabaseInstance(
         "file://db/migrations", "postgres", driver,
     )

     if err != nil{
         return fmt.Errorf("could not create the postgres driver: %w ", err)
     }

     if err = m.Up(); err != nil{
        if errors.Is(err, migrate.ErrNoChange){
            return fmt.Errorf("could not migrate the database: %w ", err)
        }
        
     }
  
    return nil
}