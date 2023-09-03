package interpreter

import (
	"fmt"
	"strconv"

	"github.com/dawidl022/antlrexamples/calculator/parser"
)

type EvalVisitor struct {
	*parser.BaseCalculatorVisitor
	memory map[string]int
}

func NewEvalVisitor() *EvalVisitor {
	return &EvalVisitor{
		BaseCalculatorVisitor: &parser.BaseCalculatorVisitor{},
		memory:                make(map[string]int),
	}
}

func (e *EvalVisitor) VisitPrintExpr(ctx *parser.PrintExprContext) interface{} {
	fmt.Println("print expr")
	val := e.Visit(ctx.Expr())
	fmt.Println(val)
	return nil
}

func (e *EvalVisitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	fmt.Println("assign")
	e.memory[ctx.ID().GetText()] = e.Visit(ctx.Expr()).(int)
	return nil
}

func (e *EvalVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	fmt.Println("mul div")

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
	fmt.Println("add sub")

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
	fmt.Println("id")

	return e.memory[ctx.ID().GetText()]
}

func (e *EvalVisitor) VisitInt(ctx *parser.IntContext) interface{} {
	fmt.Println("int")

	val, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		panic(err)
	}
	return val
}
