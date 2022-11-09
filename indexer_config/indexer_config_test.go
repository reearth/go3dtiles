package indexer_config

import (
	"fmt"
	"testing"
)

func TestParseIndexerConfig(t *testing.T) {
	indexerConfig, err := ParseIndexerConfigFile("../example/config.json")
	fmt.Println(indexerConfig)

	if err != nil {
		t.Errorf("failed to parse the indexerconfig json")
	}
}
