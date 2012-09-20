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

import "godot/attr"
import "godot/attr/color"

// A builder for a dot edge.
// For a list of all dot attrs, see: http://www.graphviz.org/doc/info/attrs.html
type Node struct {
	// Color used for node.
	Color color.Color `name:"color"`

	// Color used to fill the background of a node or cluster assuming
	// style=filled, or a filled arrowhead.
	// http://www.graphviz.org/doc/info/attrs.html#d:fillcolor
	FillColor color.Color `name:"fillcolor"`

	// Color used for text.
	FontColor color.Color `name:"fontcolor"`

	Group string `name:"group"`

	// Label attached to node.
	// BUG: To make the label appear blank, specify a space: " " 
	// This bug will be fixed at some point, I am just not sure the best way of
	// going about it.
	Label string `name:"label"`

	// Position of the node (inches)
	Position *attr.Point `name:"pos"`

	// Set the shape of a node.
	Shape *attr.NodeShape `name:"shape"`

	// Set style information for the node.
	// http://www.graphviz.org/doc/info/attrs.html#d:style
	Style string `name:"style"`
}

func (nb Node) buildAttributes() nodeattrs {
	return buildAttributes(nb)
}

func (nb Node) build(id int) *dotnode {
	atrs := nb.buildAttributes()
	return &dotnode{id, atrs}
}

func GenNodes(count int) []*Node {
	nodes := make([]*Node, 0, count)
	for i := 0; i < count; i++ {
		nodes = append(nodes, new(Node))
	}
	return nodes
}

func GenNodesFromTempl(count int, tmpl *Node) []*Node {
	nodes := make([]*Node, 0, count)
	for i := 0; i < count; i++ {
		node := *tmpl
		nodes = append(nodes, &node)
	}
	return nodes
}
