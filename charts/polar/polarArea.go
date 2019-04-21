package polar

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wuxiaoxiaoshen/gocharts/charts"
	"github.com/wuxiaoxiaoshen/gocharts/charts/options"
)

type Polar struct {
	charts.Base
}

type OptionOfPolar struct {
}

func ChartPolarArea(title []string) *Polar {
	polar := &Polar{
		charts.Base{
			Type:  charts.ChartTypes[charts.Polar],
			Title: title,
		},
	}
	polar.Options = make(map[string]interface{})
	polar.setTitle()
	return polar
}

func (p *Polar) setTitle() {
	p.Options["title"] = options.Title{
		Display: true,
		Text:    p.Title,
	}
}

func (p Polar) Plot() func(writer http.ResponseWriter, request *http.Request) {
	path, _ := os.Getwd()
	plot := filepath.Join(path, "github.com/wuxiaoxiaoshen/gocharts/template/plot.html")
	log.Println("Server on: http://localhost:port" + "/" + p.Type)
	return func(writer http.ResponseWriter, request *http.Request) {
		tpl, _ := template.ParseFiles(plot)
		tpl.Execute(writer, p)
	}

}
