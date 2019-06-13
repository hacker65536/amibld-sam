.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./api/api
	rm -rf ./cwe/cwe

	
build:
	GOOS=linux GOARCH=amd64 go build -o api/api ./api
	GOOS=linux GOARCH=amd64 go build -o cwe/cwe ./cwe
