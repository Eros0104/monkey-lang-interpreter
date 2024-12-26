package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStataments(t *testing.T) {
	input := `
    let x = 5;
    let y = 10;
    let foobar = 838383;
  `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. dot=%d", program.Statements)
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, expectedIdentifier string) bool {
	if s.TokenLiteral() != "let" {
		t.Fatalf("s.TokenLiteral is not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s is not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != expectedIdentifier {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", expectedIdentifier, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != expectedIdentifier {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", expectedIdentifier, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}
