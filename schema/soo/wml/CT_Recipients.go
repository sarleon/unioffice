// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/sarleon/unioffice"
)

type CT_Recipients struct {
	// Data About Single Data Source Record
	RecipientData []*CT_RecipientData
}

func NewCT_Recipients() *CT_Recipients {
	ret := &CT_Recipients{}
	return ret
}

func (m *CT_Recipients) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	serecipientData := xml.StartElement{Name: xml.Name{Local: "w:recipientData"}}
	for _, c := range m.RecipientData {
		e.EncodeElement(c, serecipientData)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Recipients) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Recipients:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "recipientData"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "recipientData"}:
				tmp := NewCT_RecipientData()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.RecipientData = append(m.RecipientData, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_Recipients %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Recipients
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Recipients and its children
func (m *CT_Recipients) Validate() error {
	return m.ValidateWithPath("CT_Recipients")
}

// ValidateWithPath validates the CT_Recipients and its children, prefixing error messages with path
func (m *CT_Recipients) ValidateWithPath(path string) error {
	for i, v := range m.RecipientData {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/RecipientData[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
