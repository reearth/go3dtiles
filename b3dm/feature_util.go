package b3dm

import (
	"encoding/binary"
	"math"
)

var (
	littleEndian = binary.LittleEndian
)

type BatchTableValuesEntry []interface{}

func getBatchTableValuesFromRef(ref *BinaryBodyReference, buff []byte, propName string, batchLength int) []interface{} {
	if ref != nil {
		result := BatchTableValuesEntry{}
		offset := int(ref.ByteOffset)
		containerSize := ContainerTypeSize(ref.ContainerType)
		componentByteSize := ComponentTypeSize(ref.ComponentType)
		switch ref.ComponentType {
		case COMPONENT_TYPE_BYTE:
			for i := 0; i < batchLength*componentByteSize; i += componentByteSize {
				if containerSize == 1 {
					result = append(result, buff[offset+i])
				} else {
					result = append(result, buff[offset+batchLength*containerSize:offset+(batchLength+1)*containerSize])
				}
			}
			return result
		case COMPONENT_TYPE_UNSIGNED_BYTE:
			for i := 0; i < batchLength*componentByteSize; i += componentByteSize {
				out := make([]uint8, containerSize)
				for j := 0; j < containerSize; j++ {
					out[j] = uint8(buff[offset+i*containerSize+j])
					result = append(result, out[j])
				}
			}
			return result
		case COMPONENT_TYPE_SHORT:
			for i := 0; i < batchLength*componentByteSize; i += componentByteSize {
				out := make([]int16, containerSize)
				for j := 0; j < containerSize; j++ {
					out[j] = int16(littleEndian.Uint16(buff[offset+i*containerSize+j : offset+i*containerSize+j+2]))
					result = append(result, out[j])
				}
			}
			return result
		case COMPONENT_TYPE_UNSIGNED_SHORT:
			for i := 0; i < batchLength*componentByteSize; i += componentByteSize {
				out := make([]uint16, containerSize)
				for j := 0; j < containerSize; j++ {
					out[j] = littleEndian.Uint16(buff[offset+i*containerSize+j : offset+i*containerSize+j+2])
					result = append(result, out[j])
				}
			}
			return result
		case COMPONENT_TYPE_INT:
			for i := 0; i < batchLength*componentByteSize; i += componentByteSize {
				out := make([]int32, containerSize)
				for j := 0; j < containerSize; j++ {
					out[j] = int32(littleEndian.Uint32(buff[offset+i*containerSize+j : offset+i*containerSize+j+4]))
					result = append(result, out[j])
				}
			}
			return result
		case COMPONENT_TYPE_UNSIGNED_INT:
			for i := 0; i < batchLength*componentByteSize; i += componentByteSize {
				out := make([]uint32, containerSize)
				for j := 0; j < containerSize; j++ {
					out[j] = littleEndian.Uint32(buff[offset+i*containerSize+j : offset+i*containerSize+j+4])
					result = append(result, out[j])
				}
			}
			return result
		case COMPONENT_TYPE_FLOAT:
			for i := 0; i < batchLength*componentByteSize; i += componentByteSize {
				out := make([]float32, containerSize)
				for j := 0; j < containerSize; j++ {
					inte := littleEndian.Uint32(buff[offset+i*containerSize+j : offset+i*containerSize+j+4])
					out[j] = math.Float32frombits(inte)
					result = append(result, out[j])
				}
			}
			return result
		case COMPONENT_TYPE_DOUBLE:
			for i := 0; i < batchLength*componentByteSize; i += componentByteSize {
				out := make([]float64, containerSize)
				for j := 0; j < containerSize; j++ {
					inte := littleEndian.Uint64(buff[offset+i*containerSize+j : offset+i*containerSize+j+8])
					out[j] = math.Float64frombits(inte)
					result = append(result, out[j])
				}
			}
			return result
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
