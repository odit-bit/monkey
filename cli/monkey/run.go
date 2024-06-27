package main

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/odit-bit/monkey/eval"
	"github.com/odit-bit/monkey/lexer"
	"github.com/odit-bit/monkey/object"
	"github.com/odit-bit/monkey/parser"
)

func main() {

	if len(os.Args) < 2 {
		log.Println("need source file")
		return
	}
	env := object.NewEnvironment()

	path := os.Args[1]
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}

	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(f); err != nil {
		log.Fatal(err)
	}

	l := lexer.New(buf.String())
	p := parser.New(l)

	prog := p.ParseProgram()
	obj := eval.Eval(prog, env)

	if obj != nil && obj != eval.NULL {
		io.WriteString(os.Stdout, obj.Inspect())
		io.WriteString(os.Stdout, "\n")
	}
}
