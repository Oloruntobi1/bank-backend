migrateup:
	migrate -path db/migration -database "postgresql://bb:${DB_PASSWORD}@localhost:5432/bank_data?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://bb:${DB_PASSWORD}@localhost:5432/bank_data?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./... -count=1

.PHONY: migrateup migratedown sqlc test