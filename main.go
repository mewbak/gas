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

// A tool for extracting functions from object files and transforming them into Go assembly.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/opennota/gas/program"
)

var (
	rFunctionStart = regexp.MustCompile("^[0-9a-f]{16} <([^>]+)>:$")
	rAssemblyLine  = regexp.MustCompile(`^ +([0-9a-f]+):\t((?:[0-9a-f]{2} )+) *(?:\t([a-z0-9]+)(?: *([^ ]+))?)?`)
)

func main() {
	log.SetFlags(log.Lshortfile)
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: gas object function_name")
		os.Exit(1)
	}

	cmd := exec.Command("objdump", "-d", os.Args[1])
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Print(string(output))
		log.Fatal(err)
	}

	scan := bufio.NewScanner(bytes.NewReader(output))
	functionName := os.Args[2]
	found := false
	for scan.Scan() {
		line := scan.Text()
		match := rFunctionStart.FindStringSubmatch(line)
		if match != nil && match[1] == functionName {
			found = true
			break
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	if !found {
		log.Fatal("function not found")
	}

	prog := program.New()
	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			break // end of function
		}
		match := rAssemblyLine.FindStringSubmatch(line)
		if match != nil {
			addr := match[1]
			bytes := strings.Split(strings.TrimSpace(match[2]), " ")
			op := ""
			args := ""
			if len(match) > 3 {
				op = match[3]
			}
			if len(match) > 4 {
				args = match[4]
			}
			if op == "" {
				prog.AddBytes(bytes)
				continue
			}

			prog.AddInst(addr, op, args, bytes)
		}
	}

	fmt.Printf("TEXT ·%s(SB),0,$0\n", functionName)
	for _, inst := range prog.Instrs() {
		if label := prog.Label(inst); label != "" {
			fmt.Printf("%s:\n", label)
		}
		if comment := inst.Comment(); comment != "" {
			fmt.Printf("\t// %s\n", comment)
		}
		fmt.Printf("\t%s\n", inst.String())
	}
}
