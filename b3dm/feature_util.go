package b3dm

import (
	"encoding/binary"
	"math"
)

var (
	littleEndian = binary.LittleEndian
)

func getBatchTableValuesFromRef(ref *BinaryBodyReference, buff []byte, propName string, batchLength int) interface{} {
	if ref != nil {
		offset := int(ref.ByteOffset)
		containerSize := ContainerTypeSize(ref.ContainerType)
		switch ref.ComponentType {
		case COMPONENT_TYPE_BYTE:
			if containerSize == 1 {
				return buff[offset+batchLength]
			}
			return buff[offset+batchLength*containerSize : offset+(batchLength+1)*containerSize]
		case COMPONENT_TYPE_UNSIGNED_BYTE:
			if containerSize == 1 {
				return uint8(buff[offset+batchLength])
			}
			out := make([]uint8, containerSize)
			for i := 0; i < containerSize; i++ {
				out[i] = uint8(buff[offset+batchLength*containerSize+i])
			}
			return out
		case COMPONENT_TYPE_SHORT:
			if containerSize == 1 {
				return int16(littleEndian.Uint16(buff[offset+batchLength : offset+batchLength+2]))
			}
			out := make([]int16, containerSize)
			for i := 0; i < containerSize; i++ {
				out[i] = int16(littleEndian.Uint16(buff[offset+batchLength*containerSize+i : offset+batchLength*containerSize+i+2]))
			}
			return out
		case COMPONENT_TYPE_UNSIGNED_SHORT:
			if containerSize == 1 {
				return littleEndian.Uint16(buff[offset+batchLength : offset+batchLength+2])
			}
			out := make([]uint16, containerSize)
			for i := 0; i < containerSize; i++ {
				out[i] = littleEndian.Uint16(buff[offset+batchLength*containerSize+i : offset+batchLength*containerSize+i+2])
			}
			return out
		case COMPONENT_TYPE_INT:
			if containerSize == 1 {
				return int32(littleEndian.Uint32(buff[offset+batchLength : offset+batchLength+4]))
			}
			out := make([]int32, containerSize)
			for i := 0; i < containerSize; i++ {
				out[i] = int32(littleEndian.Uint32(buff[offset+batchLength*containerSize+i : offset+batchLength*containerSize+i+4]))
			}
			return out
		case COMPONENT_TYPE_UNSIGNED_INT:
			if containerSize == 1 {
				return littleEndian.Uint32(buff[offset+batchLength : offset+batchLength+4])
			}
			out := make([]uint32, containerSize)
			for i := 0; i < containerSize; i++ {
				out[i] = littleEndian.Uint32(buff[offset+batchLength*containerSize+i : offset+batchLength*containerSize+i+4])
			}
			return out
		case COMPONENT_TYPE_FLOAT:
			if containerSize == 1 {
				i := littleEndian.Uint32(buff[offset+batchLength : offset+batchLength+4])
				return math.Float32frombits(i)
			}
			out := make([]float32, containerSize)
			for i := 0; i < containerSize; i++ {
				inte := littleEndian.Uint32(buff[offset+batchLength*containerSize+i : offset+batchLength*containerSize+i+4])
				out[i] = math.Float32frombits(inte)
			}
			return out
		case COMPONENT_TYPE_DOUBLE:
			if containerSize == 1 {
				i := littleEndian.Uint64(buff[offset+batchLength : offset+batchLength+8])
				return math.Float64frombits(i)
			}
			out := make([]float64, containerSize)
			for i := 0; i < containerSize; i++ {
				inte := littleEndian.Uint64(buff[offset+batchLength*containerSize+i : offset+batchLength*containerSize+i+8])
				out[i] = math.Float64frombits(inte)
			}
			return out
		}
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
