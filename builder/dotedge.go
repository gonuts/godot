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

type dotedge struct {
	attributes edgeattrs
	del        string
	src        *dotnode
	dst        *dotnode
}

func (e dotedge) Write(writer io.Writer) error {
  if err := e.writeAttrs(writer); err != nil {
    return err
  }
	return nil
}

func (e dotedge) writeAttrs(writer io.Writer) error {
	src := e.src.Id()
	dst := e.dst.Id()
 	atrs := attrlist(e.attributes).String(0, false)
	if atrs != "" {
		if _, err := fmt.Fprintf(writer, "\t%d %s %d [%s];\n", src, e.del, dst, atrs); err != nil {
			return err
		}
	} else {
		if _, err := fmt.Fprintf(writer, "\t%d %s %d;\n", src, e.del, dst); err != nil {
			return err
		}
	}
  return nil
}
