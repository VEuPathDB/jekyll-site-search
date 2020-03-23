package solr

import "reflect"

type DocumentCollection = []Document

type Document struct {
	Batch
	Title   string   `json:"hyperlinkName,omitempty"`
	Url     []string `json:"primaryKey"`
	Type    string   `json:"document-type"`
	Body    string   `json:"body"`
	Project string   `json:"project"`
}

// ToMap creates a flattened map for Json marshalling with
// the keys expected by Solr.
//
// Json marshal is not used directly due to the body field
// which is dynamic based on the document type from Jekyll.
func (d *Document) ToMap() map[string]interface{} {
	l := getLn()
	tags := make(map[string]string, l)
	out := make(map[string]interface{}, l)

	appendDocumentKeys(tags, d)
	appendBatchKeys(tags)

	out[tags["Title"]] = d.Title
	out[tags["Url"]] = d.Url
	out[tags["BatchType"]] = d.BatchType
	out[tags["BatchName"]] = d.BatchName
	out[tags["BatchTime"]] = d.BatchTime
	out[tags["BatchId"]] = d.BatchId
	out[tags["Body"]] = d.Body
	out[tags["Type"]] = d.Type
	out[tags["Id"]] = d.Id
	out[tags["Project"]] = d.Project

	return out
}

const (
	bodyPrefixKey = "TEXT__"
	bodySuffixKey = "_content"
)

func appendDocumentKeys(o map[string]string, d *Document) {
	tp := reflect.TypeOf(Document{})
	ln := tp.NumField()

	for i := 0; i < ln; i++ {
		o[tp.Field(i).Name] = getDocJsonKey(d, tp.Field(i))
	}
}

func appendBatchKeys(o map[string]string) {
	tp := reflect.TypeOf(Batch{})
	ln := tp.NumField()
	for i := 0; i < ln; i++ {
		o[tp.Field(i).Name] = tp.Field(i).Tag.Get("json")
	}
}

func getLn() int {
	return reflect.TypeOf(Document{}).NumField() + reflect.TypeOf(Batch{}).NumField()
}

func getDocJsonKey(d *Document, f reflect.StructField) string {
	if f.Name == "Body" {
		return bodyPrefixKey + d.Type + bodySuffixKey
	}
	return f.Tag.Get("json")
}
