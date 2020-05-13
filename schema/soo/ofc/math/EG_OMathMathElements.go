// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package math

import (
	"encoding/xml"

	"github.com/sarleon/unioffice"
)

type EG_OMathMathElements struct {
	Acc       *CT_Acc
	Bar       *CT_Bar
	Box       *CT_Box
	BorderBox *CT_BorderBox
	D         *CT_D
	EqArr     *CT_EqArr
	F         *CT_F
	Func      *CT_Func
	GroupChr  *CT_GroupChr
	LimLow    *CT_LimLow
	LimUpp    *CT_LimUpp
	M         *CT_M
	Nary      *CT_Nary
	Phant     *CT_Phant
	Rad       *CT_Rad
	SPre      *CT_SPre
	SSub      *CT_SSub
	SSubSup   *CT_SSubSup
	SSup      *CT_SSup
	R         *CT_R
}

func NewEG_OMathMathElements() *EG_OMathMathElements {
	ret := &EG_OMathMathElements{}
	return ret
}

func (m *EG_OMathMathElements) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Acc != nil {
		seacc := xml.StartElement{Name: xml.Name{Local: "m:acc"}}
		e.EncodeElement(m.Acc, seacc)
	}
	if m.Bar != nil {
		sebar := xml.StartElement{Name: xml.Name{Local: "m:bar"}}
		e.EncodeElement(m.Bar, sebar)
	}
	if m.Box != nil {
		sebox := xml.StartElement{Name: xml.Name{Local: "m:box"}}
		e.EncodeElement(m.Box, sebox)
	}
	if m.BorderBox != nil {
		seborderBox := xml.StartElement{Name: xml.Name{Local: "m:borderBox"}}
		e.EncodeElement(m.BorderBox, seborderBox)
	}
	if m.D != nil {
		sed := xml.StartElement{Name: xml.Name{Local: "m:d"}}
		e.EncodeElement(m.D, sed)
	}
	if m.EqArr != nil {
		seeqArr := xml.StartElement{Name: xml.Name{Local: "m:eqArr"}}
		e.EncodeElement(m.EqArr, seeqArr)
	}
	if m.F != nil {
		sef := xml.StartElement{Name: xml.Name{Local: "m:f"}}
		e.EncodeElement(m.F, sef)
	}
	if m.Func != nil {
		sefunc := xml.StartElement{Name: xml.Name{Local: "m:func"}}
		e.EncodeElement(m.Func, sefunc)
	}
	if m.GroupChr != nil {
		segroupChr := xml.StartElement{Name: xml.Name{Local: "m:groupChr"}}
		e.EncodeElement(m.GroupChr, segroupChr)
	}
	if m.LimLow != nil {
		selimLow := xml.StartElement{Name: xml.Name{Local: "m:limLow"}}
		e.EncodeElement(m.LimLow, selimLow)
	}
	if m.LimUpp != nil {
		selimUpp := xml.StartElement{Name: xml.Name{Local: "m:limUpp"}}
		e.EncodeElement(m.LimUpp, selimUpp)
	}
	if m.M != nil {
		sem := xml.StartElement{Name: xml.Name{Local: "m:m"}}
		e.EncodeElement(m.M, sem)
	}
	if m.Nary != nil {
		senary := xml.StartElement{Name: xml.Name{Local: "m:nary"}}
		e.EncodeElement(m.Nary, senary)
	}
	if m.Phant != nil {
		sephant := xml.StartElement{Name: xml.Name{Local: "m:phant"}}
		e.EncodeElement(m.Phant, sephant)
	}
	if m.Rad != nil {
		serad := xml.StartElement{Name: xml.Name{Local: "m:rad"}}
		e.EncodeElement(m.Rad, serad)
	}
	if m.SPre != nil {
		sesPre := xml.StartElement{Name: xml.Name{Local: "m:sPre"}}
		e.EncodeElement(m.SPre, sesPre)
	}
	if m.SSub != nil {
		sesSub := xml.StartElement{Name: xml.Name{Local: "m:sSub"}}
		e.EncodeElement(m.SSub, sesSub)
	}
	if m.SSubSup != nil {
		sesSubSup := xml.StartElement{Name: xml.Name{Local: "m:sSubSup"}}
		e.EncodeElement(m.SSubSup, sesSubSup)
	}
	if m.SSup != nil {
		sesSup := xml.StartElement{Name: xml.Name{Local: "m:sSup"}}
		e.EncodeElement(m.SSup, sesSup)
	}
	if m.R != nil {
		ser := xml.StartElement{Name: xml.Name{Local: "m:r"}}
		e.EncodeElement(m.R, ser)
	}
	return nil
}

func (m *EG_OMathMathElements) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_OMathMathElements:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "acc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "acc"}:
				m.Acc = NewCT_Acc()
				if err := d.DecodeElement(m.Acc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "bar"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "bar"}:
				m.Bar = NewCT_Bar()
				if err := d.DecodeElement(m.Bar, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "box"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "box"}:
				m.Box = NewCT_Box()
				if err := d.DecodeElement(m.Box, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "borderBox"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "borderBox"}:
				m.BorderBox = NewCT_BorderBox()
				if err := d.DecodeElement(m.BorderBox, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "d"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "d"}:
				m.D = NewCT_D()
				if err := d.DecodeElement(m.D, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "eqArr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "eqArr"}:
				m.EqArr = NewCT_EqArr()
				if err := d.DecodeElement(m.EqArr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "f"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "f"}:
				m.F = NewCT_F()
				if err := d.DecodeElement(m.F, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "func"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "func"}:
				m.Func = NewCT_Func()
				if err := d.DecodeElement(m.Func, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "groupChr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "groupChr"}:
				m.GroupChr = NewCT_GroupChr()
				if err := d.DecodeElement(m.GroupChr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "limLow"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "limLow"}:
				m.LimLow = NewCT_LimLow()
				if err := d.DecodeElement(m.LimLow, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "limUpp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "limUpp"}:
				m.LimUpp = NewCT_LimUpp()
				if err := d.DecodeElement(m.LimUpp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "m"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "m"}:
				m.M = NewCT_M()
				if err := d.DecodeElement(m.M, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "nary"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "nary"}:
				m.Nary = NewCT_Nary()
				if err := d.DecodeElement(m.Nary, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "phant"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "phant"}:
				m.Phant = NewCT_Phant()
				if err := d.DecodeElement(m.Phant, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "rad"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "rad"}:
				m.Rad = NewCT_Rad()
				if err := d.DecodeElement(m.Rad, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sPre"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "sPre"}:
				m.SPre = NewCT_SPre()
				if err := d.DecodeElement(m.SPre, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sSub"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "sSub"}:
				m.SSub = NewCT_SSub()
				if err := d.DecodeElement(m.SSub, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sSubSup"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "sSubSup"}:
				m.SSubSup = NewCT_SSubSup()
				if err := d.DecodeElement(m.SSubSup, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "sSup"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "sSup"}:
				m.SSup = NewCT_SSup()
				if err := d.DecodeElement(m.SSup, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "r"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "r"}:
				m.R = NewCT_R()
				if err := d.DecodeElement(m.R, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on EG_OMathMathElements %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_OMathMathElements
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_OMathMathElements and its children
func (m *EG_OMathMathElements) Validate() error {
	return m.ValidateWithPath("EG_OMathMathElements")
}

// ValidateWithPath validates the EG_OMathMathElements and its children, prefixing error messages with path
func (m *EG_OMathMathElements) ValidateWithPath(path string) error {
	if m.Acc != nil {
		if err := m.Acc.ValidateWithPath(path + "/Acc"); err != nil {
			return err
		}
	}
	if m.Bar != nil {
		if err := m.Bar.ValidateWithPath(path + "/Bar"); err != nil {
			return err
		}
	}
	if m.Box != nil {
		if err := m.Box.ValidateWithPath(path + "/Box"); err != nil {
			return err
		}
	}
	if m.BorderBox != nil {
		if err := m.BorderBox.ValidateWithPath(path + "/BorderBox"); err != nil {
			return err
		}
	}
	if m.D != nil {
		if err := m.D.ValidateWithPath(path + "/D"); err != nil {
			return err
		}
	}
	if m.EqArr != nil {
		if err := m.EqArr.ValidateWithPath(path + "/EqArr"); err != nil {
			return err
		}
	}
	if m.F != nil {
		if err := m.F.ValidateWithPath(path + "/F"); err != nil {
			return err
		}
	}
	if m.Func != nil {
		if err := m.Func.ValidateWithPath(path + "/Func"); err != nil {
			return err
		}
	}
	if m.GroupChr != nil {
		if err := m.GroupChr.ValidateWithPath(path + "/GroupChr"); err != nil {
			return err
		}
	}
	if m.LimLow != nil {
		if err := m.LimLow.ValidateWithPath(path + "/LimLow"); err != nil {
			return err
		}
	}
	if m.LimUpp != nil {
		if err := m.LimUpp.ValidateWithPath(path + "/LimUpp"); err != nil {
			return err
		}
	}
	if m.M != nil {
		if err := m.M.ValidateWithPath(path + "/M"); err != nil {
			return err
		}
	}
	if m.Nary != nil {
		if err := m.Nary.ValidateWithPath(path + "/Nary"); err != nil {
			return err
		}
	}
	if m.Phant != nil {
		if err := m.Phant.ValidateWithPath(path + "/Phant"); err != nil {
			return err
		}
	}
	if m.Rad != nil {
		if err := m.Rad.ValidateWithPath(path + "/Rad"); err != nil {
			return err
		}
	}
	if m.SPre != nil {
		if err := m.SPre.ValidateWithPath(path + "/SPre"); err != nil {
			return err
		}
	}
	if m.SSub != nil {
		if err := m.SSub.ValidateWithPath(path + "/SSub"); err != nil {
			return err
		}
	}
	if m.SSubSup != nil {
		if err := m.SSubSup.ValidateWithPath(path + "/SSubSup"); err != nil {
			return err
		}
	}
	if m.SSup != nil {
		if err := m.SSup.ValidateWithPath(path + "/SSup"); err != nil {
			return err
		}
	}
	if m.R != nil {
		if err := m.R.ValidateWithPath(path + "/R"); err != nil {
			return err
		}
	}
	return nil
}
