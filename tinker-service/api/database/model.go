package database

type DatabaseService struct {
	Id              int64  `json:"id"`
	ContainerName   string `json:"containerName"`
	DatabaseVendor  string `json:"vendor"`
	DefaultDatabase string `json:"defaultDatabase"`
	UserName        string `json:"username"`
	Password        string `json:"password"`
	Port            string `json:"port"`
}

