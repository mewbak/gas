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

type reg struct {
	name string
	bits int
}

var registers = map[string]reg{
	"al":   {"AX", 8},
	"ah":   {"AX", 8},
	"bl":   {"BX", 8},
	"bh":   {"BX", 8},
	"cl":   {"CX", 8},
	"ch":   {"CX", 8},
	"dl":   {"DX", 8},
	"dh":   {"DX", 8},
	"r8b":  {"R8", 8},
	"r9b":  {"R9", 8},
	"r10b": {"R10", 8},
	"r11b": {"R11", 8},
	"r12b": {"R12", 8},
	"r13b": {"R13", 8},
	"r14b": {"R14", 8},
	"r15b": {"R15", 8},
	"ax":   {"AX", 16},
	"bx":   {"BX", 16},
	"cx":   {"CX", 16},
	"dx":   {"DX", 16},
	"si":   {"SI", 16},
	"di":   {"DI", 16},
	"bp":   {"BP", 16},
	"sp":   {"SP", 16},
	"r8w":  {"R8", 16},
	"r9w":  {"R9", 16},
	"r10w": {"R10", 16},
	"r11w": {"R11", 16},
	"r12w": {"R12", 16},
	"r13w": {"R13", 16},
	"r14w": {"R14", 16},
	"r15w": {"R15", 16},
	"eax":  {"AX", 32},
	"ebx":  {"BX", 32},
	"ecx":  {"CX", 32},
	"edx":  {"DX", 32},
	"esi":  {"SI", 32},
	"edi":  {"DI", 32},
	"ebp":  {"BP", 32},
	"esp":  {"SP", 32},
	"r8d":  {"R8", 32},
	"r9d":  {"R9", 32},
	"r10d": {"R10", 32},
	"r11d": {"R11", 32},
	"r12d": {"R12", 32},
	"r13d": {"R13", 32},
	"r14d": {"R14", 32},
	"r15d": {"R15", 32},
	"rax":  {"AX", 64},
	"rbx":  {"BX", 64},
	"rcx":  {"CX", 64},
	"rdx":  {"DX", 64},
	"rsi":  {"SI", 64},
	"rdi":  {"DI", 64},
	"rbp":  {"BP", 64},
	"rsp":  {"SP", 64},
	"r":    {"R", 64},
	"cs":   {"CS", 64},
	"ss":   {"SS", 64},
	"ds":   {"DS", 64},
	"es":   {"ES", 64},
	"fs":   {"FS", 64},
	"gs":   {"GS", 64},
	"xmm":  {"X", 128},
	"ymm":  {"Y", 256},
	"zmm":  {"Z", 512},
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func stripDigitsRight(s string) (string, string) {
	for i := len(s) - 1; i >= 0; i-- {
		if !isDigit(s[i]) {
			return s[:i+1], s[i+1:]
		}
	}
	return "", s
}

func findReg(s string) (reg, string, bool) {
	if r, ok := registers[s]; ok {
		return r, "", true
	}

	s, suffix := stripDigitsRight(s)
	if r, ok := registers[s]; ok {
		return r, suffix, true
	}

	return reg{}, "", false
}
