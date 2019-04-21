package charts

const (
	Line = iota
	Bar
	Radar
	Pie
	Polar
	Bubble
	Scatter
	Area
	Mixed
)

var ChartTypes map[int]string

func init() {
	ChartTypes = make(map[int]string)
	ChartTypes[Line] = "line"
	ChartTypes[Bar] = "bar"
	ChartTypes[Radar] = "radar"
	ChartTypes[Pie] = "pie"
	ChartTypes[Polar] = "polarArea"
	ChartTypes[Bubble] = "bubble"
	ChartTypes[Scatter] = "scatter"
}
