package docker

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func commit(containerID, imageName string) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	commitResp, err := cli.ContainerCommit(ctx, containerID, types.ContainerCommitOptions{
		Reference: imageName,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(commitResp.ID)
}
