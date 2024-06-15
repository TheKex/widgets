.PHONY: postgres create-db drop-db migrate-up migrate-down sqlc test server mock

DB_IMAGE = postgres:12-alpine
DB_CONTAINER = postgres12
DB_NAME = widget_app
DB_PORT = 5432
DB_USER = root
DB_PASSWORD = secret


postgres:
	docker run --name $(DB_CONTAINER) -p $(DB_PORT):5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d $(DB_IMAGE)

create-db:
	docker exec -it $(DB_CONTAINER) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

drop-db:
	docker exec -it $(DB_CONTAINER) dropdb $(DB_NAME)

migrate-up:
	migrate -path db/migration -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mock_db -source="db/sqlc/store.go" -destination="db/mock/store.go"