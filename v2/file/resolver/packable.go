package resolver

import "github.com/AlisonOSouza/packr/v2/file"

type Packable interface {
	Pack(name string, f file.File) error
}
