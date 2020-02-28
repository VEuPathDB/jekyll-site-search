package html

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"regexp"
	"strings"
)

type Scrubber interface {
	Scrub(string) string
}

func NewScrubber() Scrubber {
	return &scrubber{true}
}

// Tags that should be ignored when parsing page content.
var ignored = map[atom.Atom]bool{
	atom.Area:     false,
	atom.Base:     false,
	atom.Canvas:   false,
	atom.Head:     false,
	atom.Html:     false,
	atom.Link:     false,
	atom.Meta:     false,
	atom.Noscript: false,
	atom.Script:   false,
	atom.Style:    false,
	atom.Title:    false,
}

type scrubber struct {
	armed bool
}

var (
	stripTags  = regexp.MustCompile("</?\\w{1,3} */?>")
	stripChars = regexp.MustCompile("[\":;><.,()&]")
	stripSpace = regexp.MustCompile(" {2,}|\t+|[ \t]*(\r\n|\r|\n)")
)

func (s *scrubber) Scrub(in string) string {
	tok := html.NewTokenizer(strings.NewReader(html.UnescapeString(in)))
	out := strings.Builder{}

	for {
		n := tok.Next()

		if n == html.ErrorToken {
			if tok.Err() == io.EOF {
				break
			}
			panic(tok.Err())
		}

		switch n {
		case html.StartTagToken:
			s.arm(tok.Token())
		case html.EndTagToken:
			s.disarm(tok.Token())
		case html.TextToken:
			if !s.armed {
				continue
			}
			out.WriteString(cleanString(tok.Token().String()))
		}
	}

	// Reset for reuse
	s.armed = true

	return out.String()
}

func (s *scrubber) arm(t html.Token) {
	if _, ok := ignored[t.DataAtom]; ok {
		s.armed = false
	} else {
		s.armed = true
	}
}

func (s *scrubber) disarm(t html.Token) {
	if _, ok := ignored[t.DataAtom]; ok {
		s.armed = true
	} else {
		s.armed = false
	}
}

func cleanString(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return ""
	}

	return strings.TrimSpace(stripSpace.ReplaceAllString(
		stripChars.ReplaceAllString(
			stripTags.ReplaceAllString(
				html.UnescapeString(s),
				""),
			""),
		" ")) + " "
}
