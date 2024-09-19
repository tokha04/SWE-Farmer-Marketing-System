postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=changeme -d postgres:16.3

createdb:
	docker exec -it postgres createdb --username=root --owner=root fms

dropdb:
	docker exec -it postgres dropdb fms

migrateup:
	migrate -path db/migration -database "postgresql://root:changeme@localhost:5432/fms?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:changeme@localhost:5432/fms?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock