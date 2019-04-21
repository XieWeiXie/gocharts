package options

type Scales struct {
	XAxes []Axes `json:"xAxes,omitempty"`
	YAxes []Axes `json:"yAxes,omitempty"`
}

type scaleLabel map[string]interface{}
type ticks map[string]interface{}

type Axes struct {
	Display    bool       `json:"display,omitempty"`
	ScaleLabel scaleLabel `json:"scaleLabel,omitempty"`
	Ticks      ticks      `json:"ticks,omitempty"`
	Type       string     `json:"type,omitempty"`
	Position   string     `json:"position,omitempty"`
	ID         string     `json:"id,omitempty"`
}

func NewAxes() *Axes {
	return &Axes{
		ScaleLabel: make(map[string]interface{}),
		Ticks:      make(map[string]interface{}),
	}
}

func (A *Axes) AddLabel(display bool, values string) {
	A.ScaleLabel["display"] = display
	A.ScaleLabel["labelString"] = values
}

func (A *Axes) AddTicks(key, beginAtZero interface{}) {
	A.Ticks[key.(string)] = beginAtZero
}
