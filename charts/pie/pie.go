package pie

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wuxiaoxiaoshen/gocharts/charts"
	"github.com/wuxiaoxiaoshen/gocharts/charts/options"
)

type Pie struct {
	charts.Base
}

type OptionOfPie struct {
}

func ChartPie(title []string) *Pie {
	pie := &Pie{
		charts.Base{
			Type:  charts.ChartTypes[charts.Pie],
			Title: title,
		},
	}
	pie.Options = make(map[string]interface{})
	pie.setTitle()
	return pie
}

func (p *Pie) setTitle() {
	p.Options["title"] = options.Title{
		Display: true,
		Text:    p.Title,
	}
}

func (p Pie) Plot() func(writer http.ResponseWriter, request *http.Request) {
	path, _ := os.Getwd()
	plot := filepath.Join(path, "github.com/wuxiaoxiaoshen/gocharts/template/plot.html")
	log.Println("Server on: http://localhost:port" + "/" + p.Type)
	return func(writer http.ResponseWriter, request *http.Request) {
		tpl, _ := template.ParseFiles(plot)
		tpl.Execute(writer, p)
	}

}
