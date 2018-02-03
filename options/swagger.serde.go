package options

import (
	"encoding/json"
)

type schemaAlias Schema

type SchemaJSON struct {
	schemaAlias
	Ref string `json:"$ref,omitempty"`
}

func (s Schema) MarshalJSON() ([]byte, error) {
	// the original Ref has omitempty, so setting it to "" removes it and allows
	// just the new one to be rendered
	ref := s.Ref
	s.Ref = ""
	return json.Marshal(SchemaJSON{schemaAlias(s), ref})
}
