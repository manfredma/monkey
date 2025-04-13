package compiler

type SymbolScop string

const (
	GlobalScop SymbolScop = "GLOBAL"
)

type Symbol struct{
	Name string
	Scope SymbolScop
	Index int
}

type SymbolTable struct {
	store map[string]Symbol
	numDefinitions int
}

func NewSymbolTable() *SymbolTable {
	s := make (map[string]Symbol)
	return &SymbolTable{store: s}
}

func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Index: s.numDefinitions, Scope: GlobalScop}
	s.store[name] = symbol
	s.numDefinitions++
	return symbol
}

func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	return obj, ok
}