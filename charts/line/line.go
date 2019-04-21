package line

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wuxiaoxiaoshen/gocharts/charts"
	"github.com/wuxiaoxiaoshen/gocharts/charts/options"
)

type Line struct {
	charts.Base
}

type OptionOfLine struct {
}

func ChartLine(title []string) *Line {
	line := &Line{
		Base: charts.Base{
			Type:  charts.ChartTypes[charts.Line],
			Title: title,
		},
	}
	line.Options = make(map[string]interface{})
	line.setTitle()
	return line
}

func (l *Line) setTitle() {
	l.Options["title"] = options.Title{
		Text:    l.Title,
		Display: true,
	}
}

func (l *Line) SetScales(scales options.Scales) {
	l.Options["scales"] = scales
}

func (l Line) Plot() func(writer http.ResponseWriter, request *http.Request) {
	path, _ := os.Getwd()
	plot := filepath.Join(path, "github.com/wuxiaoxiaoshen/gocharts/template/plot.html")
	log.Println("Server on: http://localhost:port" + "/" + l.Type)
	return func(writer http.ResponseWriter, request *http.Request) {
		tpl, _ := template.ParseFiles(plot)
		tpl.Execute(writer, l)
	}

}
