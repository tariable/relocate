include .env

migrate-up:
	docker run  -v ${MIGRATIONS_LOCAL_PATH}:/migrations --network host migrate/migrate -path=/migrations -database ${DB_DRIVER}://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DBNAME}?sslmode=${DB_SSLMODE} up
migrate-down:
	docker run  -v ${MIGRATIONS_LOCAL_PATH}:/migrations --network host migrate/migrate -path=/migrations -database ${DB_DRIVER}://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DBNAME}?sslmode=${DB_SSLMODE} down -all
