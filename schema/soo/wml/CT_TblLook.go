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

	"github.com/sarleon/unioffice/schema/soo/ofc/sharedTypes"
)

type CT_TblLook struct {
	// First Row
	FirstRowAttr *sharedTypes.ST_OnOff
	// Last Row
	LastRowAttr *sharedTypes.ST_OnOff
	// First Column
	FirstColumnAttr *sharedTypes.ST_OnOff
	// Last Column
	LastColumnAttr *sharedTypes.ST_OnOff
	// No Horizontal Banding
	NoHBandAttr *sharedTypes.ST_OnOff
	// No Vertical Banding
	NoVBandAttr *sharedTypes.ST_OnOff
	// Bitmask of Table Conditional Formatting
	ValAttr *string
}

func NewCT_TblLook() *CT_TblLook {
	ret := &CT_TblLook{}
	return ret
}

func (m *CT_TblLook) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FirstRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:firstRow"},
			Value: fmt.Sprintf("%v", *m.FirstRowAttr)})
	}
	if m.LastRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lastRow"},
			Value: fmt.Sprintf("%v", *m.LastRowAttr)})
	}
	if m.FirstColumnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:firstColumn"},
			Value: fmt.Sprintf("%v", *m.FirstColumnAttr)})
	}
	if m.LastColumnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:lastColumn"},
			Value: fmt.Sprintf("%v", *m.LastColumnAttr)})
	}
	if m.NoHBandAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:noHBand"},
			Value: fmt.Sprintf("%v", *m.NoHBandAttr)})
	}
	if m.NoVBandAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:noVBand"},
			Value: fmt.Sprintf("%v", *m.NoVBandAttr)})
	}
	if m.ValAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:val"},
			Value: fmt.Sprintf("%v", *m.ValAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TblLook) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "firstRow" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.FirstRowAttr = &parsed
			continue
		}
		if attr.Name.Local == "lastRow" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.LastRowAttr = &parsed
			continue
		}
		if attr.Name.Local == "firstColumn" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.FirstColumnAttr = &parsed
			continue
		}
		if attr.Name.Local == "lastColumn" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.LastColumnAttr = &parsed
			continue
		}
		if attr.Name.Local == "noHBand" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.NoHBandAttr = &parsed
			continue
		}
		if attr.Name.Local == "noVBand" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.NoVBandAttr = &parsed
			continue
		}
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TblLook: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TblLook and its children
func (m *CT_TblLook) Validate() error {
	return m.ValidateWithPath("CT_TblLook")
}

// ValidateWithPath validates the CT_TblLook and its children, prefixing error messages with path
func (m *CT_TblLook) ValidateWithPath(path string) error {
	if m.FirstRowAttr != nil {
		if err := m.FirstRowAttr.ValidateWithPath(path + "/FirstRowAttr"); err != nil {
			return err
		}
	}
	if m.LastRowAttr != nil {
		if err := m.LastRowAttr.ValidateWithPath(path + "/LastRowAttr"); err != nil {
			return err
		}
	}
	if m.FirstColumnAttr != nil {
		if err := m.FirstColumnAttr.ValidateWithPath(path + "/FirstColumnAttr"); err != nil {
			return err
		}
	}
	if m.LastColumnAttr != nil {
		if err := m.LastColumnAttr.ValidateWithPath(path + "/LastColumnAttr"); err != nil {
			return err
		}
	}
	if m.NoHBandAttr != nil {
		if err := m.NoHBandAttr.ValidateWithPath(path + "/NoHBandAttr"); err != nil {
			return err
		}
	}
	if m.NoVBandAttr != nil {
		if err := m.NoVBandAttr.ValidateWithPath(path + "/NoVBandAttr"); err != nil {
			return err
		}
	}
	return nil
}
