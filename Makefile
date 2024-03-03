postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

start:
	docker start postgres 
createdb: 
	docker exec -it postgres createdb --username=root --owner=root simple-bank
dropdb:
	docker exec -it postgres dropdb  simple-bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown