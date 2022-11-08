package indexer

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"os"
)

type indexConfig struct {
	Kind string `json:"kind"`
}

type IndexesConfig struct {
	IdProperty string                 `json:"idProperty"`
	Indexes    map[string]indexConfig `json:"indexes"`
}

func IndexerConfigFromJson(data io.Reader) *IndexesConfig {
	var ic *IndexesConfig
	json.NewDecoder(data).Decode(&ic)
	return ic
}

func ParseIndexerConfigFile(fileName string) (*IndexesConfig, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	return IndexerConfigFromJson(jsonFile), nil
}
