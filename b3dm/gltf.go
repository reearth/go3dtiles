package b3dm

import (
	"io"
	"math"

	"github.com/pkg/errors"
	"github.com/qmuntal/gltf"
)

func loadGltfFromByte(reader io.Reader) (*gltf.Document, error) {
	dec := gltf.NewDecoder(reader)
	doc := new(gltf.Document)
	if err := dec.Decode(doc); err != nil {
		return nil, errors.Wrap(err, "failed to decode the glTF doc")
	}
	return doc, nil
}

// Get the nth value in the buffer described by an accessor with accessorId
func getGltfBufferForValueAt(gltf *gltf.Document, accesorId, n uint32) []byte {
	accessor := gltf.Accessors[accesorId]
	bufferView := gltf.BufferViews[*accessor.BufferView]
	bufferId := bufferView.Buffer
	buffer := gltf.Buffers[bufferId].Data
	buffer = buffer[bufferView.ByteOffset : bufferView.ByteOffset+bufferView.ByteLength]
	valueSize := accessor.ComponentType.ByteSize() * accessor.Type.Components()

	// if no byteStride specified, the buffer is tightly packed
	byteStride := bufferView.ByteStride
	if byteStride == 0 {
		byteStride = valueSize
	}
	pos := accessor.ByteOffset + n*byteStride
	valueBuffer := buffer[pos : pos+valueSize]

	return valueBuffer
}

func readGltfComponent(buff []byte, componentType gltf.ComponentType, n uint32) interface{} {
	switch componentType {
	case (gltf.ComponentFloat):
		inte := littleEndian.Uint32(buff[n : n+componentType.ByteSize()])
		return math.Float32frombits(inte)
	case (gltf.ComponentByte):
		return buff[n]
	case (gltf.ComponentUbyte):
		return uint8(buff[n])
	case (gltf.ComponentShort):
		return int16(littleEndian.Uint16(buff[n : n+componentType.ByteSize()]))
	case (gltf.ComponentUshort):
		return littleEndian.Uint16(buff[n : n+componentType.ByteSize()])
	case (gltf.ComponentUint):
		return littleEndian.Uint32(buff[n:componentType.ByteSize()])
	default:
		return nil
	}
}

func ReadGltfValueAt(gltf *gltf.Document, accesorId, n uint32) []interface{} {
	buffer := getGltfBufferForValueAt(gltf, accesorId, n)
	accessor := gltf.Accessors[accesorId]
	numOfComponents := accessor.Type.Components()
	valueComponents := []interface{}{}
	componentType := accessor.ComponentType

	for i := uint32(0); i < numOfComponents*componentType.ByteSize(); i += componentType.ByteSize() {
		valueComponents = append(valueComponents, readGltfComponent(buffer, componentType, i))
	}

	return valueComponents
}

func GetGltfAttribute(primitive *gltf.Primitive, doc *gltf.Document, name string) ([]interface{}, error) {
	accessors := doc.Accessors
	attributes := primitive.Attributes
	if attributes == nil || len(attributes) == 0 {
		return nil, errors.New("no attributes found")
	}
	att, ok := attributes[name]
	if !ok {
		return nil, errors.New("can't access attribute")
	}
	count := accessors[att].Count

	var res []interface{}
	for i := uint32(0); i < count; i++ {
		res = append(res, ReadGltfValueAt(doc, att, i))
	}

	return res, nil
}
