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

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

server: 
	go run main.go
mock: 
	mockgen -package mockdb -destination db/mock/store.go  github.com/SaifAlqady51/simple-bank/db/sqlc Store 

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock