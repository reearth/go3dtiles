package tileset

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	TILE_REFINE_ADD     = "ADD"
	TILE_REFINE_REPLACE = "REPLACE"
)

type Tileset struct {
	Asset             Asset                  `json:"asset"`
	Properties        map[string]Schema      `json:"properties,omitempty"`
	GeometricError    float64                `json:"geometricError"`
	Root              Tile                   `json:"root"`
	ExtenstionsUsed   *[]string              `json:"extensionsUsed,omitempty"`
	ExtensionRequired *[]string              `json:"extensionsRequired,omitempty"`
	Extentions        map[string]interface{} `json:"extensions,omitempty"`
	Extras            interface{}            `json:"extras,omitempty"`
}

type Tile struct {
	BoundingVolume      BoundingVolume         `json:"boundingVolume"`
	ViewerRequestVolume *BoundingVolume        `josn:"viewerRequestVolume,omitempty"`
	GeometricError      float64                `json:"geometricError"`
	Refine              string                 `json:"refine,omitempty"`
	Transform           *[16]float64           `json:"transform,omitempty"`
	Content             *Content               `json:"content,omitempty"`
	Children            *[]Tile                `json:"children,omitempty"`
	Extentions          map[string]interface{} `json:"extensions,omitempty"`
	Extras              interface{}            `json:"extras,omitempty"`
}

type Content struct {
	BoundingVolume BoundingVolume          `json:"boundingVolume,omitempty"`
	URL            string                  `json:"url,omitempty"`
	URI            string                  `json:"uri,omitempty"`
	Extentions     *map[string]interface{} `json:"extension,omitempty"`
	Extras         *interface{}            `json:"extras,omitempty"`
}

type BoundingVolume struct {
	Box        *[12]float64           `json:"box,omitempty"`
	Region     *[6]float64            `json:"region,omitempty"`
	Sphere     *[4]float64            `json:"sphere,omitempty"`
	Extentions map[string]interface{} `json:"extension,omitempty"`
	Extras     interface{}            `json:"extras,omitempty"`
}

type Schema struct {
	Maximum float64 `json:"maximum,omitempty"`
	Minimum float64 `json:"minimum,omitempty"`
}

type Asset struct {
	Version        string                 `json:"version"`
	TilesetVersion string                 `json:"tileSetVersion,omitempty"`
	GltfUpAxis     string                 `json:"gltfUpAxis,omitempty"`
	Extentions     map[string]interface{} `json:"extensions,omitempty"`
	Extras         interface{}            `json:"extras,omitempty"`
}

type TilesetReader struct {
	rs io.Reader
}

func (t *Tile) Uri() (string, error) {
	content := t.Content
	if content != nil {
		if content.URI != "" {
			return content.URI, nil
		}
		if content.URL != "" {
			return content.URL, nil
		} else {
			return "", errors.New("neither URL nor URI exists for this content")
		}
	}
	return "", errors.New("content does not exist")
}

func NewTilsetReader(r io.Reader) *TilesetReader {
	return &TilesetReader{rs: r}
}

func (r *TilesetReader) Decode(ts *Tileset) error {
	if err := json.NewDecoder(r.rs).Decode(&ts); err != nil {
		return fmt.Errorf("failed to decode the JSON tilset data: %v", err)
	}
	return nil
}

func (ts *Tileset) ToJson() (string, error) {
	b, err := json.Marshal(ts)
	if err != nil {
		return "", fmt.Errorf("failed to marshal the tilset JSON: %v", err)
	}
	return string(b), nil
}

func Open(fileName string) (*Tileset, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("open failed: %v", err)
	}
	defer f.Close()
	tilsetReader := NewTilsetReader(f)
	ts := new(Tileset)
	if err := tilsetReader.Decode(ts); err != nil {
		return nil, fmt.Errorf("failed to decode the tileset: %v", err)
	}
	return ts, nil
}
