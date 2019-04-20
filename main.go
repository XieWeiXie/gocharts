package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wuxiaoxiaoshen/gocharts/charts"
)

type My struct {
	Charts []byte
}

func main() {
	data := charts.DefaultChartsBase
	data.SetLabels([]string{"red", "blue", "yellow"})
	content, err := data.JsonMarshal()
	if err != nil {
		return
	}
	fmt.Println(string(content))
	var m My
	m.Charts = content
	fmt.Println(filepath.Abs("template/plot.html"))
	currentPath, _ := os.Getwd()
	plot := filepath.Join(currentPath, "github.com/wuxiaoxiaoshen/gocharts/template/plot.html")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		tpl, _ := template.ParseFiles(plot)
		tpl.Execute(writer, data)
	})
	http.ListenAndServe(":8888", nil)

}
