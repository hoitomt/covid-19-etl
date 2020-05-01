EXE_NAME=covid-19-etl

build:
	@go build -o ${EXE_NAME} ./cmd

build_for_ubuntu:
	GOOS=linux GOARCH=amd64 go build -o ${EXE_NAME} ./cmd

clean:
	rm ${EXE_NAME}

create_test_db:
	docker-compose exec postgres createdb -U covid_user covid_test

dependencies:
	go mod vendor

migrate:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate_test:
	migrate -database ${POSTGRESQL_TEST_URL} -path db/migrations up

run:
	@./covid-19-etl

test:
	@ENV=test go test ./...

.PHONY: build build_for_ubuntu clean create_test_db dependencies migrate migrate_test run test
