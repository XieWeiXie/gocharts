package options

/*
标题
*/
type Title struct {
	Display    bool     `json:"display,omitempty"`
	Position   string   `json:"position,omitempty"`
	Padding    int      `json:"padding,omitempty"`
	LineHeight string   `json:"lineHeight,omitempty"`
	Text       []string `json:"text,omitempty"`
	Font
}
