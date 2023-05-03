package container

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"

	"github.com/saint-rivers/tinker/common"
	"github.com/saint-rivers/tinker/config"
)

var (
	cli = config.GetDockerClient()
	ctx = context.Background()
)

func VendorToDockerImage(vendor string) (string, error) {
	switch vendor {
	case "postgres":
		return "postgres:14.4-alpine", nil
	case "mysql":
		return "mysql/mysql-server:8.0-debian", nil
	default:
		return "", errors.New("unsupported vendor")
	}
}

// func toContainerConfig(r *request.DatabaseServiceRequest) *container.Config {
// 	dbEnv := fmt.Sprintf("POSTGRES_DB=%s", r.DefaultDatabase)
// 	userEnv := fmt.Sprintf("POSTGRES_USER=%s", r.UserName)
// 	passEnv := fmt.Sprintf("POSTGRES_PASSWORD=%s", r.Password)
// 	image, err := VendorToDockerImage(r.DatabaseVendor)

// 	// containerPort, err := nat.NewPort("tcp", "80")

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	return &container.Config{
// 		Hostname: r.ContainerName,
// 		Image:    image,
// 		Env:      []string{dbEnv, userEnv, passEnv},
// 		// ExposedPorts: containerPort,
// 	}
// }

type VendorPort struct {
	Vendor string
	Port   uint32
}

func CreateDatabaseContainer(r *request.DatabaseServiceRequest) container.CreateResponse {
	dbEnv := fmt.Sprintf("POSTGRES_DB=%s", r.DefaultDatabase)
	userEnv := fmt.Sprintf("POSTGRES_USER=%s", r.UserName)
	passEnv := fmt.Sprintf("POSTGRES_PASSWORD=%s", r.Password)

	// Define a PORT opening
	newport, err := nat.NewPort("tcp", "5432")

	if err != nil {
		log.Panic("Unable to create docker port")
	}
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			newport: []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: r.Port,
				},
			},
		},
	}
	networkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{},
	}
	gatewayConfig := &network.EndpointSettings{
		Gateway: "gatewayname",
	}
	networkConfig.EndpointsConfig["bridge"] = gatewayConfig

	exposedPorts := map[nat.Port]struct{}{
		newport: {},
	}
	image, _ := VendorToDockerImage(r.DatabaseVendor)

	c := &container.Config{
		Hostname:     r.ContainerName,
		Image:        image,
		Env:          []string{dbEnv, userEnv, passEnv},
		ExposedPorts: exposedPorts,
	}

	container, err := cli.ContainerCreate(ctx, c, hostConfig, networkConfig, nil, r.ContainerName)
	if err != nil {
		log.Println(err)
	}
	return container
}

func StartDatabaseContainer(cid string) {
	cli.ContainerStart(ctx, cid, types.ContainerStartOptions{})
	log.Printf("Container %s is created", cid)
}
