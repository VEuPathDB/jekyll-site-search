package solr

import "reflect"

type DocumentCollection = []Document

type Document struct {
	Id        string   `solr:"id"`
	Title     string   `solr:"hyperlinkName"`
	Url       []string `solr:"primaryKey"`
	Type      string   `solr:"document-type"`
	Body      string   `solr:"body"`
	BatchType string   `solr:"batch-type"`
	BatchName string   `solr:"batch-name"`
	BatchId   string   `solr:"batch-id"`
	BatchTime int64    `solr:"batch-timestamp"`
}

func (d *Document) ToMap() map[string]interface{} {
	v := reflect.ValueOf(d).Elem()
	t := v.Type()
	l := t.NumField()
	tags := make(map[string]string, l)
	out := make(map[string]interface{}, l)

	for i := 0; i < l; i++ {
		tags[t.Field(i).Name] = getJsonKey(d, t.Field(i))
	}

	if len(d.Title) > 0 {
		out[tags["Title"]] = d.Title
	}

	out[tags["Url"]] = d.Url
	out[tags["BatchType"]] = d.BatchType
	out[tags["BatchName"]] = d.BatchName
	out[tags["BatchTime"]] = d.BatchTime
	out[tags["BatchId"]] = d.BatchId
	out[tags["Body"]] = d.Body
	out[tags["Type"]] = d.Type
	out[tags["Id"]] = d.Id

	return out
}

const (
	bodyPrefixKey = "TEXT__"
	bodySuffixKey = "_content"
)

func getJsonKey(d *Document, f reflect.StructField) string {
	if f.Name == "Body" {
		return bodyPrefixKey + d.Type + bodySuffixKey
	}
	return f.Tag.Get("solr")
}
