# Role-Permissions Helper Package
This package supports working with Roles and Permissions

## Install
```bash
go get github.com/Limpid-LLC/roles-permissions-helper
```

## Update version
```bash
go get -u github.com/Limpid-LLC/roles-permissions-helper@v1.3.1
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/Limpid-LLC/roles-permissions-helper"
)

func main() {

	// without parsing "$stoId" etc.
	ownerRolePermissionsConfig := role_permissions_config.GetOwnerRoleConfig()

	// map for change "$stoId" to "any_sto_internal_id_you_want"
	paramsToChange := make(map[string]string)
	paramsToChange["$stoId"] = "any_sto_internal_id_you_want"

	// with parsing "$stoId" etc.
	ownerRolePermissionsParsed, errParsing := role_permissions_parser.GetOwnerRoleParsed(paramsToChange)
	if errParsing != nil {
		fmt.Printf("Error parsing role: %v", errParsing)
		return
	}
}
```
