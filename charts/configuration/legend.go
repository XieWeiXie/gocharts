package options

/*
图例
*/

type Legend struct {
	Display   bool         `json:"display"`
	Position  Position     `json:"position"`
	FullWidth bool         `json:"fullWidth"`
	OnClick   func()       `json:"onClick"`
	OnHover   func()       `json:"onHover"`
	OnLeave   func()       `json:"onLeave"`
	Reverse   bool         `json:"reverse"`
	Labels    LegendObject `json:"labels"`
}

func familyAvaliableOption() []string {
	return []string{"Helvetica Neue", "Helvetica", "Arial", "sans-serif"}
}

type LegendObject struct {
	BoxWidth       int    `json:"boxWidth"`
	FontSize       int    `json:"fontSize"`
	FontStyle      string `json:"fontStyle"`
	FontColor      string `json:"fontColor"` // #666 or rgb(255,99,132)
	FontFamily     string `json:"fontFamily"`
	Padding        int    `json:"padding"`
	GenerateLabels func() `json:"generateLabels"`
	Filter         func() `json:"filter"`
	UsePointStyle  bool   `json:"usePointStyle"`
}
