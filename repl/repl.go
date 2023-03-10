package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey_learn/evaluator"
	"monkey_learn/lexer"
	"monkey_learn/parser"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		_, _ = fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			_, _ = io.WriteString(out, "\n")
			continue
		}

		eval := evaluator.Eval(program)
		if eval != nil {
			_, _ = io.WriteString(out, eval.Inspect())
			_, _ = io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	_, _ = io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	_, _ = io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}
