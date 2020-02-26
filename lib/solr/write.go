package solr

import (
	"encoding/json"
	"io/ioutil"
)

const (
	documentFile = "_site/api/v1/solr.json"
	batchFile    = "_site/api/v1/batch.json"
)

func WriteDocumentJson(out DocumentCollection) {
	writeJson(documentFile, convertDoc(out))
}

func WriteBatchJson(batch *Batch) {
	writeJson(batchFile, []*Batch{batch})
}

func writeJson(path string, out interface{}) {
	data, err := json.Marshal(out)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, data, 0644)

	if err != nil {
		panic(err)
	}
}

func convertDoc(in DocumentCollection) []map[string]interface{} {
	out := make([]map[string]interface{}, len(in))

	for i := range in {
		out[i] = in[i].ToMap()
	}

	return out
}
