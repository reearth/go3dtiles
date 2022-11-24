package b3dm

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TESTFILE_LL_B3DM     = "testdata/ll.b3dm"
	TESTFILE_PARENT_B3DM = "testdata/parent.b3dm"
)

func TestOpen(t *testing.T) {
	tests := []struct {
		name    string
		want    *B3dm
		wantErr bool
	}{
		{"openError", nil, true},
		{TESTFILE_LL_B3DM, &B3dm{
			Header: B3dmHeader{
				Magic:                        [4]byte{98, 51, 100, 109},
				Version:                      1,
				ByteLength:                   9700,
				FeatureTableJSONByteLength:   92,
				FeatureTableBinaryByteLength: 0,
				BatchTableJSONByteLength:     640,
				BatchTableBinaryByteLength:   0,
			},
			FeatureTable: FeatureTable{
				Header: map[string]interface{}{
					"BATCH_LENGTH": float64(10),
					"RTC_CENTER":   []interface{}{1214914.5525041146, -4736388.031625768, 4081548.0407588882},
				},
				decode: B3dmFeatureTableDecode,
			},
			BatchTable: BatchTable{
				Header: map[string]interface{}{
					"Height": []interface{}{
						11.721514919772744, 12.778013898059726, 9.500697679817677,
						8.181250356137753, 10.231159372255206, 12.68863015063107,
						6.161747192963958, 7.122806219384074, 12.393268510699272,
						11.431036269292235,
					},
					"Latitude": []interface{}{
						0.6988582109, 0.6988621128191176, 0.698870582386204, 0.6988575056044288,
						0.6988603596248432, 0.6988530761634713, 0.6988687144359211,
						0.6988698975892317, 0.6988569944876143, 0.6988651780819983,
					},
					"Longitude": []interface{}{
						-1.3197004795898053, -1.3197036065852055, -1.319708772296242,
						-1.3197052536661238, -1.3197012996975566, -1.3197180493677987,
						-1.3197058762367364, -1.3196853243969762, -1.3196881546957797,
						-1.3197161145487923,
					},
					"id": []interface{}{
						float64(0), float64(1), float64(2),
						float64(3), float64(4), float64(5),
						float64(6), float64(7), float64(8),
						float64(9)},
				},
				Data: map[string][]interface{}{
					"Height": {
						11.721514919772744, 12.778013898059726, 9.500697679817677,
						8.181250356137753, 10.231159372255206, 12.68863015063107,
						6.161747192963958, 7.122806219384074, 12.393268510699272,
						11.431036269292235,
					},
					"Latitude": {
						0.6988582109, 0.6988621128191176, 0.698870582386204, 0.6988575056044288,
						0.6988603596248432, 0.6988530761634713, 0.6988687144359211,
						0.6988698975892317, 0.6988569944876143, 0.6988651780819983,
					},
					"Longitude": {
						-1.3197004795898053, -1.3197036065852055, -1.319708772296242,
						-1.3197052536661238, -1.3197012996975566, -1.3197180493677987,
						-1.3197058762367364, -1.3196853243969762, -1.3196881546957797,
						-1.3197161145487923,
					},
					"id": {
						float64(0), float64(1), float64(2),
						float64(3), float64(4), float64(5),
						float64(6), float64(7), float64(8),
						float64(9),
					},
				},
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

			// Not asserting the gltf object as its already been tested in "github.com/qmuntal/gltf"
			if tt.wantErr != true {
				assert.Equal(t, got.BatchTable, tt.want.BatchTable, "Open().BatchTable = False")
				assert.Equal(t, got.Header, tt.want.Header, "Open().HeaderTable = False")
				assert.Equal(t, got.FeatureTable.Header, tt.want.FeatureTable.Header, "Open().FeatureTable = False")
				assert.Equal(t, got.FeatureTable.Data, tt.want.FeatureTable.Data, "Open().FeatureTable = False")
			}
		})
	}
}

func TestB3dm_Decode(t *testing.T) {
	f, _ := os.Open(TESTFILE_PARENT_B3DM)
	defer f.Close()
	type args struct {
		b3dm *B3dm
	}
	tests := []struct {
		name    string
		d       *B3dmReader
		args    args
		wantErr bool
	}{
		{"empty", NewB3dmReader(bytes.NewBufferString("")), args{new(B3dm)}, true},
		{"invalidJSON", NewB3dmReader(bytes.NewBufferString("{asset: {}}")), args{new(B3dm)}, true},
		{"invalidBuffer", NewB3dmReader(bytes.NewBufferString("{\"buffers\": [{\"byteLength\": 0}]}")), args{new(B3dm)}, true},
		{"correctb3dm", NewB3dmReader(f), args{new(B3dm)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.Decode(tt.args.b3dm); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFeatureTableView(t *testing.T) {
	tests := []struct {
		name    string
		want    *B3dmFeatureTable
		wantErr bool
	}{
		{"openError", nil, true},
		{TESTFILE_LL_B3DM, &B3dmFeatureTable{
			BatchLength: 10,
			RtcCenter: [3]float64{
				1.2149145525041146e+06,
				-4.736388031625768e+06,
				4.0815480407588882e+06,
			},
		}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b3dm, err := Open(tt.name)
			if !tt.wantErr {
				if assert.NoError(t, err) {
					got := b3dm.GetFeatureTableView()
					assert.Equal(t, got, tt.want, "getFeatureTableView() = false")
				}
			}
		})
	}
}
