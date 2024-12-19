package advance

import (
	"fmt"
	"strconv"
	"strings"
)

type S struct{}

func (s S) M() {
	fmt.Println(s)
}

func TypeAssertion() {
	var i interface{} = "hello"
	var j I = S{}

	k := j.(I)
	fmt.Println(k)

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f, ok = i.(float64) // panic
	fmt.Println(f)

	getType(j)

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func getType(t interface{}) {
	switch v := t.(type) {
	case string:
		fmt.Printf("%v is a string here\n", v)
	case int:
		fmt.Printf("%v is an int here\n", v)
	default:
		fmt.Printf("type is unknown: %T\n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

// Exercise
type IPAddr [4]byte

func (ip IPAddr) StringV1() string {
	var newIp strings.Builder
	for _, j := range ip {
		newIp.WriteString(fmt.Sprintf("%v.", strconv.Itoa(int(j))))
	}
	return newIp.String()
}

func (ip IPAddr) String() string {
	newIp := make([]string, len(ip))
	for i, j := range ip {
		newIp[i] = strconv.Itoa(int(j))
	}
	return strings.Join(newIp, ".")
}
