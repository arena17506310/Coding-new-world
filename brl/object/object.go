package object

import (
	"strconv"
	"strings"
)

type ObjectType string

const (
	IntegerObj  = "INTEGER"
	BooleanObj  = "BOOLEAN"
	StringObj   = "STRING"
	ArrayObj    = "ARRAY"
	NullObj     = "NULL"
	ReturnObj   = "RETURN_VALUE"
	ErrorObj    = "ERROR"
	FunctionObj = "FUNCTION"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {
	return IntegerObj
}

func (i *Integer) Inspect() string {
	return strconv.FormatInt(i.Value, 10)
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {
	return BooleanObj
}

func (b *Boolean) Inspect() string {
	return strconv.FormatBool(b.Value)
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType {
	return StringObj
}

func (s *String) Inspect() string {
	return s.Value
}

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType {
	return ArrayObj
}

func (ao *Array) Inspect() string {
	var elements []string
	for _, el := range ao.Elements {
		elements = append(elements, el.Inspect())
	}
	return "[" + strings.Join(elements, ", ") + "]"
}

type Null struct{}

func (n *Null) Type() ObjectType {
	return NullObj
}

func (n *Null) Inspect() string {
	return "null"
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType {
	return ReturnObj
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {
	return ErrorObj
}

func (e *Error) Inspect() string {
	return "Error: " + e.Message
}

/* object 패키지를 구성하는 다양한 객체 유형을 정의하고 있습니다.
각 객체는 Object 인터페이스를 구현하며,
 해당 인터페이스는 모든 객체에 대해 공통적으로 사용되는
  Type() 및 Inspect() 메서드를 정의합니다. */

/*
Integer:

Value 필드를 가지며, 정수 값을 나타냅니다.
Type() 메서드는 "INTEGER"을 반환하고,
Inspect() 메서드는 해당 정수 값을 문자열로 반환합니다.


Boolean:

Value 필드를 가지며, 불리언 값을 나타냅니다.
Type() 메서드는 "BOOLEAN"을 반환하고,
Inspect() 메서드는 해당 불리언 값을 문자열로 반환합니다.


String:

Value 필드를 가지며, 문자열 값을 나타냅니다.
Type() 메서드는 "STRING"을 반환하고,
Inspect() 메서드는 해당 문자열 값을 반환합니다.


Array:

Elements 필드를 가지며, 객체의 배열을 나타냅니다.
Type() 메서드는 "ARRAY"를 반환하고,
Inspect() 메서드는 배열 요소를 문자열로 표현한 후
대괄호로 둘러싼 형태의 문자열을 반환합니다.


Null:

Type() 메서드는 "NULL"을 반환하고,
Inspect() 메서드는 "null"을 반환합니다.


ReturnValue:

Value 필드를 가지며, 반환값을 나타냅니다.
Type() 메서드는 "RETURN_VALUE"를 반환하고,
Inspect() 메서드는 해당 반환값을 반환합니다.


Error:

Message 필드를 가지며, 오류 메시지를 나타냅니다.
Type() 메서드는 "ERROR"를 반환하고,
Inspect() 메서드는 "Error: "를 접두어로 가지는
해당 오류 메시지를 반환합니다.*/

/* 각 객체는 Type() 메서드를 통해 자신의 유형을 식별하고,
Inspect() 메서드를 통해 해당 객체를 문자열로 표현합니다.*/
