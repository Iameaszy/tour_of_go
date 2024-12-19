package advance

import (
	"fmt"
	"math"
)

// Note: In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both
func Methods() {
	fmt.Println("This is a method")
	v := MethodVertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println("Square root after scaling", v.Abs())
	g := Greetings{
		"1": "hello",
		"2": "Hi",
	}
	fmt.Println("I dey greet", g.getElem("1"))
	var f1 MyFloat = 23
	f2 := MyFloat(24)
	fmt.Println("My float", f1.getValue(), f2.getValue())
	scaleFunc(&v, 10)
	fmt.Println("scale function", v.Abs())
}

type MethodVertex struct {
	X, Y float64
}

func (v MethodVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *MethodVertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func scaleFunc(v *MethodVertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Greetings map[string]string

func (m Greetings) getElem(elem string) string {
	data := m[elem]
	return data
}

type MyFloat float64

func (myFloat MyFloat) getValue() float64 {
	return float64(myFloat)
}
