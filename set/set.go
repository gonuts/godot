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

package set

import "container/list"

// Requirements:
//
// I wrote this because I need:
//   1) an ordered collection (uses a linked list)
//   2) fast "Contains" (uses a map) -- O(1)
//   3) reasonably fast append (linked list) -- O(1)
//   4) reasonably fast removal of an arbitrary element (linked list) -- O(n)
// and I do not believe that the go standard libraries contain such a structure.
type orderedset struct {
	index   map[interface{}]bool
	content list.List
}

func New() OrderedSet {
  index := make(map[interface{}]bool)
  return &orderedset{index: index}
}

func (s *orderedset) Count() int {
	return s.content.Len()
}

func (s *orderedset) Visit(visitor func(interface{})) {
	for e := s.content.Front(); e != nil; e = e.Next() {
		visitor(e.Value)
	}
}

func (s *orderedset) Add(element interface{}) bool {
	if _, ok := s.index[element]; ok {
		return false
	}
	s.content.PushBack(element)
  s.index[element] = true
	return true
}

func (s *orderedset) Remove(element interface{}) bool {
	if _, ok := s.index[element]; !ok {
		return false
	}

	for e := s.content.Front(); e != nil; e = e.Next() {
    if e.Value == element {
	    s.content.Remove(e)
      delete(s.index, element)
	    return true
    }
	}
  // TODO:  panic?
  return false
}

func (s *orderedset) Contains(element interface{}) bool {
	_, ok := s.index[element]
	return ok
}
