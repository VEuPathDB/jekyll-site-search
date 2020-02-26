package solr

import (
	"encoding/json"
	"io/ioutil"
)

const outputFile = "_site/api/v1/solr.json"

func WriteDocumentJson(out DocumentCollection) {
	data, err := json.Marshal(convert(out))

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(outputFile, data, 0640)

	if err != nil {
		panic(err)
	}
}


func convert(in DocumentCollection) []map[string]interface{} {
	out := make([]map[string]interface{}, len(in))

	for i := range in {
		out[i] = in[i].ToMap()
	}

	return out
}
