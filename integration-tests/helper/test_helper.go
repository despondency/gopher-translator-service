package helper

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type serviceContainer struct {
	testcontainers.Container
	URI string
}

func SetupService(ctx context.Context) (*serviceContainer, error) {
	req := testcontainers.ContainerRequest{
		AlwaysPullImage: false,
		Image:           "gopher-translator-service:1.0",
		AutoRemove:      true,
		ExposedPorts:    []string{"8080/tcp"},
		WaitingFor:      wait.ForHTTP("/v1/health").WithPort("8080"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	ip, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, "8080")
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("http://%s:%s", ip, mappedPort.Port())
	return &serviceContainer{Container: container, URI: uri}, nil
}
