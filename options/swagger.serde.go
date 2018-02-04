package options

import (
	"encoding/json"
	"fmt"
	"strings"
)

type componentsAlias Components

type ComponentsJSON struct {
	componentsAlias
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes,omitempty"`
}

func (c Components) MarshalJSON() ([]byte, error) {
	// the original SecuritySchemes has omitempty, so setting it to nil removes it and allows
	// just the new one to be rendered
	ss := c.SecuritySchemes
	c.SecuritySchemes = nil
	return json.Marshal(ComponentsJSON{componentsAlias(c), ss})
}

type opAlias Operation

type OperationJSON struct {
	opAlias
	RequestBody *RequestBody `json:"requestBody,omitempty"`
}

func (o Operation) MarshalJSON() ([]byte, error) {
	// the original RequestBody has omitempty, so setting it to nil removes it and allows
	// just the new one to be rendered
	rb := o.RequestBody
	o.RequestBody = nil
	return json.Marshal(OperationJSON{opAlias(o), rb})
}

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

func (r SecurityRequirement) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{ "%s": [] }`, r.GetName())), nil
}

type secSchemeAlias SecurityScheme

type SecuritySchemeJSON struct {
	secSchemeAlias
	Type             string `json:"type,omitempty"`
	BearerFormat     string `json:"bearerFormat,omitempty"`
	OpenIdConnectUrl string `json:"openIdConnectUrl,omitempty"`
}

func (s SecurityScheme) MarshalJSON() ([]byte, error) {
	// the original props have omitempty, so setting them to zero value removes them and allows
	// just the new ones to be rendered
	bf := s.BearerFormat
	oid := s.OpenIdConnectUrl
	s.BearerFormat = ""
	s.OpenIdConnectUrl = ""
	return json.Marshal(SecuritySchemeJSON{secSchemeAlias(s), strings.ToLower(s.GetType().String()), bf, oid})
}
