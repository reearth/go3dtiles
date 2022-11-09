package b3dm

import (
	"encoding/binary"
	"io"

	"github.com/pkg/errors"
	"github.com/qmuntal/gltf"
)

const (
	B3DM_PROP_BATCH_LENGTH = "BATCH_LENGTH"
	B3DM_PROP_RTC_CENTER   = "RTC_CENTER"
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

func (h *B3dmHeader) GetSize() int64 {
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

func B3dmFeatureTableDecode(header map[string]interface{}, buff []byte) map[string]interface{} {
	ret := make(map[string]interface{})
	l := getIntegerScalarFeatureValue(header, buff, B3DM_PROP_BATCH_LENGTH)
	ret[B3DM_PROP_BATCH_LENGTH] = l
	ret[B3DM_PROP_RTC_CENTER] = getFloat64Vec3FeatureValue(header, buff, B3DM_PROP_RTC_CENTER)
	return ret
}

func B3dmFeatureTableEncode(header map[string]interface{}, data map[string]interface{}) []byte {
	return nil
}

func (m *B3dm) GetFeatureTableView() B3dmFeatureTable {
	ret := B3dmFeatureTable{}
	ret.BatchLength = m.FeatureTable.Header[B3DM_PROP_BATCH_LENGTH].(int)
	if m.FeatureTable.Header[B3DM_PROP_RTC_CENTER] != nil {
		ret.RtcCenter = m.FeatureTable.Header[B3DM_PROP_RTC_CENTER].([]float64)
	}
	return ret
}

func (m *B3dm) GetHeader() Header {
	return &m.Header
}

func (m *B3dm) GetFeatureTable() *FeatureTable {
	return &m.FeatureTable
}

func (m *B3dm) GetBatchTable() *BatchTable {
	return &m.BatchTable
}

func (m *B3dm) Read(reader io.ReadSeeker) error {

	err := binary.Read(reader, littleEndian, &m.Header)
	if err != nil {
		return errors.Wrap(err, "open failed for b3dm file")
	}

	m.FeatureTable.decode = B3dmFeatureTableDecode

	if err := m.FeatureTable.Read(reader, m.GetHeader()); err != nil {
		return errors.Wrap(err, "failed to read FeatureTable")
	}

	if err := m.BatchTable.Read(reader, m.GetHeader(), m.FeatureTable.GetBatchLength()); err != nil {
		return errors.Wrap(err, "failed to read BatchTable")
	}

	var err1 error
	if m.Model, err1 = loadGltfFromByte(reader); err1 != nil {
		return errors.Wrap(err1, "failed to load glTF file for the given b3dm")
	}

	return nil
}
