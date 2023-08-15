# Role-Permissions Helper Package
This package supports working with Roles and Permissions

## Install
```bash
go get github.com/Limpid-LLC/roles-permissions-helper
```

## Update version
```bash
go get -u github.com/Limpid-LLC/roles-permissions-helper@v1.0.1
```

## Usage
```go
import "github.com/Limpid-LLC/roles-permissions-helper"

func main() {
    // without parsing "$stoId" etc.
    ownerRolePermissionsConfigWithoutParsing := role_permissions_config.GetOwnerRoleConfig()


    // with parsing "$stoId" etc.
    paramsToChange := make(map[string]string)
    paramsToChange["$stoId"] = "any_sto_internal_id_you_want"
    
    parsedData, errParsing := role_permissions_parser.GetOwnerRoleParsed(paramsToChange)
    if errParsing != nil {
        log.Printf("Error parsing role: %v", errParsing)
        return
    }
}
```
