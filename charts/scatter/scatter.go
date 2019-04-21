package scatter

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wuxiaoxiaoshen/gocharts/charts"
	"github.com/wuxiaoxiaoshen/gocharts/charts/options"
)

type Scatter struct {
	charts.Base
}

type OptionOfScatter struct {
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func ChartScatter(title []string) *Scatter {
	scatter := &Scatter{
		charts.Base{
			Title: title,
			Type:  charts.ChartTypes[charts.Scatter],
		},
	}
	scatter.Options = make(map[string]interface{})
	scatter.setTitle()
	return scatter
}

func (s Scatter) setTitle() {
	s.Options["title"] = options.Title{
		Display: true,
		Text:    s.Title,
	}
}

func (s Scatter) Plot() func(writer http.ResponseWriter, request *http.Request) {
	path, _ := os.Getwd()
	plot := filepath.Join(path, "github.com/wuxiaoxiaoshen/gocharts/template/plot.html")
	log.Println("Server on: http://localhost:port" + "/" + s.Type)
	return func(writer http.ResponseWriter, request *http.Request) {
		tpl, _ := template.ParseFiles(plot)
		tpl.Execute(writer, s)
	}
}
