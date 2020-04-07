EXE_NAME=covid-19-etl

build:
	@go build -o ${EXE_NAME} ./cmd

build_for_ubuntu:
	GOOS=linux GOARCH=amd64 go build -o ${EXE_NAME} ./cmd

clean:
	rm ${EXE_NAME}

migrate:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

run:
	@./covid-19-etl

test:
	@go test ./...

.PHONY: build build_for_ubuntu clean migrate run test
