package container

import (
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	request "github.com/saint-rivers/tinker/common"
)

func ConfigPostgres(r *request.DatabaseServiceRequest) DbContainerConfig {
	dbEnv := fmt.Sprintf("POSTGRES_DB=%s", r.DefaultDatabase)
	userEnv := fmt.Sprintf("POSTGRES_USER=%s", r.UserName)
	passEnv := fmt.Sprintf("POSTGRES_PASSWORD=%s", r.Password)
	envs := []string{dbEnv, userEnv, passEnv}

	defaults, _ := vendorToDockerImage(r.DatabaseVendor)

	newport, _ := nat.NewPort("tcp", defaults.Port)
	exposedPorts := map[nat.Port]struct{}{
		newport: {},
	}

	return DbContainerConfig{
		ContainerName: r.ContainerName,
		ContainerConfig: &container.Config{
			Hostname:     r.ContainerName,
			Image:        defaults.Image,
			Env:          envs,
			ExposedPorts: exposedPorts,
		},
		HostConfig:    hostConfig(r.Port, defaults.Port),
		NetworkConfig: networkConfig(),
	}
}
