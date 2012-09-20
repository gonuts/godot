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

/*
Package godot provides a method of programatically building graphs which can
be output into AT&T "dot" syntax.

For more information about AT&T dot syntax, consult the documentation:
  Grammar:  http://www.graphviz.org/doc/info/lang.html
  Examples: http://www.graphviz.org/Gallery.php
  Graphviz Software: http://www.graphviz.org/Download.php

A short example follows:

  package main

  import "os"
  import "godot"

  func main() {
    graph := NewGraph(godot.Undirected)

    nodeA := &godot.Node{Label: "A"}
    nodeB := &godot.Node{Label: "B"}

    graph.AddNodes(nodeA, nodeB)

    edge1 := &godot.Edge{Src: nodeA, Dst: nodeB}

    graph.AddEdges(edge1)

    dot := graph.Build()
    dot.Write(os.Stdout)
  }

The above example will write the following to stdout:

  graph {
    0 [label="A"];
    1 [label="B"];

    0 -- 1;
  }

Which can be converted to a png file using the "dot" program available from
www.graphviz.org:

  $ go run main.go | dot -Tpng > example.png
*/
package godot

import "io"

// Represents a dot graph.
type Dot interface {
	// Writes the dot representation of this graph to "writer".
	Write(writer io.Writer) error
}
