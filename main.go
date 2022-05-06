package main

import (
	"Proyecto2/gen"
	"bufio"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"strings"
)

type TreeShapeListener struct {
	*parser.BaseGrammarListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func Exec(path string) {
	is, _ := antlr.NewFileStream(path)
	// Create the Lexer
	lexer := parser.NewGrammarLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewGrammarParser(stream)

	p.BuildParseTrees = true
	tree := p.Start()

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}

func interpretar(entrada string) {
	is := antlr.NewInputStream(entrada)

	// crear analizador lexico
	lexer := parser.NewGrammarLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// crear analizador sintactico
	p := parser.NewGrammarParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), p.Start())
}

func main() {
	// Setup the input

	fmt.Println("Ingrese un comando:")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	comando := in.Text()

	if strings.ToLower(string(comando[0:4])) == "exec" {
		path := strings.ReplaceAll(comando[11:(len(comando))], "\"", "")
		Exec(path) // exec -path=
	} else {
		interpretar(comando)
	}

}
