package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "github.com/phodal/design/languages/design"
)

func NewDesignApp() *DesignApp {
	return &DesignApp{}
}

type DesignApp struct {
}

func (j *DesignApp) Start(path string)  {
	context := (*DesignApp)(nil).ProcessFile(path).DesignIt()
	listener := NewDesignAppListener()

	antlr.NewParseTreeWalker().Walk(listener, context)
}

func (j *DesignApp) ProcessFile(path string) *DesignParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewDesignLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewDesignParser(stream)
	return parser
}
