package role_permissions_config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"roles-permissions-helper/sdk/role_permissions_models"
)

const (
	OwnerRolePermissionFilename = "owner_role_permissions.yml"
)

func GetOwnerRoleConfig() *role_permissions_models.Role {
	// Read the content of the YAML file
	content, err := os.ReadFile(OwnerRolePermissionFilename)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return nil
	}

	// Parse the YAML content into the Config struct
	var role role_permissions_models.Role

	if err := yaml.Unmarshal(content, &role); err != nil {
		fmt.Printf("Error parsing YAML: %v", err)
		return nil
	}

	return &role
}
