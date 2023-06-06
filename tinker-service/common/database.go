package request

type DatabaseServiceRequest struct {
	ContainerName   string `json:"containerName"`
	DatabaseVendor  string `json:"vendor"`
	DefaultDatabase string `json:"defaultDatabase"`
	UserName        string `json:"username"`
	Password        string `json:"password"`
	Port            string `json:"port"`
}

func (r *DatabaseServiceRequest) SetContainerName(name string) {
	r.ContainerName = name
}
