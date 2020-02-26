package convert

import (
	"fmt"
	"github.com/VEuPathDB/jekyll-site-search/lib/jekyll"
	"github.com/VEuPathDB/jekyll-site-search/lib/solr"
	"github.com/VEuPathDB/jekyll-site-search/lib/util"
	"strings"
	"time"
)

const (
	batchType = "jekyll"
	batchName = "all"
)

// These values must match the document type enum config at
// https://github.com/VEuPathDB/SolrDeployment/blob/master/configsets/site-search/conf/enumsConfig.xml
var validTags = map[string]bool{
	"general":           true,
	"tutorial":          true,
	"workshop-exercise": true,
}

func PagesToDocs(pages *jekyll.Pages) solr.DocumentCollection {
	out := make([]solr.Document, 0, len(pages.Entries))
	for i := range pages.Entries {
		tmp, ok := pageToDoc(&(pages.Entries[i]))
		if ok {
			out = append(out, tmp)
		}
	}
	return out
}

func pageToDoc(entry *jekyll.PageEntry) (out solr.Document, ok bool) {
	tag, ok := tagsToDoctype(entry.Tags)

	if !ok {
		return
	}

	// URL is required
	if len(entry.Url) == 0 {
		return out, false
	}

	now := timeToMillis(time.Now())

	out.BatchTime = now
	out.BatchId = genBatchId(now)
	out.BatchName = batchName
	out.BatchType = batchType

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

func timeToMillis(now time.Time) int64 {
	return now.UnixNano() / int64(time.Millisecond)
}

func genBatchId(millis int64) string {
	return fmt.Sprintf("%s_%s_%d", batchType, batchName, millis)
}
