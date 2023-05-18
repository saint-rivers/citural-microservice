package container

import (
	"context"
	"fmt"
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

type ContainerResponse struct {
	ContainerName string       `json:"containerName"`
	ContainerId   string       `json:"containerId"`
	Status        string       `json:"status"`
	Ports         []types.Port `json:"ports"`
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

func ListContainers() ([]ContainerResponse, error) {
	var containers []ContainerResponse

	listedContainers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		log.Fatal(err)
	}
	if len(listedContainers) > 0 {
		for _, container := range listedContainers {
			containerName := container.Names[0]

			c := ContainerResponse {
				ContainerId:   container.ID,
				ContainerName: containerName[1:],
				Status:        container.Status,
				Ports:         container.Ports,
			}
			containers = append(containers, c)
		}
	} else {
		fmt.Println("There are no containers running")
	}

	return containers, nil
}
