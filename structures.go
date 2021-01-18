package tvmaze

// ResponseData doc
type ResponseData []showWrap

type showWrap struct {
	Score float64 `json:"score"`
	Show  show    `json:"show"`
}

type show struct {
	ID           int         `json:"id"`
	URL          string      `json:"url"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Language     string      `json:"language"`
	Genres       []string    `json:"genres"`
	Status       string      `json:"status"`
	Runtime      int         `json:"runtime"`
	Premiered    string      `json:"premiered"`
	OfficialSite string      `json:"officialSite"`
	Schedule     schedule    `json:"schedule"`
	Rating       rating      `json:"rating"`
	Weight       int         `json:"weight"`
	Network      network     `json:"network"`
	WebChannel   interface{} `json:"webChannel"`
	Externals    externals   `json:"externals"`
	Image        image       `json:"image"`
	Summary      string      `json:"summary"`
	Updated      int         `json:"updated"`
	Links        links       `json:"_links"`
}

type schedule struct {
	Time string   `json:"time"`
	Days []string `json:"days"`
}

type rating struct {
	Average float64 `json:"average"`
}

type network struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Country country `json:"country"`
}

type country struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Timezone string `json:"timezone"`
}

type externals struct {
	Tvrage  int    `json:"tvrage"`
	Thetvdb int    `json:"thetvdb"`
	Imdb    string `json:"imdb"`
}

type image struct {
	Medium   string `json:"medium"`
	Original string `json:"original"`
}

type links struct {
	Self            link `json:"self"`
	Previousepisode link `json:"previousepisode"`
}

type link struct {
	Href string `json:"href"`
}
