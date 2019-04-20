package configuration

/*
标题
*/
type Title struct {
	Display    bool     `json:"display"`
	Position   string   `json:"position"`
	Padding    int      `json:"padding"`
	LineHeight string   `json:"lineHeight"`
	Text       []string `json:"text"`
	Font       Font
}
