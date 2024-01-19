# Dependencies
dependencies:
	docker-compose -f ./deployment/dependencies/docker-compose.yaml up -d

# Database
create-migration:
	migrate create -ext sql -dir db/migrations -tz Local init

# --- for Macos
#migrate:
#	migrate -database 'postgres://simple_admin:simple_password@localhost:5454/simple_db?sslmode=disable' -path db/migrations up
#migrate-rollback:
#	migrate -database 'postgres://simple_admin:simple_password@localhost:5454/simple_db?sslmode=disable' -path db/migrations down

# ---- for linux
migrate:
	migrate -database 'postgresql://simple_admin:simple_password@localhost:5432/simple_db?sslmode=disable' -path db/migrations up
migrate-rollback:
	migrate -database 'postgresql://simple_admin:simple_password@localhost:5432/simple_db?sslmode=disable' -path db/migrations down
