package solr

import (
	"encoding/json"
	"github.com/VEuPathDB/jekyll-site-search/lib/util"
	"io/ioutil"
	"os"
)

const (
	outPath      = "api/v1"
	documentFile = outPath + "/solr.json"
	batchFile    = outPath + "/batch.json"
)

func WriteDocumentJson(out DocumentCollection) {
	ensurePath()
	writeJson(documentFile, convertDoc(out))
}

func WriteBatchJson(batch *Batch) {
	ensurePath()
	writeJson(batchFile, []*Batch{batch})
}

func writeJson(path string, out interface{}) {
	data := util.MustRead(json.Marshal(out))
	util.Require(ioutil.WriteFile(path, data, 0644))
}

func ensurePath() {
	if !dirExists(outPath) {
		util.Require(os.MkdirAll(outPath, 0755))
	}
}

func convertDoc(in DocumentCollection) []map[string]interface{} {
	out := make([]map[string]interface{}, len(in))

	for i := range in {
		out[i] = in[i].ToMap()
	}

	return out
}

func dirExists(path string) bool {
	out, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	if !out.IsDir() {
		panic(outPath + " exists and is not a directory")
	}
	return true
}
