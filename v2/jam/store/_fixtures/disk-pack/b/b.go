package b

import "github.com/AlisonOSouza/packr/v2"

func init() {
	packr.New("b-box", "../c")
	packr.New("cb-box", "../c")
}
