create: 
	~/go/bin/migrate create -ext sql -dir pkg/database/migration/ -seq $(name)

migration_up:
	~/go/bin/migrate -path pkg/database/migration/ -database "postgresql://postgres:$(password)@localhost:5433/Pokemon?sslmode=disable" -verbose up

migration_down:
	~/go/bin/migrate -path pkg/database/migration/ -database "postgresql://postgres:$(password)@localhost:5433/Pokemon?sslmode=disable" -verbose down

fix:
	~/go/bin/migrate -path pkg/database/migration/ -database "postgresql://postgres:$(password)@localhost:5433/Pokemon?sslmode=disable" force $(version)

up:
	docker-compose down
	docker-compose up --build -d
	docker image prune -f

down:
	docker-compose down