# TOTP Package

This package provides a simple TOTP (Time-Based One-Time Password) generation and validation mechanism. It supports custom configuration and includes basic validation checks for secret keys.

## Features
- Validate secret keys using base32 encoding rules.
- Generate one-time passwords based on current or custom-supplied time.
- Error handling for invalid keys and decoding issues.

## Installation
Add to your Go module:
```bash
go get github.com/dylanmazurek/go-tools/pkg/totp
```

## Usage
```go
import (
    "fmt"
    "time"

    "github.com/dylanmazurek/go-tools/pkg/totp"
)

func main() {
    // Configure options
    t, err := totp.New(
        totp.WithSecretKey("BLD56GFS34BH67F6"),
        totp.WithTime(time.Now()),
    )
    if err != nil {
        panic(err)
    }

    // Generate TOTP code
    code, err := t.Generate()
    if err != nil {
        panic(err)
    }

    fmt.Println("TOTP code:", code)
}
```

For testing examples, see [totp_test.go](#file:totp_test.go-context).
