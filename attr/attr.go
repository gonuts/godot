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

// Package attr contains the basic attributes which can be assigned to the
// fields of objects representing nodes, edges, and graphs.
package attr

import "fmt"

// Represents the shape of a node.
//
// Resources:
//
//   http://www.graphviz.org/doc/info/shapes.html
type NodeShape struct {
	name string
}

func (s *NodeShape) String() string {
	return s.name
}

// Represents the kind of graph.
//
// There are only two (valid) possibilites.  If Name returns "graph" then an
// undirected graph will result.  If Name returns "digraph" then a directed
// graph will result.  If Name returns "graph" then Delimiter *must* return
// "--" and if Name returns "digraph" then Delimiter ust return "->".  Any
// other values will result in invalid dot files being generated.
//
// There are predefined variables "Directed" and "Undirected" which contain
// the appropriate values.
//
// Resources:
//   http://www.graphviz.org/doc/info/lang.html
type GraphKind struct {
	// The string that corosponds to the "graph" nonterminal in the dot grammar.
	name string

	// The delimeter that will be output between nodes to represent an edge.
	// This corosponds to the "edgeop" nonterminal in the dot grammar.
	delimiter string
}

func (k *GraphKind) Name() string {
	return k.name
}

func (k *GraphKind) Delimiter() string {
	return k.delimiter
}

// Represents a two dimensional point.  Usually used to position a node or edge.
//
// Resources:
//   http://www.graphviz.org/doc/info/attrs.html#k:point
type Point struct {
	// The point's x coordinate.
	X float32
	// The point's y coordinate.
	Y float32
	// Set to true to (help) ensure the layout algorithms do not move the point.
	Lock bool
}

// "%f,%f('!')?" representing the point (x,y). The optional '!' indicates the
// node position should not change (input-only).
func (p Point) String() string {
	if p.Lock {
		return fmt.Sprintf("%f,%f!", p.X, p.Y)
	}
	return fmt.Sprintf("%f,%f", p.X, p.Y)
}

var (
	Directed   = &GraphKind{"digraph", "->"}
	Undirected = &GraphKind{"graph", "--"}
)

var (
	Box    = &NodeShape{"box"}
	Circle = &NodeShape{"circle"}
	Rect   = &NodeShape{"rect"}
)
