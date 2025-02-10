.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o petmap

.PHONY: push
push:
	scp petmap ali:/root/projects/