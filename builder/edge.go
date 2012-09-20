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

import "godot/attr/color"

// A builder for an edge.
//
// References:
//   http://www.graphviz.org/doc/info/attrs.html
type Edge struct {
	Src *Node
	Dst *Node

	// Color used for edge.
	Color color.Color `name:"color"`

	// Color used to fill the arrowhead. (only if style = filled)
	// http://www.graphviz.org/doc/info/attrs.html#d:fillcolor
	FillColor color.Color `name:"fillcolor"`

	// Color used for text.
	FontColor color.Color `name:"fontcolor"`

	// Label attached to edge.
	Label string `name:"label"`

	// Length of edge.
	Length string `name:"len"`

	// Set style information for the edge.
	// http://www.graphviz.org/doc/info/attrs.html#d:style
	Style string `name:"style"`
}

func (eb Edge) buildAttributes() edgeattrs {
	return buildAttributes(eb)
}

func (eb *Edge) build(nm map[*Node]*dotnode, del string) *dotedge {
	src := nm[eb.Src]
	dst := nm[eb.Dst]
	atr := eb.buildAttributes()
	return &dotedge{atr, del, src, dst}
}
