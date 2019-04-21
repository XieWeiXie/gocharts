package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/wuxiaoxiaoshen/gocharts/charts/scatter"

	"github.com/wuxiaoxiaoshen/gocharts/charts/bubble"

	"github.com/wuxiaoxiaoshen/gocharts/charts/polar"

	"github.com/wuxiaoxiaoshen/gocharts/charts/pie"

	"github.com/wuxiaoxiaoshen/gocharts/charts/radar"

	"github.com/wuxiaoxiaoshen/gocharts/charts/bar"

	"github.com/wuxiaoxiaoshen/gocharts/charts/options"

	"github.com/wuxiaoxiaoshen/gocharts/charts/line"
)

func main() {

	line := exampleLine()
	bar := exampleBar()
	radar := exampleRadar()
	pie := examplePie()
	polar := examplePolar()
	bubble := exampleBubble()
	scatter := exampleScatter()
	http.HandleFunc(fmt.Sprintf("/%s", line.TypeName()), line.Plot())
	http.HandleFunc(fmt.Sprintf("/%s", bar.TypeName()), bar.Plot())
	http.HandleFunc(fmt.Sprintf("/%s", radar.TypeName()), radar.Plot())
	http.HandleFunc(fmt.Sprintf("/%s", pie.TypeName()), pie.Plot())
	http.HandleFunc(fmt.Sprintf("/%s", polar.TypeName()), polar.Plot())
	http.HandleFunc(fmt.Sprintf("/%s", bubble.TypeName()), bubble.Plot())
	http.HandleFunc(fmt.Sprintf("/%s", scatter.TypeName()), scatter.Plot())
	http.ListenAndServe(":8888", nil)

}

func exampleLine() *line.Line {
	chart := line.ChartLine([]string{"折线图示例"})
	chart.SetLabels([]interface{}{"Red", "Blue", "Yellow", "Green", "Purple", "Orange"})
	chart.AddDataSet(*chart.NewDataSet("# of votes", []interface{}{12, 19, 3, 5, 2, 3}))
	yAes := options.NewAxes()
	yAes.AddTicks("beginAtZero", "true")
	scales := options.Scales{
		YAxes: []options.Axes{*yAes},
	}
	chart.AddOptions("scales", scales)
	return chart
}

func exampleBar() *bar.Bar {
	chart := bar.ChartBar([]string{"柱状图示例"})
	chart.SetLabels([]interface{}{"Red", "Blue", "Yellow", "Green", "Purple", "Orange"})
	chart.AddDataSet(*chart.NewDataSet("# Bar example", []interface{}{12, 19, 3, 5, 2, 3}))
	yAes := options.NewAxes()
	yAes.AddTicks("beginAtZero", "true")
	scales := options.Scales{
		YAxes: []options.Axes{*yAes},
	}
	chart.AddOptions("scales", scales)
	return chart

}

func exampleRadar() *radar.Radar {
	chart := radar.ChartRadar([]string{"雷达图示例"})
	chart.SetLabels([]interface{}{[]string{"Eating", "Dinner"}, []string{"Drinking", "Water"}, "Sleeping", []string{"Designing", "Graphics"}, "Coding", "Cycling", "Running"})
	chart.AddDataSet(*chart.NewDataSet("# Radar example", []interface{}{12, 13, 19, 4, 5, 9}))
	yAes := options.NewAxes()
	yAes.AddTicks("beginAtZero", "true")
	scales := options.Scales{
		YAxes: []options.Axes{*yAes},
	}
	chart.AddOptions("scales", scales)
	return chart
}

func examplePie() *pie.Pie {
	chart := pie.ChartPie([]string{"饼图示例"})
	chart.SetLabels([]interface{}{"Red", "Orange", "Yellow", "Green", "Blue", "HaHa"})
	chart.AddDataSet(*chart.NewDataSet("# DataSet 1", []interface{}{12, 13, 19, 4, 4, 20}))
	yAes := options.NewAxes()
	yAes.AddTicks("beginAtZero", "true")
	scales := options.Scales{
		YAxes: []options.Axes{*yAes},
	}
	chart.AddOptions("scales", scales)
	return chart
}

func examplePolar() *polar.Polar {
	chart := polar.ChartPolarArea([]string{"极坐标示例"})
	chart.SetLabels([]interface{}{"Red", "Orange", "Yellow", "Green", "Blue", "HaHa"})
	chart.AddDataSet(*chart.NewDataSet("# My DataSet 1", []interface{}{12, 13, 19, 4, 4, 20}))
	yAes := options.NewAxes()
	yAes.AddTicks("beginAtZero", "true")
	scales := options.Scales{
		YAxes: []options.Axes{*yAes},
	}
	chart.AddOptions("scales", scales)
	return chart
}

func exampleBubble() *bubble.Bubble {
	chart := bubble.ChartBubble([]string{"气泡图示例"})
	type points struct {
		X int `json:"x"`
		Y int `json:"y"`
		R int `json:"r"`
	}
	one := points{20, 30, 15}
	two := points{40, 10, 10}
	three := points{30, 20, 10}
	four := points{30, 10, 10}
	five := points{25, 15, 10}
	six := points{10, 14, 10}
	seven := points{10, 8, 10}
	eight := points{10, 19, 10}
	nine := points{10, 20, 9}
	ten := points{10, 10, 10}
	chart.AddDataSet(*chart.NewDataSet("# First DataSet", []interface{}{one, two, three, four, five, six, seven, eight, nine, ten}))
	yAes := options.NewAxes()
	yAes.AddTicks("beginAtZero", "true")
	scales := options.Scales{
		YAxes: []options.Axes{*yAes},
	}
	chart.AddOptions("scales", scales)
	return chart
}

func exampleScatter() *scatter.Scatter {
	type point struct {
		X int `json:"x"`
		Y int `json:"y"`
	}
	data := func() []interface{} {
		var results []interface{}
		for i := 0; i < 7; i++ {
			var one point
			one.X = i
			one.Y = rand.Intn(32) - i
			results = append(results, one)
		}
		return results
	}
	chart := scatter.ChartScatter([]string{"Chart.js Scatter Chart"})
	chart.AddDataSet(
		*chart.NewDataSet("# First DataSet", data()),
		*chart.NewDataSet("# Second DataSet", data()))
	return chart
}
