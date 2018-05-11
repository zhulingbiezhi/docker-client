package docker

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func stop(containerID string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if containerID == container.ID {
			if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
				panic(err)
			}
			fmt.Println("Stop Success")
		}
	}
}
