package tileset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TESTFILE_TILESET = "tileset_testdata/tileset.json"

func TestOpen(t *testing.T) {
	tests := []struct {
		name    string
		want    *Tileset
		wantErr bool
	}{
		{"openError", nil, true},
		{TESTFILE_TILESET, &Tileset{
			Asset: Asset{
				Version:        "1.0",
				TilesetVersion: "1.2.3",
			},
			Properties: map[string]Schema{
				"Height": {
					Maximum: 85.41026367992163,
					Minimum: 6.161747192963958,
				},
				"Latitude": {
					Maximum: 0.6989046192460953,
					Minimum: 0.698848878034009,
				},
				"Longitude": {
					Maximum: -1.319644104024109,
					Minimum: -1.3197192952275933,
				},
				"id": {
					Maximum: float64(9),
					Minimum: float64(0),
				},
			},
			GeometricError: float64(240),
			Root: Tile{
				BoundingVolume: BoundingVolume{
					Region: &[6]float64{
						-1.3197209591796106, 0.6988424218,
						-1.3196390408203893, 0.6989055782,
						0, 88,
					},
				},
				GeometricError: float64(70),
				Refine:         TILE_REFINE_ADD,
				Content: &Content{
					BoundingVolume: BoundingVolume{
						Region: &[6]float64{
							-1.3197004795898053, 0.6988582109,
							-1.3196595204101946, 0.6988897891,
							0, 88,
						},
					},
					URL: "",
				},
				Children: &[]Tile{
					{
						BoundingVolume: BoundingVolume{
							Region: &[6]float64{
								-1.3197209591796106, 0.6988424218, -1.31968, 0.698874, 0, 20,
							},
						},
						GeometricError: float64(0),
						Content: &Content{
							URL: "",
						},
					},
					{
						BoundingVolume: BoundingVolume{
							Region: &[6]float64{
								-1.31968, 0.6988424218, -1.3196390408203893, 0.698874, 0, 20,
							},
						},
						GeometricError: float64(0),
						Content: &Content{
							URL: "",
						},
						Extras: map[string]interface{}{
							"id": "Special Tile",
						},
					},
					{
						BoundingVolume: BoundingVolume{
							Region: &[6]float64{
								-1.31968, 0.698874, -1.3196390408203893, 0.6989055782, 0, 20,
							},
						},
						GeometricError: float64(0),
						Content: &Content{
							URL: "",
						},
					},
					{
						BoundingVolume: BoundingVolume{
							Region: &[6]float64{
								-1.3197209591796106, 0.698874, -1.31968, 0.6989055782, 0, 20,
							},
						},
						GeometricError: float64(0),
						Content: &Content{
							URL: "",
						},
					},
				},
			},
			Extras: map[string]interface{}{
				"name": "Sample Tileset",
			},
		}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Open(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, got, tt.want, "Open() = False")
		})
	}
}
