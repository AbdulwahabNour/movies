migrateDBUP:
	migrate -path ./migrations -database "postgresql://abdo:abdo@localhost:5432/moviesapp?sslmode=disable" -verbose up
migrateDBDown:
	migrate -path ./migrations -database "postgresql://abdo:abdo@localhost:5432/moviesapp?sslmode=disable" -verbose down
	