package b3dm

type Header interface {
	CalcSize() int64

	GetByteLength() uint32

	GetFeatureTableJSONByteLength() uint32
	GetFeatureTableBinaryByteLength() uint32

	GetBatchTableJSONByteLength() uint32
	GetBatchTableBinaryByteLength() uint32
}