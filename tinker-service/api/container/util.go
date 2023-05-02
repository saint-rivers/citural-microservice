package container

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"

	"github.com/saint-rivers/tinker/common"
	"github.com/saint-rivers/tinker/config"
)

var (
	cli = config.GetDockerClient()
	ctx = context.Background()
)

func toContainerConfig(r *request.DatabaseServiceRequest) *container.Config {
	dbEnv := fmt.Sprintf("POSTGRES_DB=%s", r.DefaultDatabase)
	userEnv := fmt.Sprintf("POSTGRES_USER=%s", r.UserName)
	passEnv := fmt.Sprintf("POSTGRES_PASSWORD=%s", r.Password)

	return &container.Config{
		Hostname: r.ContainerName,
		Image:    "postgres:14.4-alpine",
		Env:      []string{dbEnv, userEnv, passEnv},
	}
}

func CreateDatabaseContainer(r *request.DatabaseServiceRequest) container.CreateResponse {
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port("5432"): []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: r.Port}},
		},
	}
	container, err := cli.ContainerCreate(ctx, toContainerConfig(r), hostConfig, nil, nil, r.ContainerName)
	if err != nil {
		log.Println(err)
	}
	return container
}

func StartDatabaseContainer(cid string) {
	cli.ContainerStart(ctx, cid, types.ContainerStartOptions{})
	log.Printf("Container %s is created", cid)
}
