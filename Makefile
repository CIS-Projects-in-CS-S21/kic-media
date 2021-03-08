.DEFAULT_GOAL := run_tests

test_container:
	docker build -f test.Dockerfile -t media-tests .

run_tests: test_container
	docker network create tests
	docker run --rm -d -p 27017:27017 --net tests --name mongo mongo
	- docker run --rm --net tests media-tests
	docker stop mongo
	docker network rm tests

client:
	go build -o ./bin/client ./cmd/client/clientTest.go

build:
	go build -o ./bin/server ./cmd/server/server.go

push:
	docker build -t gcr.io/keeping-it-casual/kic-media:dev .
	docker push gcr.io/keeping-it-casual/kic-media:dev