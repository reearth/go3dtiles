package indexer

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

type BatchTable struct {
	Header map[string]interface{}
	Data   map[string]interface{}
}

func (t *BatchTable) readJSONHeader(data io.Reader) error {
	t.Header = make(map[string]interface{})
	dec := json.NewDecoder(data)
	if err := dec.Decode(&t.Header); err != nil {
		return err
	}
	t.Header = transformBinaryBodyReference(t.Header)
	return nil
}

func (h *BatchTable) Read(reader io.ReadSeeker, header Header, batchLength int) error {
	if header.GetBatchTableJSONByteLength() <= 0 {
		return nil
	}
	jsonLen := header.GetBatchTableJSONByteLength()

	if jsonLen == 0 {
		h.Data = make(map[string]interface{})
		h.Header = make(map[string]interface{})
		return nil
	}

	jsonb := make([]byte, jsonLen)
	if _, err := reader.Read(jsonb); err != nil {
		return errors.Wrap(err, "failed to read json file")
	}

	jsonr := bytes.NewReader(jsonb)
	if err := h.readJSONHeader(jsonr); err != nil {
		return errors.Wrap(err, "failed to read jsonHeader file")
	}

	batchdata := make([]byte, header.GetBatchTableBinaryByteLength())
	if _, err := reader.Read(batchdata); err != nil {
		return errors.Wrap(err, "failed to read batchdata")
	}
	h.Data = make(map[string]interface{})
	for k, v := range h.Header {
		switch t := v.(type) {
		case BinaryBodyReference:
			h.Data[k] = getBatchTableValuesFromRef(&t, batchdata, k, batchLength)
		case []interface{}:
			h.Data[k] = t
		default:
			continue
		}
	}

	return nil
}
