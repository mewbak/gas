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

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/opennota/gas/scanner"
)

// A Program represents an assembler program.
type Program struct {
	instrs []*Inst
	labels map[string]struct{}
}

// New returns a new Program.
func New() *Program {
	p := Program{
		labels: make(map[string]struct{}),
	}
	return &p
}

// AddInst adds an instruction to the program.
func (p *Program) AddInst(addr, op, args string, bytes []string) {
	i := &Inst{
		addr:  addr,
		op:    op,
		args:  args,
		bytes: bytes,
	}
	p.transform(i)
	p.instrs = append(p.instrs, i)
}

// AddBytes adds opcodes to a latest instruction.
func (p *Program) AddBytes(b []string) {
	p.instrs[len(p.instrs)-1].bytes = append(p.instrs[len(p.instrs)-1].bytes, b...)
}

// Instrs returns a slice of instructions the program contains.
func (p *Program) Instrs() []*Inst {
	return p.instrs
}

func stripSuffix(s string) string {
	switch s[len(s)-1] {
	case 'b', 'w', 'l', 'q':
		return s[:len(s)-1]
	}
	return s
}

func reverseOperands(p []byte) []byte {
	i := bytes.IndexByte(p, ',')
	buf := make([]byte, 0, len(p))
	buf = append(buf, p[i+2:]...)
	buf = append(buf, ", "...)
	buf = append(buf, p[:i]...)
	return buf
}

var rAddress = regexp.MustCompile("^[0-9a-f]+$")

func (p *Program) transform(i *Inst) error {
	var ii instInfo
	var ok bool
	op := i.op
	if ii, ok = instructions[op]; ok {
		if ii.goName != "" {
			op = ii.goName
		}
	} else {
		op2 := stripSuffix(op)
		if ii, ok = instructions[op2]; ok {
			op = op2
		}
	}

	scan := scanner.NewScanner(strings.NewReader(i.args))
	nextRegister := false
	insideParentheses := false
	n := 0
	bits := 64
	var buf []byte
	for scan.Scan() {
		tok := scan.Text()
		switch tok {
		case "%":
			nextRegister = true
		case "(":
			insideParentheses = true
			buf = append(buf, '(')
		case ")":
			insideParentheses = false
			buf = append(buf, ')')
			n = 0
		case ",":
			if !insideParentheses {
				buf = append(buf, ", "...)
			} else {
				// (%ax,%bx,2) => (AX)(BX*2)
				// (,%bx,2) => (BX*2)
				if n == 0 && buf[len(buf)-1] != '(' {
					buf = append(buf, ")("...)
				} else if n > 0 {
					buf = append(buf, '*')
				}
				n++
			}
		default:
			if nextRegister {
				if r, suffix, ok := findReg(tok); ok {
					buf = append(buf, r.name...)
					buf = append(buf, suffix...)
					if !insideParentheses {
						bits = r.bits
					}
				} else {
					buf = append(buf, strings.ToUpper(tok)...)
				}
				nextRegister = false
				continue
			}

			buf = append(buf, tok...)
		}
	}
	if err := scan.Err(); err != nil {
		return err
	}

	if op == "cmp" {
		// cmp $0, AX => cmp AX, $0
		buf = reverseOperands(buf)
	}

	if ii.mustHaveSuffix {
		switch bits {
		case 64:
			op += "q"
		case 32:
			op += "l"
		case 16:
			op += "w"
		case 8:
			op += "b"
		}
	}

	op = strings.ToUpper(op)

	if ii.jumpOrCall && rAddress.Match(buf) {
		addr := string(buf)
		p.labels[addr] = struct{}{}
		i.goAsm = op + " _L" + addr
	} else {
		asm := op
		if len(buf) > 0 {
			asm += " " + string(buf)
		}
		if ok := tryCompile(asm); !ok {
			return nil
		}

		i.goAsm = asm
	}

	return nil
}

// Label returns a label for an instruction, if there are any jumps or calls referencing it.  Otherwise it returns an empty string.
func (p *Program) Label(i *Inst) string {
	if _, ok := p.labels[i.addr]; !ok {
		return ""
	}
	return "_L" + i.addr
}
