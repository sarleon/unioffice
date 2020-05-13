// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/sarleon/unioffice/schema/soo/dml"
)

func TestEG_LineJoinPropertiesConstructor(t *testing.T) {
	v := dml.NewEG_LineJoinProperties()
	if v == nil {
		t.Errorf("dml.NewEG_LineJoinProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_LineJoinProperties should validate: %s", err)
	}
}

func TestEG_LineJoinPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_LineJoinProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_LineJoinProperties()
	xml.Unmarshal(buf, v2)
}
