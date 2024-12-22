package mustget

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const JSON_TEST_STR = `{
	"$schema": "http://localhost:8080/schemas/User.json",
	"id": "c2b7db34-b202-4004-8f76-54ae9c7d604f",
	"created_at": "2024-09-26T12:36:38.802574782Z",
	"updated_at": "2024-09-26T12:36:38.802574782Z",
	"some_number": 1234567890,
	"_links": {
		"self": {
			"href": "/users/c2b7db34-b202-4004-8f76-54ae9c7d604f"
		}
	}
}`

// oldLogWriter is a variable that holds the previous value of the log writer.
// It is used to restore the original log writer after a test has completed.
var oldLogWriter io.Writer

// getTestMap unmarshals a JSON string defined by JSON_TEST_STR into a map with string keys and values of any type.
// It returns the resulting map. If the unmarshalling process encounters an error, the function will panic.
func getTestMap() map[string]any {
	testMap := make(map[string]any)

	if err := json.Unmarshal([]byte(JSON_TEST_STR), &testMap); err != nil {
		panic(err)
	}

	return testMap
}

// getTestMapString returns a string representation of the test map.
// It uses the fmt.Sprintf function to format the map returned by getTestMap.
func getTestMapString() string {
	return fmt.Sprintf("%v", getTestMap())
}

// setupTest configures the test environment by redirecting log output to io.Discard.
// This prevents log messages from cluttering the test output.
// It also saves the current log writer to oldLogWriter for later restoration.
func setupTest(_ *testing.T) {
	oldLogWriter = log.Writer()

	log.SetOutput(io.Discard)
}

// teardownTest restores the original log output by setting the log output
// back to the old log writer. This function is typically used in test
// teardown to ensure that any changes to the log output during tests do not
// affect other tests or the application.
func teardownTest(_ *testing.T) {
	log.SetOutput(oldLogWriter)
}

// TestPanic tests the MustGet function to ensure it panics when a key is not found in the map.
// It sets up the test environment, defers the teardown, and uses a deferred function to recover from the panic.
// The test checks if the recovered error message contains the expected substring and suffix.
func TestPanic(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	defer func() {
		var err any

		if err = recover(); err == nil {
			t.Errorf("The code did not panic")
		}

		errString := err.(string)

		assert.True(t, strings.Contains(errString, "key \"some_number1\" not found in map[string]interface {} ("))
		assert.True(t, strings.HasSuffix(errString, ")\n"))
	}()

	testMap := getTestMap()

	_ = MustGet(testMap, "some_number1").(float64)
}

// TestPanicWithValues tests the behavior of the MustGet function when a key is not found in the map.
// It sets up the test environment, defers the teardown, and uses a deferred function to recover from a panic.
// The test expects a panic to occur and verifies that the panic message contains specific substrings.
// The PanicWithValues flag is set to true to enable panic on missing keys.
// The test map is retrieved and the MustGet function is called with a key that does not exist in the map.
func TestPanicWithValues(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	defer func() {
		var err any

		if err = recover(); err == nil {
			t.Errorf("The code did not panic")
		}

		errString := err.(string)

		assert.True(t, strings.Contains(errString, "key \"some_number1\" not found in map[string]interface {} ("))
		assert.True(t, strings.Contains(errString, getTestMapString()))
	}()

	PanicWithValues = true

	testMap := getTestMap()

	_ = MustGet(testMap, "some_number1").(float64)
}

// TestDisabledPanic tests the behavior of the MustGet function when the DisablePanic flag is set to true.
// It ensures that MustGet returns nil instead of panicking when the key is not found in the map.
func TestDisabledPanic(t *testing.T) {
	setupTest(t)
	defer teardownTest(t)

	DisablePanic = true

	testMap := getTestMap()

	assert.Nil(t, MustGet(testMap, "some_number1"))
}
