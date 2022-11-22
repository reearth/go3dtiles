package b3dm

import (
	"fmt"
	"io"

	"github.com/qmuntal/gltf"
)

func loadGltfFromByte(reader io.Reader) (*gltf.Document, error) {
	dec := gltf.NewDecoder(reader)
	doc := new(gltf.Document)
	if err := dec.Decode(doc); err != nil {
		return nil, fmt.Errorf("failed to decode the glTF doc: %v", err)
	}
	return doc, nil
}
