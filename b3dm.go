package indexer

import (
	"github.com/qmuntal/gltf"
)

type B3dmHeader struct {
	Magic                        [4]byte
	Version                      uint32
	ByteLength                   uint32
	FeatureTableJSONByteLength   uint32
	FeatureTableBinaryByteLength uint32
	BatchTableJSONByteLength     uint32
	BatchTableBinaryByteLength   uint32
}

func (h *B3dmHeader) CalcSize() int64 {
	return 28
}

func (h *B3dmHeader) GetByteLength() uint32 {
	return h.ByteLength
}

func (h *B3dmHeader) GetFeatureTableJSONByteLength() uint32 {
	return h.FeatureTableJSONByteLength
}

func (h *B3dmHeader) GetFeatureTableBinaryByteLength() uint32 {
	return h.FeatureTableBinaryByteLength
}

func (h *B3dmHeader) GetBatchTableJSONByteLength() uint32 {
	return h.BatchTableJSONByteLength
}

func (h *B3dmHeader) GetBatchTableBinaryByteLength() uint32 {
	return h.BatchTableBinaryByteLength
}

type B3dmFeatureTable struct {
	BatchLength int
	RtcCenter   []float64
}

type B3dm struct {
	Header       B3dmHeader
	FeatureTable FeatureTable
	BatchTable   BatchTable
	Model        *gltf.Document
}