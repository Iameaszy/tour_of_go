package advance

import (
	"fmt"
	"io"
	"strings"
)

func Readers() {
	readBytes()
	// s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// r := rot13Reader{s}
	// io.Copy(os.Stdout, &r)
}

func readBytes() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// Todo: Finish exercise later
type rot13Reader struct {
	r io.Reader
}

// func (r rot13Reader) Read([]byte) (int, error) {
// 	b := make([]byte, 13)
// 	return r.Read(b)
// }
