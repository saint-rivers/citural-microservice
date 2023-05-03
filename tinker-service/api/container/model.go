package container

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

type DbContainerConfig struct {
	ContainerName   string
	ContainerConfig *container.Config
	HostConfig      *container.HostConfig
	NetworkConfig   *network.NetworkingConfig
}
