.PHONY: test build docker

help:
	echo "#################### OPTIONS ####################"
	echo "## test: run test suite                        ##"
	echo "## build: build application binary ##"
	echo "## docker: build application docker image      ##"
	echo "#################################################"

test:
	go test ./...

build: test
	go build .

docker: test
	docker build -t jvrmaia/dist .