package tileset

import (
	"fmt"
	"testing"
)

func TestParseTilesetFile(t *testing.T) {
	tileset, err := ParseTilesetFile("../example/tileset.json")
	fmt.Println(tileset)

	if err != nil {
		t.Errorf("failed to parse the tileset")
	}
}
