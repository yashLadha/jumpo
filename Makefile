.DEFAULT_GOAL := build

build-mac:
	GOOS=darwin go build -o jumpo-mac

build-linux:
	GOOS=linux go build -o jumpo-linux

build: build-mac build-linux
