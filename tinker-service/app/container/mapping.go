package container

import "errors"

type VendorPort struct {
	Image string
	Port  string
}

func vendorToDockerImage(vendor string) (*VendorPort, error) {
	switch vendor {
	case "postgres":
		return &VendorPort{"postgres:14.4-alpine", "5432"}, nil
	case "mysql":
		return &VendorPort{"mysql/mysql-server:8.0-debian", "3306"}, nil
	default:
		return nil, errors.New("unsupported vendor")
	}
}
