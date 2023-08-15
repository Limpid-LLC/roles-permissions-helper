# Role-Permissions Helper Package
This package supports working with Roles and Permissions

## Install
```bash
go get github.com/Limpid-LLC/roles-permissions-helper
```

## Update version
```bash
go get -u github.com/Limpid-LLC/roles-permissions-helper@v1.1.2
```

## Usage
```go
import "github.com/Limpid-LLC/roles-permissions-helper"

func test() {
	
    // without parsing "$stoId" etc.
    ownerRolePermissionsConfig := role_permissions_config.GetOwnerRoleConfig()


	// map for change "$stoId" to "any_sto_internal_id_you_want"
    paramsToChange := make(map[string]string)
    paramsToChange["$stoId"] = "any_sto_internal_id_you_want"

    // with parsing "$stoId" etc.
    ownerRolePermissionsParsed, errParsing := role_permissions_parser.GetOwnerRoleParsed(paramsToChange)
    if errParsing != nil {
        log.Printf("Error parsing role: %v", errParsing)
        return
    }
}
```
