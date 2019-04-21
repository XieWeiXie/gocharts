package main

import (
	"fmt"
	"net/http"

	"github.com/wuxiaoxiaoshen/gocharts/charts/radar"

	"github.com/wuxiaoxiaoshen/gocharts/charts/bar"

	"github.com/wuxiaoxiaoshen/gocharts/charts/options"

	"github.com/wuxiaoxiaoshen/gocharts/charts/line"
)

func main() {

	line := exampleLine()
	bar := exampleBar()
	http.HandleFunc(fmt.Sprintf("/%s", line.TypeName()), line.Plot())
	http.HandleFunc(fmt.Sprintf("/%s", bar.TypeName()), bar.Plot())
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
	return chart
}
