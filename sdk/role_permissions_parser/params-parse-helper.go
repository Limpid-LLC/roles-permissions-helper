package role_permissions_parser

import (
	"roles-permissions-helper/sdk/role_permissions_config"
	"roles-permissions-helper/sdk/role_permissions_models"
)

func GetOwnerRoleParsed(paramsToChange map[string]string) (*role_permissions_models.Role, error) {
	ownerRoleConfig, errRole := role_permissions_config.GetOwnerRoleConfig().Clone()
	if errRole != nil {
		return nil, errRole
	}

	for changeFrom, changeTo := range paramsToChange {
		ownerRoleConfig.ChangeParamInRole(changeFrom, changeTo)
	}

	return ownerRoleConfig, nil
}
