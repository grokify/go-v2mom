// Package schema provides the embedded JSON Schema for V2MOM validation.
package schema

import (
	_ "embed"
)

// SchemaVersion is the version of the V2MOM JSON Schema.
const SchemaVersion = "1.0.0"

// SchemaID is the canonical ID for the V2MOM JSON Schema.
const SchemaID = "https://github.com/grokify/structured-goals/schema/v2mom.schema.json"

//go:embed v2mom.schema.json
var schemaJSON []byte

// JSON returns the embedded V2MOM JSON Schema as bytes.
func JSON() []byte {
	return schemaJSON
}

// JSONString returns the embedded V2MOM JSON Schema as a string.
func JSONString() string {
	return string(schemaJSON)
}
