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

type opAlias Operation

type OperationJSON struct {
	opAlias
	RequestBody *RequestBody `json:"requestBody,omitempty"`
}

func (o Operation) MarshalJSON() ([]byte, error) {
	// the original RequestBody has omitempty, so setting it to "" removes it and allows
	// just the new one to be rendered
	rb := o.RequestBody
	o.RequestBody = nil
	return json.Marshal(OperationJSON{opAlias(o), rb})
}
