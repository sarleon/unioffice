// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"

	"github.com/sarleon/unioffice"
)

type EG_SlideListChoice struct {
	// All Slides
	SldAll *CT_Empty
	// Slide Range
	SldRg *CT_IndexRange
	// Custom Show
	CustShow *CT_CustomShowId
}

func NewEG_SlideListChoice() *EG_SlideListChoice {
	ret := &EG_SlideListChoice{}
	return ret
}

func (m *EG_SlideListChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.SldAll != nil {
		sesldAll := xml.StartElement{Name: xml.Name{Local: "p:sldAll"}}
		e.EncodeElement(m.SldAll, sesldAll)
	}
	if m.SldRg != nil {
		sesldRg := xml.StartElement{Name: xml.Name{Local: "p:sldRg"}}
		e.EncodeElement(m.SldRg, sesldRg)
	}
	if m.CustShow != nil {
		secustShow := xml.StartElement{Name: xml.Name{Local: "p:custShow"}}
		e.EncodeElement(m.CustShow, secustShow)
	}
	return nil
}

func (m *EG_SlideListChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_SlideListChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sldAll"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "sldAll"}:
				m.SldAll = NewCT_Empty()
				if err := d.DecodeElement(m.SldAll, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sldRg"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "sldRg"}:
				m.SldRg = NewCT_IndexRange()
				if err := d.DecodeElement(m.SldRg, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "custShow"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "custShow"}:
				m.CustShow = NewCT_CustomShowId()
				if err := d.DecodeElement(m.CustShow, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on EG_SlideListChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_SlideListChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_SlideListChoice and its children
func (m *EG_SlideListChoice) Validate() error {
	return m.ValidateWithPath("EG_SlideListChoice")
}

// ValidateWithPath validates the EG_SlideListChoice and its children, prefixing error messages with path
func (m *EG_SlideListChoice) ValidateWithPath(path string) error {
	if m.SldAll != nil {
		if err := m.SldAll.ValidateWithPath(path + "/SldAll"); err != nil {
			return err
		}
	}
	if m.SldRg != nil {
		if err := m.SldRg.ValidateWithPath(path + "/SldRg"); err != nil {
			return err
		}
	}
	if m.CustShow != nil {
		if err := m.CustShow.ValidateWithPath(path + "/CustShow"); err != nil {
			return err
		}
	}
	return nil
}
