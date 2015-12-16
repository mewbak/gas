// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General
// Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

package program

import "strings"

type instInfo struct {
	goName         string
	mustHaveSuffix bool
	jumpOrCall     bool
}

var instructions = map[string]instInfo{
	"add":    {"", true, false},
	"and":    {"", true, false},
	"call":   {"", false, true},
	"cmp":    {"", true, false},
	"dec":    {"", true, false},
	"div":    {"", true, false},
	"hlt":    {"", false, false},
	"imul":   {"", true, false},
	"inc":    {"", true, false},
	"ja":     {"", false, true},
	"jae":    {"", false, true},
	"jb":     {"", false, true},
	"jbe":    {"", false, true},
	"je":     {"", false, true},
	"jge":    {"", false, true},
	"jle":    {"", false, true},
	"jmp":    {"", false, true},
	"jne":    {"", false, true},
	"jns":    {"", false, true},
	"js":     {"", false, true},
	"lea":    {"", true, false},
	"mov":    {"", true, false},
	"movdqa": {"MOVOA", false, false},
	"movdqu": {"MOVOU", false, false},
	"nop":    {"", false, false},
	"not":    {"", true, false},
	"or":     {"", true, false},
	"pop":    {"", true, false},
	"push":   {"", true, false},
	"rep":    {"", false, false},
	"ret":    {"", false, false},
	"sal":    {"", true, false},
	"sar":    {"", true, false},
	"setne":  {"", false, false},
	"sete":   {"", true, false},
	"shl":    {"", true, false},
	"shr":    {"", true, false},
	"sub":    {"", true, false},
	"test":   {"", true, false},
	"xchg":   {"", true, false},
	"xor":    {"", true, false},
}

// An Inst represents an assembler instruction.
type Inst struct {
	addr  string
	op    string
	args  string
	bytes []string

	goAsm string
}

// Comment returns the original GNU assembler instruction string, if it cannot be transformed into Go assembly.  Otherwise it returns an empty string.
func (i *Inst) Comment() string {
	if i.goAsm != "" {
		return ""
	}
	if i.args == "" {
		return i.op
	}
	return i.op + " " + i.args
}

// String returns a string representation of the instruction.
func (i *Inst) String() string {
	if i.goAsm == "" {
		return "BYTE $0x" + strings.Join(i.bytes, "; BYTE $0x")
	}
	return i.goAsm
}
