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

package main

import "os"
import "math"

import "godot/builder"
import "godot/attr"
import "godot/attr/color"

func setPos(n int, radius float32, node *builder.Node) {
	pos := float64((n * 72) % 360) // Where n is the nth point of the pentagon
	radians := pos * (math.Pi / 180)

	node.Position = &attr.Point{
		X:    radius * float32(math.Cos(radians)),
		Y:    radius * float32(math.Sin(radians)),
		Lock: true,
	}
}

func setColor(nodes []*builder.Node) {
	nodes[0].FillColor = color.Red
	nodes[1].FillColor = color.Cornflowerblue
	nodes[2].FillColor = color.Chartreuse
	nodes[3].FillColor = color.Red
	nodes[4].FillColor = color.Cornflowerblue
	nodes[5].FillColor = color.Cornflowerblue
	nodes[6].FillColor = color.Red
	nodes[7].FillColor = color.Red
	nodes[8].FillColor = color.Chartreuse
	nodes[9].FillColor = color.Chartreuse
}

func makeGraph() *builder.Graph {
	gb := builder.NewGraph(attr.Undirected)
	gb.Label = "Petersen Graph (3-Coloring)"

	// Create default attributes for nodes.
	nTmpl := &builder.Node{
		Label: " ",
		Color: color.Black,
		Shape: attr.Circle,
		Style: "filled",
	}
	gb.SetNodeTemplate(nTmpl)

	return gb
}

func makeNodes() []*builder.Node {
	// Creates 10 nodes
	nodes := builder.GenNodes(10)

	setColor(nodes)

	// Set the position of the outer ring.
	for i, n := range nodes[0:5] {
		setPos(i, 2, n)
	}

	// Set the position of the inner star.
	for i, n := range nodes[5:] {
		setPos(i, 1, n)
	}
	return nodes
}

func makeEdges(nodes []*builder.Node) []*builder.Edge {
	edges := []*builder.Edge{
		// Wire-up the edges for the outer ring.
		&builder.Edge{Src: nodes[0], Dst: nodes[1]},
		&builder.Edge{Src: nodes[1], Dst: nodes[2]},
		&builder.Edge{Src: nodes[2], Dst: nodes[3]},
		&builder.Edge{Src: nodes[3], Dst: nodes[4]},
		&builder.Edge{Src: nodes[4], Dst: nodes[0]},

		// Wire-up the edges from the outer ring to the inner star.
		&builder.Edge{Src: nodes[0], Dst: nodes[5]},
		&builder.Edge{Src: nodes[1], Dst: nodes[6]},
		&builder.Edge{Src: nodes[2], Dst: nodes[7]},
		&builder.Edge{Src: nodes[3], Dst: nodes[8]},
		&builder.Edge{Src: nodes[4], Dst: nodes[9]},

		// Wire-up the edges in the inner star.
		&builder.Edge{Src: nodes[5], Dst: nodes[7]},
		&builder.Edge{Src: nodes[7], Dst: nodes[9]},
		&builder.Edge{Src: nodes[9], Dst: nodes[6]},
		&builder.Edge{Src: nodes[6], Dst: nodes[8]},
		&builder.Edge{Src: nodes[8], Dst: nodes[5]},
	}
	return edges
}

func main() {
	graph := makeGraph()
	nodes := makeNodes()
	edges := makeEdges(nodes)

	graph.AddNodes(nodes...)
	graph.AddEdges(edges...)

	dot := graph.Build()
	dot.Write(os.Stdout)
}
