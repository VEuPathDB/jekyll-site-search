package solr

import (
	"encoding/json"
	"github.com/VEuPathDB/jekyll-site-search/lib/util"
	"os"
)

const (
	outPath      = "./api/v1"
	documentFile = outPath + "/solr.json"
	batchFile    = outPath + "/batch.json"
	doneFile     = outPath + "/DONE"
)

func WriteDocumentJson(out DocumentCollection) {
	ensurePath()
	writeJson(documentFile, convertDoc(out))
}

func WriteBatchJson(batch *Batch) {
	ensurePath()
	writeJson(batchFile, []*Batch{batch})
}

func WriteDoneFile() {
	file, err := os.OpenFile(doneFile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	util.Require(file.Close())
}

func writeJson(path string, out interface{}) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	util.Require(enc.Encode(out))
	util.Require(file.Close())
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
