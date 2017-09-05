package structs

type Server struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Protocol       string `json:"protocol"`
	RepositoryId   int    `json:"repository_id"`
	EnvironmentId  int    `json:"environment_id"`
	PreDeployHook  string `json:"pre_deploy_hook"`
	PostDeployHook string `json:"post_deploy_hook"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type Servers struct {
	Meta
	Entries []Server `json:"entries"`
}
