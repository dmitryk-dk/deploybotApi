package structs

type Deployment struct {
	Id                  int    `json:"id"`
	RepositoryId        int    `json:"repository_id"`
	EnvironmentId       int    `json:"environment_id"`
	UserId              int    `json:"user_id"`
	DeployedVersion     string `json:"deployed_version"`
	DeployFromScratch   bool   `json:"deploy_from_scratch"`
	TriggerNotification bool   `json:"trigger_notifications"`
	IsAutomatic         bool   `json:"is_automatic"`
	Comment             string `json:"comment"`
	AuthorName          string `json:"author_name"`
	State               string `json:"state"`
	Retries             int    `json:"retries"`
	CreatedAt           string `json:"created_at"`
	DeployedAt          string `json:"deployed_at"`
}

type Deployments struct {
	Meta
	Entries []Deployment `json:"entries"`
}
