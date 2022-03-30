test:
	make build-image && go test ./...

test-unit:
	go test ./internal/...

test-integration:
	make build-image && go test ./integration-tests/...

build-image:
	docker build . -f build/Dockerfile -t gopher-translator-service:1.0

run-container:
	docker-compose -f ./deployments/docker-compose-local.yml up --build --force-recreate

generate:
	go generate ./...
