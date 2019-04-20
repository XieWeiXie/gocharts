package configuration

type Elements struct {
	PointOfElements
	LineOfElements
	RectangleElements
	ArcElements
}

type PointOfElements struct {
	Radius           int    `json:"radius"`
	PointStyle       string `json:"pointStyle"`
	style            PointOfElements
	Rotation         int    `json:"rotation"`
	BackgroundColor  string `json:"backgroundColor"`
	backColor        Color
	BorderWidth      int    `json:"borderWidth"`
	BorderColor      string `json:"borderColor"`
	borderColor      Color
	HitRadius        int `json:"hitRadius"`
	HoverRadius      int `json:"hoverRadius"`
	HoverBorderWidth int `json:"hoverBorderWidth"`
}

type PointOfElementStyle struct {
}

func (p PointOfElementStyle) Circle() string {
	return "circle"
}

func (p PointOfElementStyle) Cross() string {
	return "cross"
}
func (p PointOfElementStyle) CrossRot() string {
	return "crossRot"
}

func (p PointOfElementStyle) Dash() string {
	return "dash"
}

func (p PointOfElementStyle) Line() string {
	return "line"
}
func (p PointOfElementStyle) Rect() string {
	return "rect"
}

func (p PointOfElementStyle) RectRounded() string {
	return "rectRounded"
}

func (p PointOfElementStyle) RectRot() string {
	return "rectRot"
}

func (p PointOfElementStyle) Star() string {
	return "star"
}

func (p PointOfElementStyle) Triangle() string {
	return "triangle"
}

type LineOfElements struct {
	Tension          int    `json:"tension"`
	BackgroundColor  string `json:"backgroundColor"`
	backColor        Color
	BorderWidth      int    `json:"borderWidth"`
	BorderColor      string `json:"borderColor"`
	borderColor      Color
	BorderCapStyle   string `json:"borderCapStyle"`
	BorderDash       []int  `json:"borderDash"`
	BorderDashOffset int    `json:"borderDashOffset"`
	BorderJoinStyle  string `json:"borderJoinStyle"`
	CapBezierPoints  bool   `json:"capBezierPoints"`
	Fill             bool   `json:"fill"`
	Stepped          bool   `json:"stepped"`
}

type RectangleElements struct {
	BackgroundColor string `json:"backgroundColor"`
	backColor       Color
	BorderWidth     int    `json:"borderWidth"`
	BorderColor     string `json:"borderColor"`
	borderColor     Color
	BorderSkipped   string `json:"borderSkipped"`
}

type ArcElements struct {
	BackgroundColor string `json:"backgroundColor"`
	backColor       Color
	BorderWidth     int    `json:"borderWidth"`
	BorderColor     string `json:"borderColor"`
	borderColor     Color
	BorderAlign     string `json:"borderAlign"`
}
