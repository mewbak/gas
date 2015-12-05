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
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func tryCompile(inst string) bool {
	f, err := ioutil.TempFile(os.TempDir(), "asm")
	if err != nil {
		log.Print(err)
		return false
	}

	_, err = f.WriteString("TEXT ·foo(SB),$0\n\t" + inst + "\n")
	if err != nil {
		log.Print(err)
		f.Close()
		return false
	}

	err = f.Close()
	if err != nil {
		log.Print(err)
		return false
	}

	defer os.Remove(f.Name())

	cmd := exec.Command("go", "tool", "asm", "-S", "-dynlink", "-o", f.Name()+".out", f.Name())
	output, err := cmd.CombinedOutput()
	os.Remove(f.Name() + ".out")
	if err != nil {
		log.Println(inst)
		log.Println(string(output))
		return false
	}

	return true
}
