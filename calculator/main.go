package main

import (
	"github.com/antlr4-go/antlr/v4"
	"os"

	"github.com/dawidl022/antlr4-examples/calculator/interpreter"
	"github.com/dawidl022/antlr4-examples/calculator/parser"
)

func main() {
	input := antlr.NewIoStream(os.Stdin)
	lexer := parser.NewCalculatorLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewCalculatorParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	tree := p.Init()
	eval := interpreter.NewEvalVisitor()
	eval.Visit(tree)
}
