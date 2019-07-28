package model

type Feed struct {
	Title   string `xml:"title"`
	ID      string `xml:"id"`
	Updated string `xml:"updated"`
	Entry   []Entry `xml:"entry"`
}

type Entry struct {
	ID        string `xml:"id"`
	Title     string `xml:"title"`
	Summary   string `xml:"summary"`
	Published string `xml:"published"`
	Updated   string `xml:"updated"`
	Author    struct {
		Name   string `xml:"name"`
		URI    string `xml:"uri"`
		Avatar string `xml:"avatar"`
	} `xml:"author"`
	Link struct {
		Rel  string `xml:"rel,attr"`
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Diggs    string `xml:"diggs"`
	Views    string `xml:"views"`
	Comments string `xml:"comments"`
}
