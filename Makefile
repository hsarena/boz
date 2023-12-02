db:
	podman run --name db -p 5432:5432 --env POSTGRES_USER=root --env POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	podman exec db createdb --username=root --owner=root $DB_NAME

dropdb:
	podman exec db dropdb $DB_NAME

migrationup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/$DB_NAME?sslmode=disable" -verbose up

migrationdown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/$DB_NAME?sslmode=disable" -verbose down -f

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: db creatdb dropdb migrationup migrationdown sqlc test