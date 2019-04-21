package main

import (
	"fmt"
	"net/http"

	"github.com/wuxiaoxiaoshen/gocharts/charts/line"
)

type My struct {
	Charts []byte
}

func main() {

	chart := line.ChartLine("line", []string{"折线图示例"})
	chart.SetLabels([]string{"Red", "Blue", "Yellow", "Green", "Purple", "Orange"})
	chart.AddDataSet(*chart.NewDataSet("# of votes", []interface{}{12, 19, 3, 5, 2, 3}))
	fmt.Println(chart)

	http.HandleFunc("/", chart.Plot())
	http.ListenAndServe(":8888", nil)

}
