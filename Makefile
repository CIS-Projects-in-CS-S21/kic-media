.DEFAULT_GOAL := build

build:
	go build -o ./bin ./cmd/server/server.go

push:
	docker build -t gcr.io/keeping-it-casual/kic-media:dev .
	docker push gcr.io/keeping-it-casual/kic-media:dev