// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml

import (
	"encoding/xml"

	"github.com/sarleon/unioffice"
)

type CT_GvmlGroupShapeNonVisual struct {
	CNvPr      *CT_NonVisualDrawingProps
	CNvGrpSpPr *CT_NonVisualGroupDrawingShapeProps
}

func NewCT_GvmlGroupShapeNonVisual() *CT_GvmlGroupShapeNonVisual {
	ret := &CT_GvmlGroupShapeNonVisual{}
	ret.CNvPr = NewCT_NonVisualDrawingProps()
	ret.CNvGrpSpPr = NewCT_NonVisualGroupDrawingShapeProps()
	return ret
}

func (m *CT_GvmlGroupShapeNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "a:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvGrpSpPr := xml.StartElement{Name: xml.Name{Local: "a:cNvGrpSpPr"}}
	e.EncodeElement(m.CNvGrpSpPr, secNvGrpSpPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GvmlGroupShapeNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = NewCT_NonVisualDrawingProps()
	m.CNvGrpSpPr = NewCT_NonVisualGroupDrawingShapeProps()
lCT_GvmlGroupShapeNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cNvPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cNvPr"}:
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cNvGrpSpPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cNvGrpSpPr"}:
				if err := d.DecodeElement(m.CNvGrpSpPr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_GvmlGroupShapeNonVisual %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GvmlGroupShapeNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GvmlGroupShapeNonVisual and its children
func (m *CT_GvmlGroupShapeNonVisual) Validate() error {
	return m.ValidateWithPath("CT_GvmlGroupShapeNonVisual")
}

// ValidateWithPath validates the CT_GvmlGroupShapeNonVisual and its children, prefixing error messages with path
func (m *CT_GvmlGroupShapeNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvGrpSpPr.ValidateWithPath(path + "/CNvGrpSpPr"); err != nil {
		return err
	}
	return nil
}
