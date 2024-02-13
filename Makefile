.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/credit main.go
	
clean:
	rm -rf ./bin ./vendor Gopkg.lock

