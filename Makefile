create: 
	migrate create -ext sql -dir pkg/database/migration/ -seq $(name)

migration_up:
	migrate -path pkg/database/migration/ -database "postgresql://postgres:$(password)@localhost:5433/Pokemon?sslmode=disable" -verbose up

migration_down:
	migrate -path pkg/database/migration/ -database "postgresql://postgres:$(password)@localhost:5433/Pokemon?sslmode=disable" -verbose down

fix:
	migrate -path pkg/database/migration/ -database "postgresql://postgres:$(password)@localhost:5433/Pokemon?sslmode=disable" force $(version)

up:
	docker-compose down
	docker-compose up --build -d
	docker image prune -f

down:
	docker-compose down