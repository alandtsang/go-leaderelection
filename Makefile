Output := main

.PHONY: build clean

build:
	GO111MODULE=on go build -mod vendor -o $(Output) cmd/main/main.go

clean:
	@rm -rf $(Output) 
