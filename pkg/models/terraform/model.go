package terraform

import (
	"github.com/mia-platform/mipy/pkg/models"
)

type Terraform struct {
	Name string
	config map[string]string
	templateType models.TemplateType
}

func (t *Terraform) Provision() error {
	return nil
}

func (t *Terraform) GetType() models.TemplateType {
	return t.templateType
}

func NewTerraform(name string, config map[string]string) *Terraform {
	return &Terraform{
		Name: name,
		config: config,
		templateType: models.Terraform,
	}
}
