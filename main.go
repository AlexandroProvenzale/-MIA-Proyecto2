package main

import (
	"Proyecto2/gen"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BaseGrammarListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func main() {

	// Setup the input
	is, _ := antlr.NewFileStream("entrada.txt")
	// Create the Lexer
	lexer := parser.NewGrammarLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewGrammarParser(stream)

	p.BuildParseTrees = true
	tree := p.Start()

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

}
