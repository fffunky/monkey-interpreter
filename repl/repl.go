package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey/lexer"
	"monkey/parser"
	"monkey/evaluator"
)

const PROMPT = ">> "

const MONKEY_ASCII = `
  ╭~╮
 C oo⦘
 _❪ ^❫
/   ~\
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if (!scanned) {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_ASCII)
	io.WriteString(out, "Whoops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t" + msg + "\n")
	}
}