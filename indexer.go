package indexer

import (
	"fmt"
	"github.com/pkg/errors"
	"path/filepath"
)

func Indexer(tilesetFile, indexConfigFile, outDir string) error {
	tileset, err := ParseTilesetFile(tilesetFile)
	if err != nil {
		return errors.Wrap(err, "failed to parse tileset")
	}
	indexesConfig, err := ParseIndexerConfigFile(indexConfigFile)
	if err != nil {
		return errors.Wrap(err, "failed to parse tileset")

	}

	fmt.Println("Tileset: ", tileset)
	fmt.Println("IndexesConfig: ", indexesConfig)

	tilesetDir := filepath.Dir(tilesetFile)
	fmt.Println("TilesetDir: ", tilesetDir)

	fmt.Println("OutDir: ", outDir)

	return nil
}
