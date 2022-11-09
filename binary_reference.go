package indexer

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
)

type BinaryBodyReference struct {
	ByteOffset    uint32 `json:"byteOffset"`
	ComponentType string `json:"componentType,omitempty"`
}

func (r *BinaryBodyReference) FromMap(d map[string]interface{}) {
	if d[REF_PROP_BYTE_OFFSET] != nil {
		r.ByteOffset = uint32(d[REF_PROP_BYTE_OFFSET].(float64))
	}
	if d[REF_PROP_COMPONENT_TYPE] != nil {
		r.ComponentType = d[REF_PROP_COMPONENT_TYPE].(string)
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
