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

//68282 Dev 1 Kiev sftp 33598 46319   2016/01/21 08:17:33 -0800 2017/05/17 08:31:23 -0700

type Servers struct {
	Meta    `json:"meta"`
	Entries []Server `json:"entries,omitempty"`
}
