// Copyright 2012 John Connor. All rights reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package builder

import "fmt"
import "reflect"
import "strings"

// Represents an attribute in the generated dot file, for example a node name
// or an edge color.
type attribute struct {
	Name  string
	Value string
}

type attrlist []*attribute

// TODO:  Is this "good" go?
// Intended to provide type safety.
type nodeattrs attrlist

// Intended to provide type safety.
type edgeattrs attrlist

// Intended to provide type safety.
type graphattrs attrlist

func (al attrlist) String(indent int, multiline bool) string {
	var idnt string = ""
	for i := 1; i < indent; i++ {
		idnt += "\t"
	}

	if len(al) > 0 {
		var str string
		atrs := make([]string, 0, len(al))
		if multiline {
			for _, a := range al {
				atrs = append(atrs, fmt.Sprintf("\t%s%s=\"%s\"", idnt, a.Name, a.Value))
			}
			str = strings.Join(atrs, "\n")
			str = fmt.Sprintf("\n%s\n%s", str, idnt)
		} else {
			for _, a := range al {
				atrs = append(atrs, fmt.Sprintf("%s%s=\"%s\"", idnt, a.Name, a.Value))
			}
			str = strings.Join(atrs, ", ")
			str = fmt.Sprintf("%s", str)
		}
		return str
	}
	return ""
}

// TODO: I can't figure out how to make this work with pointers.  Currently I
// just copy the structure being built into the obj paramater.
// Reflects on the type of "obj" to find tagged fields.  It then extracts the
// field information and encodes as a slice of pointers to "attribute" 
// structures.
// Used to extract information from Graph, Node and Edge objects.
func buildAttributes(obj interface{}) []*attribute {
	typ := reflect.TypeOf(obj)
	atrs := make([]*attribute, 0)

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if name := f.Tag.Get("name"); name != "" {
			def := f.Tag.Get("default")
			if fval := reflect.ValueOf(obj).FieldByName(f.Name); fval.IsValid() {
				str := getStr(name, fval)
				if atr := getAttr(name, str, def); atr != nil {
					atrs = append(atrs, atr)
				}
			}
		}
	}
	return atrs
}

func getStr(name string, val reflect.Value) string {
	if val.CanInterface() {
		if ifc := val.Interface(); ifc != nil {
			if v, ok := ifc.(string); ok {
				return v
			} else {
				if val.IsNil() {
					return ""
				}
				return fmt.Sprintf("%s", ifc)
			}
		}
	}
	return ""
}

func getAttr(key string, val string, def string) *attribute {
	if key == "" || (val == "" && def == "") {
		return nil
	}
	if val == "" {
		return &attribute{key, def}
	}
	return &attribute{key, val}
}
