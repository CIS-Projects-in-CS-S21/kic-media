.DEFAULT_GOAL := build

client:
	go build -o ./bin/client ./cmd/client/clientTest.go

build:
	go build -o ./bin/server ./cmd/server/server.go

push:
	docker build -t gcr.io/keeping-it-casual/kic-media:dev .
	docker push gcr.io/keeping-it-casual/kic-media:dev