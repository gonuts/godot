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

import "bytes"
import "testing"

import "godot/attr"

func TestWrite(t *testing.T) {
	var b bytes.Buffer

	nodes := []*Node{
		&Node{Label: "test0"},
		&Node{Label: "test1"},
		&Node{Label: "test2"},
		&Node{Label: "test3"},
		&Node{Label: "test4"},
	}

	edges := []*Edge{
		&Edge{Label: "a", Dst: nodes[0], Src: nodes[1]},
		&Edge{Label: "a", Dst: nodes[1], Src: nodes[2]},
		&Edge{Label: "a", Dst: nodes[2], Src: nodes[3]},
		&Edge{Label: "a", Dst: nodes[3], Src: nodes[4]},
		&Edge{Label: "a", Dst: nodes[4], Src: nodes[0]},
	}

	g := NewGraph(attr.Directed)
	g.AddNodes(nodes...)
	g.AddEdges(edges...)
	d := g.Build()

	d.Write(&b)

	dot := `digraph {
	0 [label="test0"];
	1 [label="test1"];
	2 [label="test2"];
	3 [label="test3"];
	4 [label="test4"];

	1 -> 0 [label="a"];
	2 -> 1 [label="a"];
	3 -> 2 [label="a"];
	4 -> 3 [label="a"];
	0 -> 4 [label="a"];
}
`

	if dot != b.String() {
		t.Errorf("Output was incorrect.")
	}
}
