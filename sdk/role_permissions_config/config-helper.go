package role_permissions_config

import (
	"embed"
	"fmt"
	"github.com/Limpid-LLC/roles-permissions-helper/sdk/role_permissions_models"
	"gopkg.in/yaml.v2"
	"io/fs"
	"path/filepath"
)

const (
	OwnerRolePermissionFilename = "owner_role_permissions.yml"
)

//go:embed config/owner_role_permissions.yml
var content embed.FS

func GetOwnerRoleConfig() (*role_permissions_models.Role, error) {
	//libRootPath, errPath := roles_permissions_helper.GetLibraryRootPath()
	//if errPath != nil {
	//	fmt.Printf("Error getting root path: %v", errPath)
	//	return nil, errPath
	//}

	path := filepath.Join("config", OwnerRolePermissionFilename)
	contentRole, err := fs.ReadFile(content, path)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return nil, err
	}

	// Parse the YAML content into the Config struct
	var role role_permissions_models.Role

	if err := yaml.Unmarshal(contentRole, &role); err != nil {
		fmt.Printf("Error parsing YAML: %v", err)
		return nil, err
	}

	return &role, nil
}
