package indexer

import (
	"fmt"
	"os"
	"testing"
)

func TestReadb3dm(t *testing.T) {
	f, _ := os.Open("example/data/data1.b3dm")
	b3d := &B3dm{}
	b3d.Read(f)
	fmt.Println(b3d)
}
