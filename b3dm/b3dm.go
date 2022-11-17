package b3dm

import (
	"encoding/binary"
	"errors"
	"io"
	"os"

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

type B3dmReader struct {
	rs io.Reader
}

func NewB3dmReader(r io.Reader) *B3dmReader {
	return &B3dmReader{rs: r}
}

func (r *B3dmReader) DecodeHeader(d *B3dmHeader) error {
	if err := binary.Read(r.rs, littleEndian, d); err != nil {
		return errors.New("failed to read header")
	}
	return nil
}

func (r *B3dmReader) Decode(m *B3dm) error {
	if err := r.DecodeHeader(&m.Header); err != nil {
		return errors.New("failed to decode header")
	}

	m.FeatureTable.decode = B3dmFeatureTableDecode

	if err := m.FeatureTable.Read(r.rs, m.GetHeader()); err != nil {
		return errors.New("failed to read FeatureTable")
	}

	if err := m.BatchTable.Read(r.rs, m.GetHeader(), m.FeatureTable.GetBatchLength()); err != nil {
		return errors.New("failed to read BatchTable")
	}

	var err1 error
	if m.Model, err1 = loadGltfFromByte(r.rs); err1 != nil {
		return errors.New( "failed to load glTF file")
	}

	return nil
}

// Open will open a b3dm file specified by name and return the B3dm.
func Open(fileName string) (*B3dm, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New("open failed")
	}
	defer f.Close()
	b3dmReader := NewB3dmReader(f)
	b3d := new(B3dm)
	if err := b3dmReader.Decode(b3d); err != nil {
		return nil, errors.New("failed to decode the b3dm file")
	}
	return b3d, nil
}
