package line

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/wuxiaoxiaoshen/gocharts/charts/line"
)

func TestLine(test *testing.T) {
	chart := line.ChartLine("line", []string{"折线图示例"})
	chart.SetLabels([]string{"Red", "Blue", "Yellow", "Green", "Purple", "Orange"})
	chart.AddDataSet(*chart.NewDataSet("# of votes", []interface{}{12, 19, 3, 5, 2, 3}))
	fmt.Println(chart)

	fmt.Println(filepath.Abs("template/plot.html"))
	currentPath, _ := os.Getwd()
	plot := filepath.Join(currentPath, "github.com/wuxiaoxiaoshen/gocharts/template/plot.html")
	fmt.Println(plot)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tpl, _ := template.ParseFiles(plot)
		tpl.Execute(writer, chart)
	})
	http.ListenAndServe(":8888", nil)
}
