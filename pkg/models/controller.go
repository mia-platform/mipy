package models

type TemplateType string
type PipelineType string

const (
	AzureDevOps PipelineType = "azure-devops"
	GithubActions PipelineType = "github-actions"
	TypeGitlab PipelineType = "gitlab"
)

const (
	Terraform TemplateType = "terraform"
)

type Resource struct {
	Type TemplateType
	Files []string
}

type TemplatesConfig struct {
	Type TemplateType
	Path string
}


type MipyConfig struct {
	BasePath string
	TemplatesConfig []TemplatesConfig
	LogLevel string
}

type Controller struct {
	Resources []Resource
	Pipeline PipelineType
	Config *MipyConfig
}
