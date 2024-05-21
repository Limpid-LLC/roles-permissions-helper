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

	return prepareRole(ownerRoleConfigLoaded, paramsToChange)
}

func GetSupportAdminRoleParsed(paramsToChange map[string]string) (*role_permissions_models.Role, error) {
	supportRoleConfigLoaded, errRoleLoad := role_permissions_config.GetSupportAdminRoleConfig()
	if errRoleLoad != nil {
		return nil, errRoleLoad
	}

	return prepareRole(supportRoleConfigLoaded, paramsToChange)
}

func GetPublicRoleParsed(paramsToChange map[string]string) (*role_permissions_models.Role, error) {
	publicRoleConfigLoaded, errRoleLoad := role_permissions_config.GetPublicRoleConfig()
	if errRoleLoad != nil {
		return nil, errRoleLoad
	}

	return prepareRole(publicRoleConfigLoaded, paramsToChange)
}

func prepareRole(loadedRole *role_permissions_models.Role, paramsToChange map[string]string) (*role_permissions_models.Role, error) {
	ownerRoleConfig, errRoleClone := loadedRole.Clone()
	if errRoleClone != nil {
		return nil, errRoleClone
	}

	for changeFrom, changeTo := range paramsToChange {
		ownerRoleConfig.ChangeParamInRole(changeFrom, changeTo)
	}

	return ownerRoleConfig, nil
}
