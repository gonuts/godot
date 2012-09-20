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

import "godot"
import "godot/set"
import "godot/attr"

// A builder for a dot graph.
// For a list of all dot attrs, see: http://www.graphviz.org/doc/info/attrs.html
type Graph struct {
	kind  *attr.GraphKind
	nodes set.OrderedSet
	edges set.OrderedSet
	nTmpl *Node
	eTmpl *Edge

	// Label will appear centered at bottom of graph.
	Label string `name:"label"`
}

// Convenience constructor for the graph builder, which populates all required
// fields.  If using this constructor, it should be possible (although not
// practical) to immediately call Write.
func NewGraph(kind *attr.GraphKind) *Graph {
	nodes := set.New()
	edges := set.New()
	return &Graph{kind: kind, nodes: nodes, edges: edges}
}

// Returns a slice of nodes.  Although the nodes are mutable, assigning Nodes
// to elements of the slice has no effect on the graph.
func (gb *Graph) Nodes() []*Node {
	nodes := make([]*Node, 0, gb.nodes.Count())

	gb.nodes.Visit(func(e interface{}) {
		node := e.(*Node)
		nodes = append(nodes, node)
	})

	return nodes
}

// Returns a slice of edges.  Although the edges are mutable, assigning Edges
// to elements of the slice has no effect on the graph.
func (gb *Graph) Edges() []*Edge {
	edges := make([]*Edge, 0, gb.edges.Count())

	gb.nodes.Visit(func(e interface{}) {
		edge := e.(*Edge)
		edges = append(edges, edge)
	})

	return edges
}

// If set, all nodes will take their attributes from this "template" node, unless
// they specifically override them.
func (gb *Graph) NodeTemplate() *Node {
	return gb.nTmpl
}

func (gb *Graph) SetNodeTemplate(node *Node) {
	gb.nTmpl = node
}

// If set, all edges will take their attributes from this "template" edge, unless
// they specifically override them.
func (gb *Graph) EdgeTemplate() *Edge {
	return gb.eTmpl
}

func (gb *Graph) SetEdgeTemplate(edge *Edge) {
	gb.eTmpl = edge
}

// Adds nodes to the graph, returns number of nodes added.
// Adding a node to the graph multiple times has no effect.
func (gb *Graph) AddNodes(nodes ...*Node) int {
	count := 0
	for _, n := range nodes {
		if gb.nodes.Add(n) {
			count++
		}
	}
	return count
}

// Removes nodes from the graph, returns number of nodes removed.
// Attempting to remove a node which is not a part of the graph has no effect.
func (gb *Graph) RemoveNodes(nodes ...*Node) int {
	count := 0
	for _, n := range nodes {
		if gb.nodes.Remove(n) {
			count++
		}
	}
	return count
}

// Adds edges to the graph, returns number of edges added.
// Adding an edge to the graph multiple times has no effect.
// If an edge has endpoints which are not already in the graph, they will be
// added, Src first, then Dst.  It is therefore possible to create a graph
// entirely through adding edges.
// If an edge has nill endpoints, it will be added successfully and available to
// updates, however if actual nodes are not specified by the time Build is
// called, the edge will not be output.
func (gb *Graph) AddEdges(edges ...*Edge) int {
	count := 0
	for _, e := range edges {
		if gb.edges.Add(e) {
			gb.AddNodes(e.Src, e.Dst)
			count++
		}
	}
	return count
}

// Removes edges from the graph, returns number of edges removed.
// Attempting to remove an edge which is not a part of the graph has no effect.
// If removing the edge disconnects a node from the graph, the node will NOT
// be removed.
func (gb *Graph) RemoveEdges(edges ...*Edge) int {
	count := 0
	for _, e := range edges {
		if gb.edges.Remove(e) {
			count++
		}
	}
	return count
}

// Returns an immutable structure representing the current graph.
func (gb *Graph) Build() godot.Dot {
	nodes, nodemap := buildNodes(gb.nodes)
	edges := buildEdges(gb.edges, nodemap, gb.kind.Delimiter())

	var nTmpl []*attribute
	var eTmpl []*attribute

	if gb.nTmpl != nil {
		nTmpl = gb.nTmpl.buildAttributes()
	}

	if gb.eTmpl != nil {
		eTmpl = gb.eTmpl.buildAttributes()
	}

	return &dotgraph{
		kind:       gb.kind,
		nodes:      nodes,
		edges:      edges,
		nTmpl:      nTmpl,
		eTmpl:      eTmpl,
		attributes: gb.buildAttributes(),
	}
}

// Reflects on the graph and extracts all dot attribute information into
// attribute structures.
func (gb *Graph) buildAttributes() graphattrs {
	return buildAttributes(*gb)
}

func buildNodes(bldrs set.OrderedSet) ([]*dotnode, map[*Node]*dotnode) {
	nodes := make([]*dotnode, 0, bldrs.Count())
	nodemap := make(map[*Node]*dotnode)

	id := 0
	bldrs.Visit(func(e interface{}) {
		bldr := e.(*Node)
		node := bldr.build(id)
		nodes = append(nodes, node)
		nodemap[bldr] = node
		id++
	})

	return nodes, nodemap
}

func buildEdges(bldrs set.OrderedSet, nm map[*Node]*dotnode, del string) []*dotedge {
	edges := make([]*dotedge, 0, bldrs.Count())

	bldrs.Visit(func(e interface{}) {
		bldr := e.(*Edge)
		if bldr.Src != nil && bldr.Dst != nil {
			edges = append(edges, bldr.build(nm, del))
		}
	})
	return edges
}
