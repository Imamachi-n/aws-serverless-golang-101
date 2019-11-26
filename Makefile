.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/getHelloWorld helloworld/getHelloWorld.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/postHelloWorld helloworld/postHelloWorld.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

remove:
	sls remove