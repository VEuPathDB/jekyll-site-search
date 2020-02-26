package main

import (
	"github.com/VEuPathDB/jekyll-site-search/lib/convert"
	"github.com/VEuPathDB/jekyll-site-search/lib/jekyll"
	"github.com/VEuPathDB/jekyll-site-search/lib/solr"
)

func main() {
	batch := solr.NewBatch()
	solr.WriteDocumentJson(convert.PagesToDocs(jekyll.ReadPageJson(), batch))
	solr.WriteBatchJson(batch)
}
