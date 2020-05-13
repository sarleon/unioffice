// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/sarleon/unioffice/schema/soo/dml/chartDrawing"
)

func TestCT_PictureNonVisualConstructor(t *testing.T) {
	v := chartDrawing.NewCT_PictureNonVisual()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_PictureNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_PictureNonVisual should validate: %s", err)
	}
}

func TestCT_PictureNonVisualMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_PictureNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_PictureNonVisual()
	xml.Unmarshal(buf, v2)
}
