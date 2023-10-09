package role_permissions_models

import (
	"gopkg.in/yaml.v2"
)

type Role struct {
	Id          string                 `json:"internal_id,omitempty" yaml:"internal_id,omitempty"`
	Type        string                 `json:"type" yaml:"type"`
	StoId       string                 `json:"sto_id" yaml:"sto_id"`
	Data        map[string]interface{} `json:"data" yaml:"data"`
	Permissions []RolePermission       `json:"permissions" yaml:"permissions"`
}

type RolePermission struct {
	Microservice     string  `json:"microservice" yaml:"microservice"`
	Method           string  `json:"method" yaml:"method"`
	RequiredParams   []Param `json:"required_params" yaml:"required_params"`
	RestrictedParams []Param `json:"restricted_params" yaml:"restricted_params"`
}

type Param struct {
	Param  string   `json:"param" yaml:"param"`
	Values []string `json:"values" yaml:"values"`
	All    bool     `json:"all" yaml:"all"`
}

func (role *Role) Clone() (*Role, error) {
	marshalledData, errM := yaml.Marshal(role)
	if errM != nil {
		return nil, errM
	}

	var clonedRole Role
	errUnM := yaml.Unmarshal(marshalledData, &clonedRole)
	if errUnM != nil {
		return nil, errUnM
	}

	return &clonedRole, nil
}

func (role *Role) ChangeParamInRole(changeFrom string, changeTo string) {
	if role.StoId == changeFrom {
		role.StoId = changeTo
	}

	for permKey, rolePermission := range role.Permissions {
		for reqParamKey, reqParam := range rolePermission.RequiredParams {
			for reqParamValueKey, reqParamValue := range reqParam.Values {
				if reqParamValue == changeFrom {
					role.Permissions[permKey].RequiredParams[reqParamKey].Values[reqParamValueKey] = changeTo
				}
			}
		}

		for resParamKey, resParam := range rolePermission.RestrictedParams {
			for resParamValueKey, resParamValue := range resParam.Values {
				if resParamValue == changeFrom {
					role.Permissions[permKey].RestrictedParams[resParamKey].Values[resParamValueKey] = changeTo
				}
			}
		}
	}
}
