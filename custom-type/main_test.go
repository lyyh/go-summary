package custom_type

import (
	"fmt"
	"testing"
)

type A = int
type B int

// type C = []int
type S struct {
	Prop string
}

func (s S) name() string {
	return ""
}

type X S

func TestTypeAlias(t *testing.T) {
	s := S{}
	s.name()
	x := X{
		Prop: "asdf",
	}
	fmt.Println(x)
	// x.name 方法无法集成，属性可以集成
}

func TestType(t *testing.T) {
	i := 5
	a := A(i)
	b := B(i)
	fmt.Println("a,b", a, b)
	// if a == A(b) {
	// }
}
