package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/odit-bit/monkey/eval"
	"github.com/odit-bit/monkey/lexer"
	"github.com/odit-bit/monkey/object"
	"github.com/odit-bit/monkey/parser"
)

const PROMPT = ">>"

func Start(r io.Reader, w io.Writer) {
	scn := bufio.NewScanner(r)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(w, PROMPT)
		ok := scn.Scan()
		if !ok {
			return
		}

		line := scn.Text()
		if line == "" {
			io.WriteString(w, "cannot be nil")
			io.WriteString(w, "\n")

			continue
		}
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseError(w, p.Errors())
			continue
		}

		obj := eval.Eval(program, env)
		if obj != nil {
			io.WriteString(w, obj.Inspect())
			io.WriteString(w, "\n")
		}

	}
}

func printParseError(w io.Writer, errors []error) {
	for _, err := range errors {
		io.WriteString(w, "\t"+err.Error()+"\n")
	}
}
