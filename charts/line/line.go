package line

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wuxiaoxiaoshen/gocharts/charts"
	"github.com/wuxiaoxiaoshen/gocharts/charts/configuration"
)

type Line struct {
	charts.Base
}

type OptionOfLine struct {
}

func ChartLine(typ string, title []string) *Line {
	line := &Line{
		Base: charts.Base{
			Type:  typ,
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

func (l Line) Plot() func(writer http.ResponseWriter, request *http.Request) {
	path, _ := os.Getwd()
	plot := filepath.Join(path, "github.com/wuxiaoxiaoshen/gocharts/template/plot.html")
	log.Println(path, plot)
	return func(writer http.ResponseWriter, request *http.Request) {
		tpl, _ := template.ParseFiles(plot)
		tpl.Execute(writer, l)
	}

}

func (l Line) Save(fileName string) {}
