// Package gtbstring contains string helper functions
package gtbstring

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"text/template"
)

// CreateXML takes an XML template string and v data structure to produce the final XML.
func CreateXML(templateStr string, v any) (bytes.Buffer, error) {
	var buf bytes.Buffer

	xmlBody, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return buf, fmt.Errorf("could not marshal data to XML: %w", err)
	}

	tmpl, err := template.New("xmlTemplate").Parse(templateStr)
	if err != nil {
		return buf, fmt.Errorf("could not parse template: %w", err)
	}

	if err := tmpl.Execute(&buf, string(xmlBody)); err != nil {
		return buf, fmt.Errorf("could not execute template: %w", err)
	}

	return buf, nil
}
