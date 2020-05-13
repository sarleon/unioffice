// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/sarleon/unioffice/schema/soo/sml"
)

func TestCT_PivotSelectionConstructor(t *testing.T) {
	v := sml.NewCT_PivotSelection()
	if v == nil {
		t.Errorf("sml.NewCT_PivotSelection must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotSelection should validate: %s", err)
	}
}

func TestCT_PivotSelectionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotSelection()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotSelection()
	xml.Unmarshal(buf, v2)
}
