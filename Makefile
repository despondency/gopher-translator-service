test:
	go test ./...

test-unit:
	go test ./... -run ./internal

test-integration:
	make build-image && go test ./... -run ./integration-test

build-image:
	docker build . -f build/Dockerfile -t gopher-translator-service:1.0

run:
	docker-compose -f ./deployments/docker-compose-local.yml up --build --force-recreate

generate:
	go generate ./...
