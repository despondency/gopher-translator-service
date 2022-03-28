test:
	go test ./...

build-image:
	docker build . -f build/Dockerfile -t gopher-translator-service:1.0

run:
	docker-compose -f ./deployments/docker-compose-local.yml up