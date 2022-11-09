package indexer

import (
	"fmt"
	"os"
	"testing"
)

func TestReadb3dm(t *testing.T) {
	f, err := os.Open("example/data/data1.b3dm")
	if err != nil {
		t.Errorf("failed to open the b3dm file")
	}
	
	b3d := &B3dm{}
	err1 := b3d.Read(f)
	if err1 != nil {
		t.Errorf("failed to open the b3dm file")
	}
	fmt.Println(b3d)
}
