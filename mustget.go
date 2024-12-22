package mustget

import (
	"log"
)

const PANIC_WITH_VALUES_FORMAT_ = "key \"%v\" not found in %T (%p) %v\n"
const PANIC_FORMAT = "key \"%v\" not found in %T (%p)\n"

var (
	// PanicWithValues will panic with the values of the map and the key
	PanicWithValues bool

	// DisablePanic indicates whether the application should disable panic behavior.
	// When set to true, the application will not panic in case of an missing key.
	DisablePanic bool
)

// MustGet retrieves the value associated with the specified key from a map.
// If the key is found, the corresponding value is returned.
// If the key is not found and PanicWithValues is true, a panic is triggered with additional information.
// If the key is not found and PanicWithValues is false, a panic is triggered without additional information.
func MustGet[K comparable](m any, k K) any {
	if v, ok := m.(map[K]any)[k]; ok {
		return v
	}

	if !DisablePanic {
		if PanicWithValues {
			log.Panicf(PANIC_WITH_VALUES_FORMAT_, k, m, &m, m)
		} else {
			log.Panicf(PANIC_FORMAT, k, m, &m)
		}
	}

	return nil
}
