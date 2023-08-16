package role_permissions_config

import (
	"fmt"
	"github.com/Limpid-LLC/roles-permissions-helper/sdk/role_permissions_models"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

const (
	OwnerRolePermissionFilename = "owner_role_permissions.yml"
)

func GetOwnerRoleConfig() (*role_permissions_models.Role, error) {
	//libRootPath, errPath := roles_permissions_helper.GetLibraryRootPath()
	//if errPath != nil {
	//	fmt.Printf("Error getting root path: %v", errPath)
	//	return nil, errPath
	//}
	libRootPath := ""

	// Read the content of the YAML file
	content, err := os.ReadFile(filepath.Join(libRootPath, OwnerRolePermissionFilename))
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return nil, err
	}

	// Parse the YAML content into the Config struct
	var role role_permissions_models.Role

	if err := yaml.Unmarshal(content, &role); err != nil {
		fmt.Printf("Error parsing YAML: %v", err)
		return nil, err
	}

	return &role, nil
}
