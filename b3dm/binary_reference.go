package b3dm

const (
	COMPONENT_TYPE_BYTE           = "BYTE"
	COMPONENT_TYPE_UNSIGNED_BYTE  = "UNSIGNED_BYTE"
	COMPONENT_TYPE_SHORT          = "SHORT"
	COMPONENT_TYPE_UNSIGNED_SHORT = "UNSIGNED_SHORT"
	COMPONENT_TYPE_INT            = "INT"
	COMPONENT_TYPE_UNSIGNED_INT   = "UNSIGNED_INT"
	COMPONENT_TYPE_FLOAT          = "FLOAT"
	COMPONENT_TYPE_DOUBLE         = "DOUBLE"
)

const (
	REF_PROP_BYTE_OFFSET    = "byteOffset"
	REF_PROP_COMPONENT_TYPE = "componentType"
	REF_PROP_TYPE           = "type"
)

const (
	CONTAINER_TYPE_SCALAR = "SCALAR"
	CONTAINER_TYPE_VEC2   = "VEC2"
	CONTAINER_TYPE_VEC3   = "VEC3"
	CONTAINER_TYPE_VEC4   = "VEC4"
)

func ContainerTypeSize(tp string) int {
	switch tp {
	case CONTAINER_TYPE_SCALAR:
		return 1
	case CONTAINER_TYPE_VEC2:
		return 2
	case CONTAINER_TYPE_VEC3:
		return 3
	case CONTAINER_TYPE_VEC4:
		return 4
	default:
		return 0
	}
}

type BinaryBodyReference struct {
	ByteOffset    uint32 `json:"byteOffset"`
	ComponentType string `json:"componentType,omitempty"`
	ContainerType string `json:"type,omitempty"`
}

func (r *BinaryBodyReference) FromMap(d map[string]interface{}) {
	if d[REF_PROP_BYTE_OFFSET] != nil {
		r.ByteOffset = uint32(d[REF_PROP_BYTE_OFFSET].(float64))
	}
	if d[REF_PROP_COMPONENT_TYPE] != nil {
		r.ComponentType = d[REF_PROP_COMPONENT_TYPE].(string)
	}
	if d[REF_PROP_TYPE] != nil {
		r.ContainerType = d[REF_PROP_TYPE].(string)
	}
}

func transformBinaryBodyReference(m map[string]interface{}) map[string]interface{} {
	ref := make(map[string]interface{})
	for k, v := range m {
		switch tp := v.(type) {
		case map[string]interface{}:
			r := new(BinaryBodyReference)
			r.FromMap(tp)
			ref[k] = *r
		default:
			ref[k] = v
		}
	}
	return ref
}
