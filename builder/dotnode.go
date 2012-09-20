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

import "io"
import "fmt"

type dotnode struct {
	id         int
	attributes nodeattrs
}

func (n dotnode) Id() int {
	return n.id
}

func (n dotnode) Write(writer io.Writer) error {
	if err := n.writeAttrs(writer); err != nil {
		return err
	}
	return nil
}

func (n dotnode) writeAttrs(writer io.Writer) error {
	atrs := attrlist(n.attributes).String(0, false)
	if atrs != "" {
		if _, err := fmt.Fprintf(writer, "\t%d [%s];\n", n.Id(), atrs); err != nil {
			return err
		}
	} else {
		if _, err := fmt.Fprintf(writer, "\t%d;\n", n.Id()); err != nil {
			return err
		}
	}
	return nil
}
