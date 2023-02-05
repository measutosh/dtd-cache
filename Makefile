# .PHONY: run
# run: main
#	 ./$<
#
# main: *.go go.mod
# 	go build -o $@ .
#
# .PHONY: all
# all: main

build:
	go build -o bin/dtb-cache

run: build
	./bin/dtb-cache

runfollower: build
	./bin/dtb-cache --listenaddr :4000 --leaderaddr :3000
test:
	@go test -v ./...