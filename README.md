# Go-MustGet

`Go-MustGet` is a Go package that provides a utility function to retrieve values from a map with panic behavior when keys are not found.

## Installation

To install the package, use `go get`:

```sh
go get github.com/skazanyNaGlany/go-mustget
```

## Usage

```go
package main

import (
	"log"

	"github.com/skazanyNaGlany/go-mustget"
)

func main() {
	someMap := map[string]any{
		"key": "some value",
	}

	value := mustget.MustGet(someMap, "key").(string)

	log.Println(value)
}
```

### MustGet Function

The `MustGet` function retrieves the value associated with the specified key from a map. If the key is found, the corresponding value is returned. If the key is not found, the code will panic, this behavior can be configured with `PanicWithValues` and `DisablePanic` flags.

```go
func MustGet[K comparable](m any, k K) any
```

#### Parameters

- `m`: The map from which to retrieve the value.
- `k`: The key whose associated value is to be returned.

#### Return Value

- The value associated with the specified key, or `nil` if the key is not found and `DisablePanic` is true.

### Flags

- `PanicWithValues`: When set to `true`, `MustGet` will panic with additional information about the map and the key if the key is not found.
- `DisablePanic`: When set to `true`, `MustGet` will return `nil` instead of panicking if the key is not found.

## Testing

The package includes tests to verify the behavior of the `MustGet` function. To run the tests, use the `go test` command:

```sh
go test
```

## License

© Paweł Kacperski, 2024 ~ time.Now

Released under the [MIT License](https://github.com/go-gorm/gorm/blob/master/LICENSE)
