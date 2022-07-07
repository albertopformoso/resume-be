# Build app and db
build:
	docker-compose -f ./docker/docker-compose.yml up -d --build

# Remove app and db
remove:
	docker-compose -f ./docker/docker-compose.yml down

# Clean docker system
clean: remove
	docker system prune -fa

# DB migration up
migrateup:
	./scripts/migrate.sh up

# DB migration down
migratedown:
	./scripts/migrate.sh down

populate:
	./scripts/populate.sh
