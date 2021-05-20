package convert

import (
	"strings"

	"github.com/VEuPathDB/jekyll-site-search/lib/jekyll"
	"github.com/VEuPathDB/jekyll-site-search/lib/solr"
)

func PagesToDocs(
	pages jekyll.Pages,
	batch *solr.Batch,
) solr.DocumentCollection {
	out := make([]solr.Document, 0, len(pages))
	for i := range pages {
		pageToDocs(pages[i], batch, &out)
	}
	return out
}

func pageToDocs(
	page *jekyll.Page,
	batch *solr.Batch,
	docs *[]solr.Document,
) {
	tag, ok := page.IsUsable()

	if !ok {
		return
	}

	splitUrl := strings.Split(strings.Trim(page.Link, "/"), "/")

	for _, project := range parseProjects(splitUrl[0], page.Header.Categories) {
		var id string
		var url []string

		if project == "" {
			id = tag + ":all:" + strings.Join(splitUrl, ":")
		} else if project == splitUrl[0] {
			id = tag + ":" + strings.Join(splitUrl, ":")
		} else {
			id = tag + ":" + project + ":" + strings.Join(splitUrl, ":")
		}

		if page.PrependContent {
			url = append([]string{"content"}, splitUrl...)
		} else {
			url = splitUrl
		}

		*docs = append(*docs, solr.Document{
			Batch:   solr.Batch{
				BatchType: batch.BatchType,
				BatchId:   batch.BatchId,
				BatchTime: batch.BatchTime,
				BatchName: batch.BatchName,
				Id:        id,
			},
			Title:   page.Title,
			Url:     url,
			Type:    tag,
			Body:    page.Content,
			Project: project,
		})
	}
}

func parseProjects(url string, cats []string) (out []string) {
	// Prioritize categories
	for _, cat := range cats {
		if _, ok := projects[strings.ToLower(cat)]; ok {
			out = append(out, cat)
		}
	}
	
	if len(out) > 0 {
		return
	}
	
	if _, ok := projects[strings.ToLower(url)]; ok {
		out = append(out, url)
		return
	}

	out = append(out, "")
	return
}