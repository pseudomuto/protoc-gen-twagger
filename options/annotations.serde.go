package options

import (
	"encoding/json"
	"fmt"
	"strings"
)

type componentsAlias Components

// ComponentsJSON describes a custom serialization object for Components
type ComponentsJSON struct {
	componentsAlias
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes,omitempty"`
}

// MarshalJSON marshals the Components object without SecuritySchemas
func (c Components) MarshalJSON() ([]byte, error) {
	// the original SecuritySchemes has omitempty, so setting it to nil removes it and allows
	// just the new one to be rendered
	ss := c.SecuritySchemes
	c.SecuritySchemes = nil
	return json.Marshal(ComponentsJSON{componentsAlias(c), ss})
}

type opAlias Operation

// OperationJSON describes a custom serialization object for Operation
type OperationJSON struct {
	opAlias
	RequestBody *RequestBody `json:"requestBody,omitempty"`
}

// MarshalJSON marshals the Operation object without the RequestBody
func (o Operation) MarshalJSON() ([]byte, error) {
	// the original RequestBody has omitempty, so setting it to nil removes it and allows
	// just the new one to be rendered
	rb := o.RequestBody
	o.RequestBody = nil
	return json.Marshal(OperationJSON{opAlias(o), rb})
}

type schemaAlias Schema

// SchemaJSON describes a custom serialization object for Schema
type SchemaJSON struct {
	schemaAlias
	Ref string `json:"$ref,omitempty"`
}

// MarshalJSON marshals the Schema without the Ref
func (s Schema) MarshalJSON() ([]byte, error) {
	// the original Ref has omitempty, so setting it to "" removes it and allows
	// just the new one to be rendered
	ref := s.Ref
	s.Ref = ""
	return json.Marshal(SchemaJSON{schemaAlias(s), ref})
}

// MarshalJSON marshals the name of the SecurityRequirement with an empty set of scopes
func (r SecurityRequirement) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{ "%s": [] }`, r.GetName())), nil
}

type secSchemeAlias SecurityScheme

// SecuritySchemeJSON describes a custom serialization object for SecurityScheme
type SecuritySchemeJSON struct {
	secSchemeAlias
	Type             string `json:"type,omitempty"`
	BearerFormat     string `json:"bearerFormat,omitempty"`
	OpenIdConnectUrl string `json:"openIdConnectUrl,omitempty"`
}

// MarshalJSON marshals the SecurityScheme
func (s SecurityScheme) MarshalJSON() ([]byte, error) {
	// the original props have omitempty, so setting them to zero value removes them and allows
	// just the new ones to be rendered
	bf := s.BearerFormat
	oid := s.OpenIdConnectUrl
	s.BearerFormat = ""
	s.OpenIdConnectUrl = ""
	return json.Marshal(SecuritySchemeJSON{secSchemeAlias(s), strings.ToLower(s.GetType().String()), bf, oid})
}
