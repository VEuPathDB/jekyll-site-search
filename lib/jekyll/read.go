package jekyll

import (
	"bufio"
	"github.com/VEuPathDB/jekyll-site-search/lib/html"
	"io"
	"os"
)

func ParseInput() []*Page {
	reader := bufio.NewReader(os.Stdin)
	scrub  := html.NewScrubber()
	pages  := make([]*Page, 0, 1024)
	for {
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		page := NewPage(bytes)
		preLn := len(page.Content)
		page.Content = scrub.Scrub(page.Content)
		if preLn > 0 && len(page.Content) == 0 {
			page.Content = "-"
		}
		pages = append(pages, page)
	}

	return pages
}
