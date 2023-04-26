.PHONY: createdb dropdb postgres start stop migrateup migratedown sqlc test server mock clean

createdb:
	docker exec -it postgres createdb --username=root --owner=root domain_design
dropdb:
	docker exec -it postgres dropdb domain_design

start:
	docker compose up -d

clean:
	docker compose down -v && rm -rf docker

stop:
	docker compose down
migrateup:
	docker run --rm \
	-v ${PWD}/db/migration:/migrations \
	--network host \
	migrate/migrate -path=/migrations/ -database "postgresql://root:secret@localhost:5432/domain_design?sslmode=disable" -verbose up
migratedown:
	docker run --rm \
	-v ${PWD}/db/migration:/migrations \
	--network host \
	migrate/migrate -path=/migrations/ -database "postgresql://root:secret@localhost:5432/domain_design?sslmode=disable" -verbose down
sqlc:
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...
