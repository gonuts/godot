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
Package color exports commonly used colors.

The contents of package color logically belong to package attr, however, due to
the sheer number of colors it is probably best if they get their own package.

Currently only a few colors are defined, and only the X11 color scheme is
supported.
*/
package color

// Represents the color of a node, edge, or subgraph background.
//
// Resources:
//   http://www.graphviz.org/doc/info/attrs.html#k:color
type Color interface {
	String() string
}

// Exported to make it easier for users to extend the available named colors.
type Named struct {
	name string
}

func (c Named) String() string {
	return c.name
}

// There are thousands of colors in the X11 color scheme.  A few are available
// here for convenience.
//
// Resources:
//   http://www.graphviz.org/doc/info/colors.html
var (
	Aliceblue      Color = Named{"aliceblue"}
	Aquamarine     Color = Named{"aquamarine"}
	Black          Color = Named{"black"}
	Blue           Color = Named{"blue"}
	Chartreuse     Color = Named{"chartreuse"}
	Cornflowerblue Color = Named{"cornflowerblue"}
	Crimson        Color = Named{"crimson"}
	Cyan           Color = Named{"cyan"}
	Gray           Color = Named{"gray"}
	Red            Color = Named{"red"}
)
