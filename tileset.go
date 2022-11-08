package indexer

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"os"
)

const (
	TILE_REFINE_ADD     = "ADD"
	TILE_REFINE_REPLACE = "REPLACE"
)

var (
	TileDefaultTransform = [16]float64{1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0}
)

type Schema struct {
	Maximum float64 `json:"maximum,omitempty"`
	Minimum float64 `json:"minimum,omitempty"`
}

type Asset struct {
	Version        string                  `json:"version"`
	TilesetVersion *string                 `json:"tileSetVersion,omitempty"`
	Extentions     *map[string]interface{} `json:"extensions,omitempty"`
	Extras         *interface{}            `json:"extras,omitempty"`
}

type BoundingVolume struct {
	Box        [12]float64             `json:"box,omitempty"`
	Region     [6]float64              `json:"region,omitempty"`
	Sphere     [4]float64              `json:"sphere,omitempty"`
	Extentions *map[string]interface{} `json:"extension,omitempty"`
	Extras     *interface{}            `json:"extras,omitempty"`
}

type Content struct {
	BoundingVolume BoundingVolume `json:"boundingVolume,omitempty"`
	// change it to "uri" https://github.com/CesiumGS/3d-tiles/tree/main/specification#content
	URL        string                  `json:"url"`
	Extentions *map[string]interface{} `json:"extension,omitempty"`
	Extras     *interface{}            `json:"extras,omitempty"`
}

type Tile struct {
	BoundingVolume      BoundingVolume          `json:"boundingVolume"`
	ViewerRequestVolume *BoundingVolume         `josn:"viewerRequestVolume,omitempty"`
	GeometricError      float64                 `json:"geometricError"`
	Refine              *string                 `json:"refine,omitempty"`
	Transform           [16]float64             `json:"transform,omitempty"`
	Content             *Content                `json:"content,omitempty"`
	Children            []Tile                  `json:"children,omitempty"`
	Extentions          *map[string]interface{} `json:"extensions,omitempty"`
	Extras              *interface{}            `json:"extras,omitempty"`
}

type Tileset struct {
	Asset             Asset                   `json:"asset"`
	Properties        **map[string]Schema     `json:"properties,omitempty"`
	GeometricError    float64                 `json:"geometricError"`
	Root              Tile                    `json:"root"`
	ExtenstionsUsed   []string                `json:"extensionsUsed,omitempty"`
	ExtensionRequired []string                `json:"extensionsRequired,omitempty"`
	Extentions        *map[string]interface{} `json:"extensions,omitempty"`
	Extras            *interface{}            `json:"extras,omitempty"`
}

func (ts *Tileset) ToJson() (string, error) {
	b, e := json.Marshal(ts)
	return string(b), errors.Wrap(e, "failed to convert tileset to JSON")
}

func TilesetFromJson(data io.Reader) *Tileset {
	var ts *Tileset
	json.NewDecoder(data).Decode(&ts)
	return ts
}

func ParseTilesetFile(fileName string) (*Tileset, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	return TilesetFromJson(jsonFile), nil
}
