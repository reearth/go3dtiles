package tileset

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

const (
	TILE_REFINE_ADD     = "ADD"
	TILE_REFINE_REPLACE = "REPLACE"
)

type Tileset struct {
	Asset             Asset                   `json:"asset"`
	Properties        map[string]Schema      `json:"properties,omitempty"`
	GeometricError    float64                 `json:"geometricError"`
	Root              Tile                    `json:"root"`
	ExtenstionsUsed   *[]string               `json:"extensionsUsed,omitempty"`
	ExtensionRequired *[]string               `json:"extensionsRequired,omitempty"`
	Extentions        map[string]interface{} `json:"extensions,omitempty"`
	Extras            interface{}            `json:"extras,omitempty"`
}

type Tile struct {
	BoundingVolume      BoundingVolume          `json:"boundingVolume"`
	ViewerRequestVolume *BoundingVolume         `josn:"viewerRequestVolume,omitempty"`
	GeometricError      float64                 `json:"geometricError"`
	Refine              string                 `json:"refine,omitempty"`
	Transform           *[16]float64            `json:"transform,omitempty"`
	Content             *Content                `json:"content,omitempty"`
	Children            *[]Tile                 `json:"children,omitempty"`
	Extentions          map[string]interface{} `json:"extensions,omitempty"`
	Extras              interface{}            `json:"extras,omitempty"`
}

type Content struct {
	BoundingVolume BoundingVolume `json:"boundingVolume,omitempty"`
	// change it to "uri" https://github.com/CesiumGS/3d-tiles/tree/main/specification#content
	URL        string                  `json:"url"`
	Extentions *map[string]interface{} `json:"extension,omitempty"`
	Extras     *interface{}            `json:"extras,omitempty"`
}

type BoundingVolume struct {
	Box        *[12]float64            `json:"box,omitempty"`
	Region     *[6]float64             `json:"region,omitempty"`
	Sphere     *[4]float64             `json:"sphere,omitempty"`
	Extentions map[string]interface{} `json:"extension,omitempty"`
	Extras     interface{}            `json:"extras,omitempty"`
}

type Schema struct {
	Maximum float64 `json:"maximum,omitempty"`
	Minimum float64 `json:"minimum,omitempty"`
}

type Asset struct {
	Version        string                  `json:"version"`
	TilesetVersion string                  `json:"tileSetVersion,omitempty"`
	GltfUpAxis     string                  `json:"gltfUpAxis,omitempty"`
	Extentions     map[string]interface{} `json:"extensions,omitempty"`
	Extras         interface{}            `json:"extras,omitempty"`
}

func (ts *Tileset) ToJson() (string, error) {
	b, e := json.Marshal(ts)
	return string(b), errors.Wrap(e, "failed to convert tileset to JSON")
}

func Open(fileName string) (*Tileset, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}

	var ts *Tileset
	if err = json.NewDecoder(jsonFile).Decode(&ts); err != nil {
		return nil, errors.Wrap(err, "failed to decode the tileset file")
	}
	return ts, nil
}
