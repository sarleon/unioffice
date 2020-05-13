// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/sarleon/unioffice"
)

type CT_CommentPr struct {
	// Locked Flag
	LockedAttr *bool
	// Default Size Flag
	DefaultSizeAttr *bool
	// Print Flag
	PrintAttr *bool
	// Disabled Flag
	DisabledAttr *bool
	// Automatic Fill Flag
	AutoFillAttr *bool
	// Automatic Line Flag
	AutoLineAttr *bool
	// Alternative Text
	AltTextAttr *string
	// Text Horizontal Alignment
	TextHAlignAttr ST_TextHAlign
	// ext Vertical Alignment
	TextVAlignAttr ST_TextVAlign
	// Text Lock Flag
	LockTextAttr *bool
	// Far East Alignment Flag
	JustLastXAttr *bool
	// Automatic Text Scaling Flag
	AutoScaleAttr *bool
	Anchor        *CT_ObjectAnchor
}

func NewCT_CommentPr() *CT_CommentPr {
	ret := &CT_CommentPr{}
	ret.Anchor = NewCT_ObjectAnchor()
	return ret
}

func (m *CT_CommentPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LockedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "locked"},
			Value: fmt.Sprintf("%d", b2i(*m.LockedAttr))})
	}
	if m.DefaultSizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defaultSize"},
			Value: fmt.Sprintf("%d", b2i(*m.DefaultSizeAttr))})
	}
	if m.PrintAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "print"},
			Value: fmt.Sprintf("%d", b2i(*m.PrintAttr))})
	}
	if m.DisabledAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "disabled"},
			Value: fmt.Sprintf("%d", b2i(*m.DisabledAttr))})
	}
	if m.AutoFillAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoFill"},
			Value: fmt.Sprintf("%d", b2i(*m.AutoFillAttr))})
	}
	if m.AutoLineAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoLine"},
			Value: fmt.Sprintf("%d", b2i(*m.AutoLineAttr))})
	}
	if m.AltTextAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "altText"},
			Value: fmt.Sprintf("%v", *m.AltTextAttr)})
	}
	if m.TextHAlignAttr != ST_TextHAlignUnset {
		attr, err := m.TextHAlignAttr.MarshalXMLAttr(xml.Name{Local: "textHAlign"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TextVAlignAttr != ST_TextVAlignUnset {
		attr, err := m.TextVAlignAttr.MarshalXMLAttr(xml.Name{Local: "textVAlign"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.LockTextAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lockText"},
			Value: fmt.Sprintf("%d", b2i(*m.LockTextAttr))})
	}
	if m.JustLastXAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "justLastX"},
			Value: fmt.Sprintf("%d", b2i(*m.JustLastXAttr))})
	}
	if m.AutoScaleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoScale"},
			Value: fmt.Sprintf("%d", b2i(*m.AutoScaleAttr))})
	}
	e.EncodeToken(start)
	seanchor := xml.StartElement{Name: xml.Name{Local: "ma:anchor"}}
	e.EncodeElement(m.Anchor, seanchor)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CommentPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Anchor = NewCT_ObjectAnchor()
	for _, attr := range start.Attr {
		if attr.Name.Local == "altText" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AltTextAttr = &parsed
			continue
		}
		if attr.Name.Local == "defaultSize" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DefaultSizeAttr = &parsed
			continue
		}
		if attr.Name.Local == "print" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PrintAttr = &parsed
			continue
		}
		if attr.Name.Local == "disabled" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DisabledAttr = &parsed
			continue
		}
		if attr.Name.Local == "autoFill" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoFillAttr = &parsed
			continue
		}
		if attr.Name.Local == "autoLine" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoLineAttr = &parsed
			continue
		}
		if attr.Name.Local == "locked" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LockedAttr = &parsed
			continue
		}
		if attr.Name.Local == "textHAlign" {
			m.TextHAlignAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "textVAlign" {
			m.TextVAlignAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "lockText" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LockTextAttr = &parsed
			continue
		}
		if attr.Name.Local == "justLastX" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.JustLastXAttr = &parsed
			continue
		}
		if attr.Name.Local == "autoScale" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoScaleAttr = &parsed
			continue
		}
	}
lCT_CommentPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "anchor"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "anchor"}:
				if err := d.DecodeElement(m.Anchor, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_CommentPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CommentPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CommentPr and its children
func (m *CT_CommentPr) Validate() error {
	return m.ValidateWithPath("CT_CommentPr")
}

// ValidateWithPath validates the CT_CommentPr and its children, prefixing error messages with path
func (m *CT_CommentPr) ValidateWithPath(path string) error {
	if err := m.TextHAlignAttr.ValidateWithPath(path + "/TextHAlignAttr"); err != nil {
		return err
	}
	if err := m.TextVAlignAttr.ValidateWithPath(path + "/TextVAlignAttr"); err != nil {
		return err
	}
	if err := m.Anchor.ValidateWithPath(path + "/Anchor"); err != nil {
		return err
	}
	return nil
}
