package container

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"

	request "github.com/saint-rivers/tinker/common"
	"github.com/saint-rivers/tinker/config"
)

var (
	cli = config.GetDockerClient()
	ctx = context.Background()
)

type PortMappings struct {
	HostPort      string
	ContainerPort string
}

func CreateDatabaseContainer(r *request.DatabaseServiceRequest) container.CreateResponse {
	c := ConfigPostgres(r)

	container, err := cli.ContainerCreate(ctx, c.ContainerConfig, c.HostConfig, c.NetworkConfig, nil, c.ContainerName)
	if err != nil {
		log.Println(err)
	}
	return container
}

func StartDatabaseContainer(cid string) error {
	err := cli.ContainerStart(ctx, cid, types.ContainerStartOptions{})
	log.Printf("Container %s is started", cid)
	return err
}

func StopDatabaseContainer(cid string) error {
	err := cli.ContainerStop(ctx, cid, container.StopOptions{})
	log.Printf("Container %s is stopped", cid)
	return err
}
