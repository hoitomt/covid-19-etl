EXE_NAME=covid-19-etl

build:
	go build -o ${EXE_NAME}

build_for_ubuntu:
	GOOS=linux GOARCH=amd64 go build -o ${EXE_NAME}

clean:
	rm ${EXE_NAME}

run:
	./covid-19-etl

.PHONY: build clean run build_for_ubuntu