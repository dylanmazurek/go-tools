# Truncate Tool

A simple utility for shortening strings to a specified length. Useful for logs, summaries, or any scenario requiring controlled string length.

## Installation
```bash
go get github.com/dylanmazurek/go-tools/pkg/truncate
```

## Usage
```go
import "github.com/dylanmazurek/go-tools/pkg/truncate"

func main() {
    result := truncate.String("Example text", 5)
    // result: "Examp"
}
```

## Features
- Efficient and straightforward implementation
- Customizable maximum length

## License
Distributed under the MIT License.