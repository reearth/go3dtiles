package indexer

type featureTableDecode func(header map[string]interface{}, buff []byte) map[string]interface{}
type featureTableEncode func(header map[string]interface{}, data map[string]interface{}) []byte

type FeatureTable struct {
	Header map[string]interface{}
	Data   map[string]interface{}

	decode featureTableDecode
	encode featureTableEncode
}
