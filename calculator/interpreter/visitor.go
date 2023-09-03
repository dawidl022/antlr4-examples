package interpreter

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"

	"github.com/dawidl022/antlr4-examples/calculator/parser"
)

type EvalVisitor struct {
	memory map[string]int
}

func NewEvalVisitor() *EvalVisitor {
	return &EvalVisitor{
		memory: make(map[string]int),
	}
}

func (e *EvalVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(e)
}

func (e *EvalVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, child := range node.GetChildren() {
		e.Visit(child.(antlr.ParseTree))
	}
	return nil
}

func (e *EvalVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

func (e *EvalVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}

func (e *EvalVisitor) VisitInit(ctx *parser.InitContext) interface{} {
	return e.VisitChildren(ctx)
}

func (e *EvalVisitor) VisitAssign(ctx *parser.AssignContext) interface{} {
	return e.VisitChildren(ctx)
}

func (e *EvalVisitor) VisitParens(ctx *parser.ParensContext) interface{} {
	return e.Visit(ctx.Expr())
}

func (e *EvalVisitor) VisitPrintExpr(ctx *parser.PrintExprContext) interface{} {
	val := e.Visit(ctx.Expr())
	fmt.Println(val)
	return nil
}

func (e *EvalVisitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	e.memory[ctx.ID().GetText()] = e.Visit(ctx.Expr()).(int)
	return nil
}

func (e *EvalVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	left := e.Visit(ctx.Expr(0)).(int)
	right := e.Visit(ctx.Expr(1)).(int)

	switch ctx.GetOp().GetTokenType() {
	case parser.CalculatorParserMUL:
		return left * right
	case parser.CalculatorParserDIV:
		return left / right
	default:
		panic("unexpected op")
	}
}

func (e *EvalVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	left := e.Visit(ctx.Expr(0)).(int)
	right := e.Visit(ctx.Expr(1)).(int)

	switch ctx.GetOp().GetTokenType() {
	case parser.CalculatorParserADD:
		return left + right
	case parser.CalculatorLexerSUB:
		return left - right
	default:
		panic("unexpected op")
	}
}

func (e *EvalVisitor) VisitId(ctx *parser.IdContext) interface{} {
	return e.memory[ctx.ID().GetText()]
}

func (e *EvalVisitor) VisitInt(ctx *parser.IntContext) interface{} {
	val, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		panic(err)
	}
	return val
}

func (e *EvalVisitor) VisitClear(ctx *parser.ClearContext) interface{} {
	e.memory = make(map[string]int)
	return nil
}
