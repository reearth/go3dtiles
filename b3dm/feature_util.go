package b3dm

import "encoding/binary"

var (
	littleEndian = binary.LittleEndian
)

func getBatchTableValuesFromRef(ref *BinaryBodyReference, buff []byte, propName string, batchLength int) interface{} {
	if ref != nil {
		offset := int(ref.ByteOffset)
		return buff[offset+batchLength]
	}
	return nil
}

func getIntegerScalarFeatureValue(header map[string]interface{}, buff []byte, propName string) int32 {
	objValue := header[propName]
	switch oref := objValue.(type) {
	case float64:
		return int32(oref)
	}
	return 0
}

func getFloat64Vec3FeatureValue(header map[string]interface{}, buff []byte, propName string) [3]float64 {
	objValue := header[propName]
	switch oref := objValue.(type) {
	case []float64:
		ret := [3]float64{}
		for i := 0; i < 3; i++ {
			ret[i] = oref[i]
		}
		return ret
	case []interface{}:
		ret := [3]float64{}
		for i := 0; i < 3; i++ {
			ret[i] = oref[i].(float64)
		}
		return ret
	}
	return [3]float64{0, 0, 0}
}
