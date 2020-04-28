package jekyll

import (
	"encoding/json"
	"github.com/VEuPathDB/jekyll-site-search/lib/util"
	"strings"
)

//==========================================================
//
// Internal Statics & Constants
//
//==========================================================

// These values must exist in the document type enum config at
// https://github.com/VEuPathDB/SolrDeployment/blob/master/configsets/site-search/conf/enumsConfig.xml
var validTags = map[string]bool{
	"general":           true,
	"tutorial":          true,
	"news":              true,
	"workshop-exercise": true,
}

//==========================================================
//
// Exported Types
//
//==========================================================

type Pages = []*Page

type Header struct {
	Title      string   `json:"title,omitempty"`
	Tags       []string `json:"tags"`
	Categories []string `json:"categories"`
	Link       string   `json:"permalink"`
}

type Page struct {
	Header `json:"header"`

	Content string `json:"output"`
}

//==========================================================
//
// Exported Functions
//
//==========================================================

func NewPage(b []byte) *Page {
	out := new(Page)
	util.Require(json.Unmarshal(b, out))
	return out
}

func (p *Page) IsUsable() (string, bool) {
	if len(p.Content) == 0 || len(p.Link) == 0 {
		return "", false
	}

	for i := range p.Tags {
		tmp := strings.ToLower(p.Tags[i])
		if _, ok := validTags[tmp]; ok {
			if p.Content == "-" {
				p.Content = ""
			}
			return tmp, true
		}
	}

	return "", false
}
