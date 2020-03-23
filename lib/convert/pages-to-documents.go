package convert

import (
	"github.com/VEuPathDB/jekyll-site-search/lib/jekyll"
	"github.com/VEuPathDB/jekyll-site-search/lib/solr"
	"strings"
)

func PagesToDocs(
	pages jekyll.Pages,
	batch *solr.Batch,
) solr.DocumentCollection {
	out := make([]solr.Document, 0, len(pages))
	for i := range pages {
		tmp, ok := pageToDoc(pages[i], batch)
		if ok {
			out = append(out, tmp)
		}
	}
	return out
}

func pageToDoc(
	page *jekyll.Page,
	batch *solr.Batch,
) (out solr.Document, ok bool) {
	tag, ok := page.IsUsable()

	out.BatchTime = batch.BatchTime
	out.BatchId = batch.BatchId
	out.BatchName = batch.BatchName
	out.BatchType = batch.BatchType

	out.Title = page.Title
	out.Url = strings.Split(strings.Trim(page.Link, "/"), "/")
	out.Body = page.Content
	out.Type = tag
	out.Id = tag + ":" + strings.Join(out.Url, ":")
	out.Project = out.Url[0]
	return
}
