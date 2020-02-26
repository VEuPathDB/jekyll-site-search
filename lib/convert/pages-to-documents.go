package convert

import (
	"github.com/VEuPathDB/jekyll-site-search/lib/jekyll"
	"github.com/VEuPathDB/jekyll-site-search/lib/solr"
	"github.com/VEuPathDB/jekyll-site-search/lib/util"
	"strings"
)

// These values must match the document type enum config at
// https://github.com/VEuPathDB/SolrDeployment/blob/master/configsets/site-search/conf/enumsConfig.xml
var validTags = map[string]bool{
	"general":           true,
	"tutorial":          true,
	"news":              true,
	"workshop-exercise": true,
}

func PagesToDocs(
	pages *jekyll.Pages,
	batch *solr.Batch,
) solr.DocumentCollection {
	out := make([]solr.Document, 0, len(pages.Entries))
	for i := range pages.Entries {
		tmp, ok := pageToDoc(&(pages.Entries[i]), batch)
		if ok {
			out = append(out, tmp)
		}
	}
	return out
}

func pageToDoc(
	entry *jekyll.PageEntry,
	batch *solr.Batch,
) (out solr.Document, ok bool) {
	tag, ok := tagsToDoctype(entry.Tags)

	if !ok {
		return
	}

	// URL is required
	if len(entry.Url) == 0 {
		return out, false
	}

	out.BatchTime = batch.BatchTime
	out.BatchId = batch.BatchId
	out.BatchName = batch.BatchName
	out.BatchType = batch.BatchType

	out.Title = entry.Title
	out.Url = strings.Split(util.Trim(entry.Url, '/'), "/")
	out.Body = entry.Body
	out.Type = tag
	out.Id = tag + ":" + strings.Join(out.Url, ":")
	return
}

func tagsToDoctype(tags []string) (string, bool) {
	for i := range tags {
		if _, ok := validTags[strings.ToLower(tags[i])]; ok {
			return tags[i], true
		}
	}
	return "", false
}
