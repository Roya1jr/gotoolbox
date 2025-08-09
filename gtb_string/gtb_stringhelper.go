package gtb_string

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"text/template"
)

// AddPre adds prefix to suffix
func AddPre(suffix, prefix string) string {
	return prefix + suffix
}

// CreateXML takes an XML template string and v data structure to produce the final XML.
func CreateXml(template_str string, v any) (bytes.Buffer, error) {
	var buf bytes.Buffer

	xml_body, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return buf, fmt.Errorf("could not marshal data to XML: %w", err)
	}

	tmpl, err := template.New("xmlTemplate").Parse(template_str)
	if err != nil {
		return buf, fmt.Errorf("could not parse template: %w", err)
	}

	if err := tmpl.Execute(&buf, string(xml_body)); err != nil {
		return buf, fmt.Errorf("could not execute template: %w", err)
	}

	return buf, nil
}
