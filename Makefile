test:
	go test ./...

test-integration:
	go test ./... -run ./integration-test

build-image:
	docker build . -f build/Dockerfile -t gopher-translator-service:1.0

run:
	docker-compose -f ./deployments/docker-compose-local.yml up --build --force-recreate