package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
)

type Tileset struct {
	Asset             Asset       `json:"asset"`
	Properties        *Properties `json:"properties,omitempty"`
	GeometricError    float64     `json:"geometricError"`
	Root              Tile        `json:"root"`
	ExtenstionsUsed   []string    `json:"extensionsUsed,omitempty"`
	ExtensionRequired []string    `json:"extensionsRequired,omitempty"`
	Extentions        *Extention  `json:"extensions,omitempty"`
	Extras            *Extras     `json:"extras,omitempty"`
}

type Asset struct {
	Version        string     `json:"version"`
	TilesetVersion *string    `json:"tileSetVersion,omitempty"`
	Extentions     *Extention `json:"extensions,omitempty"`
	Extras         *Extras    `json:"extras,omitempty"`
}

type Properties map[string]interface{}

type Tile struct {
	BoundingVolume      BoundingVolume  `json:"boundingVolume"`
	ViewerRequestVolume *BoundingVolume `josn:"viewerRequestVolume,omitempty"`
	GeometricError      float64         `json:"geometricError"`
	Refine              *string         `json:"refine,omitempty"`
	Transform           [16]int         `json:"transform,omitempty"`
	Content             *Content        `json:"content,omitempty"`
	Children            []Tile          `json:"children,omitempty"`
	Extentions          *Extention      `json:"extensions,omitempty"`
	Extras              *Extras         `json:"extras,omitempty"`
}

type BoundingVolume struct {
	Box        [12]float64 `json:"box,omitempty"`
	Region     [6]float64  `json:"region,omitempty"`
	Sphere     [4]float64  `json:"sphere,omitempty"`
	Extentions *Extention  `json:"extension,omitempty"`
	Extras     *Extras     `json:"extras,omitempty"`
}

type Content struct {
	BoundingVolume BoundingVolume `json:"boundingVolume,omitempty"`
	// change it to "uri" https://github.com/CesiumGS/3d-tiles/tree/main/specification#content
	URL        string     `json:"url"`
	Extentions *Extention `json:"extension,omitempty"`
	Extras     *Extras    `json:"extras,omitempty"`
}

type Extention interface{}
type Extras interface{}

type IndexesConfig struct {
	IdProperty string  `json:"idProperty"`
	Indexes    Indexes `json:"indexes"`
}

type Indexes map[string]IndexConfig

type IndexConfig struct {
	Kind string `json:"kind"`
}

func readJSONFile(fileName string) ([]byte, error) {
	fmt.Println(fileName)
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	fmt.Println("Successfully Opened ", fileName)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}

	return byteValue, nil
}

func parseJSONFile[T any](fileName string, data T) (T, error) {
	var zero T
	byteValue, err := readJSONFile(fileName)
	if err != nil {
		return zero, errors.Wrap(err, "could not read config")
	}

	errUnmarshal := json.Unmarshal([]byte(byteValue), &data)
	if errUnmarshal != nil {
		return zero, errors.Wrap(errUnmarshal, "could not unmarshal the json")
	}

	return data, nil
}

func ParseTilesetFile(fileName string) (Tileset, error) {
	return parseJSONFile(fileName, Tileset{})
}

func ParseIndexerConfigFile(fileName string) (IndexesConfig, error) {
	return parseJSONFile(fileName, IndexesConfig{})
}

func main() {

	tilesetFile := "example/tileset.json"
	tileset, err := ParseTilesetFile(tilesetFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	indexConfigFile := "example/3dtiles-config.json"
	indexesConfig, err := ParseIndexerConfigFile(indexConfigFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Tileset: ", tileset)
	fmt.Println("IndexesConfig: ", indexesConfig)

}
