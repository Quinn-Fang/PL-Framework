package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"quinn007.com/listeners"
	"quinn007.com/navigator"
	"quinn007.com/parser"
	"quinn007.com/sym_tables"
	"quinn007.com/sym_tables/utils"
	"quinn007.com/uspace"
)

func runListener() {
	input, _ := antlr.NewFileStream("samples/sample_5.go")
	// Create First SymTable
	sym_tables.NewEntryTable()
	// Create Cursor
	navigator.InitCursor()
	curNavigator := navigator.NewNavigator()
	navigator.SetCurNavigator(curNavigator)

	// Create the Lexer
	lexer := parser.NewGoLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewGoParser(stream)

	//listen := GoListener{}
	tree := p.SourceFile()
	antlr.ParseTreeWalkerDefault.Walk(listeners.NewGoListener(p, tree), tree)
	utils.PrintAllSymTale()
	curNavigator.PrintStack()

	curNavigator.PrintCodeSegments()
}

func main() {
	runListener()
	fmt.Println("................... Start testing .......................")
	newNavigator := navigator.GetNavigator()
	event, err := newNavigator.GetNextEvent()
	for err == nil {
		fmt.Println("------------------  ***  --------------------")
		fmt.Printf("%+v ", event)
		if event.GetEventType() == uspace.EventTypeFunction {
			fFunction := event.GetFunction(event.GetEventContext())
			// fmt.Printf("| %+v %+v Executable: %+v\n", fFunction, fFunction.GetReturnValue(), event.GetSymTable().IsExecutable())
			fmt.Printf("| %+v %+v SymTable: %+v \n", fFunction, fFunction.GetReturnValue(), event.GetSymTable().IsExecutable())
		} else if event.GetEventType() == uspace.EventTypeIfElseExpr {
			ifElseExpr, ifElseExprVarNames := event.GetExpr(event.GetEventContext())
			fmt.Printf("| %+v %+v \n", ifElseExpr, ifElseExprVarNames)
		}
		//fmt.Printf("%+v ", event.GetEventType())
		//if event.GetEventType() == uspace.EventTypeFunction {
		//	fmt.Printf("%+v \n", event.GetFunction(event))
		//} else if event.GetEventType() == uspace.EventTypeIfElseExpr {
		//	v1, v2 := event.GetExpr(event)
		//	fmt.Printf("%+v %+v \n", v1, v2)
		//}
		event, err = newNavigator.GetNextEvent()
	}
}
