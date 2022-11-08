package indexer

type Features map[string]FeatureValues

type FeatureValues struct {
	Position   interface{}
	Properties interface{}
}