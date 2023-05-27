package action

import "errors"

func buildComposeUp(f []string, isBuild bool) []string {
	var compose = []string{"compose"}
	for _, s := range f {
		compose = append(compose, "-f", s)
	}
	compose = append(compose, "up", "-d")
	if isBuild {
		compose = append(compose, "--build")
	}
	return compose
}

func buildComposeDown(f []string) []string {
	var compose = []string{"compose"}
	for _, s := range f {
		compose = append(compose, "-f", s)
	}
	compose = append(compose, "down")
	return compose
}

func parseArguments(args []string) ([]string, error) {
	if len(args) == 0 {
		return nil, errors.New("at least one argument is expected")
	}
	var composeFiles []string
	for _, s := range args {
		file, ok := serviceMap[s]
		if ok {
			composeFiles = append(composeFiles, file)
		} else {
			println("service not found")
		}
	}
	return composeFiles, nil
}
