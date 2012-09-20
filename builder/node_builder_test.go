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

import "testing"

import "godot/attr"

func TestAddNode(t *testing.T) {
	nb := &Node{Label: "test"}
	gb := NewGraph(attr.Undirected)
	gb.AddNodes(nb)

	l := gb.Nodes()

	if len(l) != 1 {
		t.Fatalf("Length is incorrect.  Should be 1, but is %d.", len(l))
	}

	n := gb.Nodes()[0]

	if n != nb {
		t.Fatal("not equal")
	}
}

func TestAddNodes(t *testing.T) {
	gb := NewGraph(attr.Undirected)

	n1 := &Node{Label: "test1"}
	n2 := &Node{Label: "test2"}
	n3 := &Node{Label: "test3"}

	if count := gb.AddNodes(n1, n2, n3); count != 3 {
		t.Errorf("count is incorrect.  Should be '3', but is '%d'.", count)
	}

	l := gb.Nodes()

	if len(l) != 3 {
		t.Fatalf("Length of node array should be 3, but is %d.", len(l))
	}

	if n1 != l[0] {
		t.Errorf("nodes[0] is out of order.")
	}

	if n2 != l[1] {
		t.Errorf("nodes[1] is out of order.")
	}

	if n3 != l[2] {
		t.Errorf("nodes[2] is out of order.")
	}
}

func TestAddDuplicateNodes(t *testing.T) {
	gb := NewGraph(attr.Undirected)

	n1 := &Node{Label: "test1"}
	n2 := n1
	n3 := n2

	if count := gb.AddNodes(n1, n2, n3); count != 1 {
		t.Errorf("count is incorrect.  Should be '1', but is '%d'.", count)
	}

	l := gb.Nodes()

	if len(l) != 1 {
		t.Fatalf("Length of node array should be 1, but is %d.", len(l))
	}

	if n1 != l[0] {
		t.Errorf("Something truly horriable has happend.")
	}
}

func TestRemoveNode(t *testing.T) {
	gb := NewGraph(attr.Undirected)
	nb := &Node{Label: "test"}
	gb.AddNodes(nb)

	if count := gb.RemoveNodes(nb); count != 1 {
		t.Errorf("count is incorrect.  Should be '1', but is '%d'.", count)
	}

	l := gb.Nodes()

	if len(l) != 0 {
		t.Fatalf("Length is incorrect.  Should be 0, but is %d.", len(l))
	}
}

func TestRemoveNodes(t *testing.T) {
	gb := NewGraph(attr.Undirected)

	n1 := &Node{Label: "test1"}
	n2 := &Node{Label: "test2"}
	n3 := &Node{Label: "test3"}

	gb.AddNodes(n1, n2, n3)

	if count := gb.RemoveNodes(n2); count != 1 {
		t.Errorf("count is incorrect.  Should be '1', but is '%d'.", count)
	}

	l := gb.Nodes()

	if len(l) != 2 {
		t.Errorf("length is incorrect.  Should be '2', but is '%d'.", len(l))
	}

	if n1 != l[0] {
		t.Errorf("nodes[0] is out of order after remove")
	}

	if n3 != l[1] {
		t.Errorf("nodes[1] is out of order after remove")
	}
}

func TestRemoveNonexistentNode(t *testing.T) {
	gb := NewGraph(attr.Undirected)

	n1 := &Node{Label: "test1"}
	n2 := &Node{Label: "test2"}
	n3 := &Node{Label: "test3"}

	ne := &Node{Label: "I don't exist :("}

	gb.AddNodes(n1, n2, n3)

	if count := gb.RemoveNodes(ne); count != 0 {
		t.Errorf("count is incorrect.  Should be '0', but is '%d'.", count)
	}

	l := gb.Nodes()

	if len(l) != 3 {
		t.Fatalf("Length is incorrect.  Should be 3, but is %d.", len(l))
	}
}
