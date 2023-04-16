DB_NAME ?= gobank

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root $(DB_NAME)


dropdb:
	docker exec -it postgres12 dropdb --username=root --owner=root $(DB_NAME) 

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/$(DB_NAME)?sslmode=disable" -verbose up


sqlc:
	sqlc generate

test:
	go test ./...


.PHONY: createdb dropdb postgres migrateup migratedown sqlc test