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

// type DbContainer struct {
// 	ContainerName string
// 	Image         string
// 	Tag           string `default:"latest"`
// 	Network       string `default:"bridge"`
// 	PortMappings  []PortMappings
// 	Env           []string
// }

func CreateDatabaseContainer(r *request.DatabaseServiceRequest) container.CreateResponse {
	db := ConfigPostgres(r)

	container, err := cli.ContainerCreate(
		ctx, db.ContainerConfig, db.HostConfig, db.NetworkConfig, nil, db.ContainerName,
	)
	if err != nil {
		log.Println(err)
	}
	return container
}

func StartDatabaseContainer(cid string) {
	cli.ContainerStart(ctx, cid, types.ContainerStartOptions{})
	log.Printf("Container %s is created", cid)
}
