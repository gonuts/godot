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

type OrderedSet interface {
	Count() int

	// Visit (in order of addition)
	Visit(func(interface{}))

	// Returns true if the element was added to the set. False if the element is
	// already a member of the set.
	Add(interface{}) bool

	// Returns true if the element was removed, false if the element was not
	// a member of the set.
	Remove(interface{}) bool

	// Returns true if the element is a member of the set.
	Contains(interface{}) bool
}
