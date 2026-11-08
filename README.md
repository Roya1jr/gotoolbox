# GoToolbox

A collection of useful Go utilities, tools, and string constants for common programming tasks.

## Features

- **HTTP Headers**: Pre-defined HTTP header constants for authorization, content types, cache control, etc.
- **String Utilities**: Case-insensitive comparison functions and utility helpers.
- **DB Helpers**: Database-related helper functions (placeholder package).

## Installation

```bash
go get github.com/Roya1jr/gotoolbox
```

## Usage Examples

### HTTP Headers

```go
package main

import (
    "log"
    gtbhttp "github.com/Roya1jr/gotoolbox/http"
)

func main() {
    // Authorization header
    log.Println(gtbhttp.HeaderAuthorization)
    
    // Content-Type headers
    log.Println(gtbhttp.ContentTypeJSON)
    log.Println(gtbhttp.ContentTypeHTML)
    log.Println(gtbhttp.ContentTypeXML)
    
    // Cache control headers
    log.Println(gtbhttp.CacheControlNoCache)
    log.Println(gtbhttp.CacheControlMustRevalidate)
    log.Println(gtbhttp.CacheControlMaxAge("300")) 
    
    // CORS headers
    log.Println(gtbhttp.HeaderAccessControlAllowOrigin)
    
    // Content security policy
    log.Println(gtbhttp.ContentSecurityPolicyBaseURI)
}
```

### String Utilities

```go
package main

import (
    "strings"
    gtbstring "github.com/Roya1jr/gotoolbox/string"
)

func main() {
    // Case-insensitive string comparison
    result := gtbstring.Casefold("Hello", "hello")
    log.Println(result) // true
    

}
```

## Package Structure

```
gotoolbox/
├── db/              # Database helpers (placeholder)
├── http/            # HTTP headers and utilities
│   ├── gtb_httpheaders.go
│   └── gtb_httpheaders_test.go
├── string/          # String utility functions
│   └── gtb_stringcasefold.go
└── go.mod           # Go module definition
```

## Development

### Run Tests

```bash
go test -v ./...
```

### Lint and Check

```bash
golangci-lint run
```

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
