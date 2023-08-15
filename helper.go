package roles_permissions_helper

import (
	"os"
)

const (
	OwnerRolePermissionFilename = "owner_role_permissions.yml"
)

func GetOwnerRoleContent() ([]byte, error) {
	return os.ReadFile(OwnerRolePermissionFilename)
}
