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

import "io"
import "fmt"

import "godot/attr"

type dotgraph struct {
	kind       *attr.GraphKind
	nodes      []*dotnode
	edges      []*dotedge
	nTmpl      nodeattrs
	eTmpl      edgeattrs
	attributes graphattrs
}

func (g dotgraph) Write(writer io.Writer) error {
	if _, err := fmt.Fprintf(writer, "%s {\n", g.kind.Name()); err != nil {
		return err
	}

	if len(g.attributes) > 0 {
		str := attrlist(g.attributes).String(1, true)
		if _, err := fmt.Fprintf(writer, "%s\n", str); err != nil {
			return err
		}
	}

	if g.nTmpl != nil {
		str := attrlist(g.nTmpl).String(2, true)
		if _, err := fmt.Fprintf(writer, "\tnode [%s]\n\n", str); err != nil {
			return err
		}
	}

	if g.eTmpl != nil {
		str := attrlist(g.eTmpl).String(2, true)
		if _, err := fmt.Fprintf(writer, "\tedge [%s]\n\n", str); err != nil {
			return err
		}
	}

	for _, node := range g.nodes {
		if err := node.Write(writer); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintln(writer, ""); err != nil {
		return err
	}

	for _, edge := range g.edges {
		if err := edge.Write(writer); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(writer, "}\n"); err != nil {
		return err
	}

	return nil
}
