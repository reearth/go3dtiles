package indexer

import "testing"

func TestIndexer(t *testing.T) {
	Indexer("example/tileset.json", "example/config.json", ".")
}