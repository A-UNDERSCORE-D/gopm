package main

import (
	"encoding"
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
)

const testData = `
some = "test"
[patterns]
data = "asd"
stuff = ["x", "y", "z"]
`

type testStruct struct {
	Some     string    `toml:"some"`
	Patterns RawStruct `toml:"patterns"`
}

type RawStruct struct {
	data []byte
}

func (r *RawStruct) UnmarshalText(data []byte) error {
	r.data = append([]byte(nil), data...)
	fmt.Println("CALLED")

	return nil
}

type Raw []byte

var _ encoding.TextUnmarshaler = (*Raw)(nil)

func (r *Raw) UnmarshalText(data []byte) error {
	if r == nil {
		return errors.New("nil raw instance used")
	}

	*r = append((*r)[0:0], data...)
	return nil
}

func main() {
	test := &testStruct{}

	fmt.Println(toml.Unmarshal([]byte(testData), test))

	fmt.Printf("%#v -- %#v\n", test, test.Patterns)
}
