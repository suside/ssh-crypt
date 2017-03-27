CURRENT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
VERSION := $(shell git describe --tag --always --long)

build: ssh-crypt ssh-crypt.exe ssh-crypt_darwin

ssh-crypt: $(wildcard *.go) $(wildcard **/*.go)
	go get ./...
	go get github.com/stretchr/testify/assert
	go test -v github.com/suside/ssh-crypt/lib
	go build -o ssh-crypt -i -ldflags "-s -w -X main.version=${VERSION}"

ssh-crypt.exe: ssh-crypt
	env GOOS=windows go build -o ssh-crypt.exe -i -ldflags "-s -w -X main.version=${VERSION}"

ssh-crypt_darwin: ssh-crypt
	env GOOS=darwin go build -o ssh-crypt_darwin -i -ldflags "-s -w -X main.version=${VERSION}"
