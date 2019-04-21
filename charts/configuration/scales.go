package options

type Scales struct {
	XAxes []Axes `json:"xAxes"`
	YAxes []Axes `json:"yAxes"`
}

type scaleLabel map[string]interface{}
type ticks map[string]interface{}

type Axes struct {
	Display    bool       `json:"display"`
	ScaleLabel scaleLabel `json:"scaleLabel"`
	Ticks      ticks      `json:"ticks"`
}

func (A *Axes) AddLabel(display bool, values string) {
	A.ScaleLabel["display"] = display
	A.ScaleLabel["labelString"] = values
}

func (A *Axes) AddTicks(key, beginAtZero interface{}) {
	A.Ticks[key.(string)] = beginAtZero
}
