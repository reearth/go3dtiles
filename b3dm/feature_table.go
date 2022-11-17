package b3dm

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

type featureTableDecode func(header map[string]interface{}, buff []byte) map[string]interface{}

type FeatureTable struct {
	Header map[string]interface{}
	Data   map[string]interface{}

	decode featureTableDecode
}

func (h *FeatureTable) GetBatchLength() int {
	if h.Header["BATCH_LENGTH"] != nil {
		switch d := h.Header["BATCH_LENGTH"].(type) {
		case int:
			return d
		case float64:
			return int(d)
		}
	}
	return 0
}

func (t *FeatureTable) readJSONHeader(data io.Reader, jsonLength int) error {
	jdata := make([]byte, jsonLength)
	_, err := data.Read(jdata)
	dec := json.NewDecoder(bytes.NewBuffer(jdata))
	if err != nil {
		return nil
	}
	t.Header = make(map[string]interface{})
	if err := dec.Decode(&t.Header); err != nil {
		return errors.New("failed to decode the json file")
	}
	t.Header = transformBinaryBodyReference(t.Header)
	return nil
}

func (h *FeatureTable) readData(reader io.Reader, buffLength int) error {
	if buffLength == 0 {
		return nil
	}
	bdata := make([]byte, buffLength)
	_, err := reader.Read(bdata)
	if err != nil {
		return errors.New("failed to read the binary data")
	}
	h.Data = h.decode(h.Header, bdata)
	return nil
}

func (h *FeatureTable) Read(reader io.Reader, header Header) error {
	err := h.readJSONHeader(reader, int(header.GetFeatureTableJSONByteLength()))
	if err != nil {
		return errors.New("failed to read FeatureTable header")
	}
	err = h.readData(reader, int(header.GetFeatureTableBinaryByteLength()))
	if err != nil {
		return errors.New("failed to read FeatureTable Data")
	}
	return nil
}
