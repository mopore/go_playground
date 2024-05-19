package main

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestExample(t *testing.T) {
	t.Log(("We are testing 'testing containers' package."))
	assert.Equal(t, 1, 1, "The two numbers should be the same.")

	req := testcontainers.ContainerRequest{
		Image:        "nginx:latest",
		ExposedPorts: []string{"80/tcp"},
		WaitingFor:   wait.ForLog("Configuration complete; ready for start up"),
	}

	ctx := context.Background()
	container, err := testcontainers.GenericContainer(
		ctx, 
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		},
	)
	assert.NoError(t, err)

	endpoint, err := container.Endpoint(ctx, "")
	assert.NoError(t, err)
	t.Logf("Container started on %s", endpoint)

	// Make a request to the container
	response, err := http.Get(fmt.Sprintf("http://%s", endpoint))
	assert.NoError(t, err)

	assert.Equal(t, 200, response.StatusCode, "Status code should be 200")

}
