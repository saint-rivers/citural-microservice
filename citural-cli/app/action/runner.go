package action

import (
	"github.com/urfave/cli"
)

var serviceMap = map[string]string{
	"keycloak": "keycloak-service/docker-compose.yml",
	"tinker":   "tinker-service/docker-compose.yml",
}

func StopAction(c *cli.Context) error {
	composeFiles, _ := parseArguments(c.Args())
	args := buildComposeDown(composeFiles)
	runDockerCompose(args)
	return nil
}

func RunAction(c *cli.Context) error {
	composeFiles, _ := parseArguments(c.Args())
	args := buildComposeUp(composeFiles, c.IsSet("build"))
	runDockerCompose(args)
	return nil
}
