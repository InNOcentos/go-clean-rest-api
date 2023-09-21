
build:
	go mod download && GOOS=darwin go build -o ./.bin/app ./cmd/main.go

run: build
	docker-compose up --build postgres

migrate-create:
	migrate create -ext sql -dir migrations/ -seq $(name)

migrate-up:
	migrate -path migrations/ -database postgresql://root:pass@localhost:5432/test?sslmode=disable up

