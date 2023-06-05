package ast

type Node interface {
	// 공통적인 메서드나 속성들을 정의할 수 있습니다.
}

// Assignment 노드
type Assignment struct {
	Left  *Identifier
	Right Node
}

type Expression interface {
	Node
	expressionNode()
}

// Identifier 노드
type Identifier struct {
	Name string
}

// FunctionDeclaration 노드
type FunctionDeclaration struct {
	Name       *Identifier
	Parameters []*Parameter
	ReturnType *Identifier
	Body       *Block
}

// Parameter 노드
type Parameter struct {
	Name     *Identifier
	DataType *Identifier
}

// Block 노드
type Block struct {
	Statements []Node
}

// IfStatement 노드
type IfStatement struct {
	Condition Expression
	ThenBlock *Block
	ElseBlock *Block
}
