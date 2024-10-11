package utils

import (
	"fmt"
	"os"
)


// the function will return a list of templates types from local filesystem
// manifests/{template_name}/environments/{ENVIRONMENT_NAME}/
func GetTemplatesTypes() []string {

	dir, err := os.ReadDir("./manifests")
	if err != nil {
		fmt.Println(err)
	}

	var templatesTypes []string
	for _, d := range dir {
		templatesTypes = append(templatesTypes, d.Name())
	}

	return templatesTypes
}