package advance

import "fmt"

type Vertex struct {
	X int
	Y int
}

func Struct() {
	fmt.Println(Vertex{1, 2})
	structAccess()
	pointerToStruct()
	structLiteral()
}

func structAccess() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func pointerToStruct() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v, *p)
}

func structLiteral() {
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, v2, v3, p)
}
