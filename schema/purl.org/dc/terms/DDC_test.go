// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package terms_test

import (
	"encoding/xml"
	"testing"

	"github.com/sarleon/unioffice/schema/purl.org/dc/terms"
)

func TestDDCConstructor(t *testing.T) {
	v := terms.NewDDC()
	if v == nil {
		t.Errorf("terms.NewDDC must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.DDC should validate: %s", err)
	}
}

func TestDDCMarshalUnmarshal(t *testing.T) {
	v := terms.NewDDC()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewDDC()
	xml.Unmarshal(buf, v2)
}
