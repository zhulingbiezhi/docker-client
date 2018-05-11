package docker

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func listContainerAll() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("|%-20s|%-32s|%-25s|%-20s|%-25s|%-25s|%s\n", "CONTAINER ID", "IMAGE", "COMMAND", "CREATED", "STATUS", "PORTS", "NAMES")
	for _, container := range containers {
		fmt.Printf("|%-20s|%-32s|%-25s|%-20d|%-25s|%-25s|%s\n", container.ID, container.Image, container.Command, container.Created, container.State, container.Ports, container.Names)
	}
}


func listImagesAll(){
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{
		All: true,
	})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}
}