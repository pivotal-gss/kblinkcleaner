package main


import (
	"regexp"
	"fmt"
)

var (
	// TODO IMAGE ATTACHMENTS
	linkPattern = regexp.MustCompile(`href="http\S+/discuss.pivotal.io/hc/en-us/articles/\S+"`)
	kbIDPattern = regexp.MustCompile(`hc/en-us/articles/(\d+)`)
	relativeLinkPattern = regexp.MustCompile(`href="/hc/en-us/articles/\S+"`)

	//linkPattern = regexp.MustCompile(`<.*href=\"(\S+)\"`)
	//httpSchemaPattern = regexp.MustCompile(`http\S+\/(\S+)`)
	//linkHashPattern = regexp.MustCompile(`href=\\"#`)
	//articleRootPath = "https://discuss.pivotal.io"
)

type RootObj struct {
	Articles []Article `json:"articles"`
}

type Article struct {
	ID int64 `json:"id"`
	URL string `json:"url"`
	HTMLURL string `json:"html_url"`
	AuthorID int64 `json:"author_id"`
	CommentsDisabled bool `json:"comments_disabled"`
	Draft bool `json:"draft"`
	Promoted bool `json:"promoted"`
	Position int64 `json:"position"`
	VoteSum int64 `json:"vote_sum"`
	VoteCount int64 `json:"vote_count"`
	SectionID int64 `json:"section_id"`
	CreatedAT string `json:"created_at"`
	UpdatedAT string `json:"updated_at"`
	Name string `json:"name"`
	Title string `json:"title"`
	Body string `json:"body"`
	SourceLocal string `json:"source_locale"`
	Locale string `json:"locale"`
	Outdated bool `json:"outdated"`
	OutdataedLocales []string `json:"outdated_locales"`
	LabelNames []string `json:"label_names"`
}


func (a *RootObj) Clean() {
	for i := range a.Articles {
		a.Articles[i].CleanLinks()
	}
}

func (a *Article) CleanLinks() {
	a.Body = linkPattern.ReplaceAllStringFunc(a.Body, replaceDirectLinks)
	a.Body = relativeLinkPattern.ReplaceAllStringFunc(a.Body, replaceRelativeLinks)
}

// TODO look up id maps from cache
func ZDKBMap(id string) string {
	//fmt.Printf("replacing: %s with https://the.force.yup/kb\n", id)
	return `href=\"https:\/\/the.force.yup\/kb\"`
}

func replaceDirectLinks(s string) string {
	//fmt.Println("here")
	id := kbIDPattern.FindAllStringSubmatch(s, 1)
	if len(id) <= 0 {
		fmt.Printf("Warning: could not find kbid in %s\n", s)
		return s
	}

	if len(id[0]) < 2 {
		fmt.Printf("Warning: could not find kbid in %s from array %v\n", s, id)
		return s
	}
	return ZDKBMap(id[0][1])
}

func replaceRelativeLinks(s string) string {
	//fmt.Println("rel here")
	id := kbIDPattern.FindAllStringSubmatch(s, 1)
	if len(id) <= 0 {
		fmt.Printf("Warning: could not find kbid in %s\n", s)
		return s
	}

	if len(id[0]) < 2 {
		fmt.Printf("Warning: could not find kbid in %s from array %v\n", s, id)
		return s
	}
	return ZDKBMap(id[0][1])
}