package role_permissions_parser

import (
	"github.com/Limpid-LLC/roles-permissions-helper/sdk/role_permissions_config"
	"github.com/Limpid-LLC/roles-permissions-helper/sdk/role_permissions_models"
)

func GetOwnerRoleParsed(paramsToChange map[string]string) (*role_permissions_models.Role, error) {
	ownerRoleConfigLoaded, errRoleLoad := role_permissions_config.GetOwnerRoleConfig()
	if errRoleLoad != nil {
		return nil, errRoleLoad
	}

	ownerRoleConfig, errRoleClone := ownerRoleConfigLoaded.Clone()
	if errRoleClone != nil {
		return nil, errRoleClone
	}

	for changeFrom, changeTo := range paramsToChange {
		ownerRoleConfig.ChangeParamInRole(changeFrom, changeTo)
	}

	return ownerRoleConfig, nil
}
