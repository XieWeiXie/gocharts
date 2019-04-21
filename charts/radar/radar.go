package radar

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wuxiaoxiaoshen/gocharts/charts/options"

	"github.com/wuxiaoxiaoshen/gocharts/charts"
)

type Radar struct {
	charts.Base
}

type OptionOfRadar struct {
}

func ChartRadar(title []string) *Radar {
	radar := &Radar{
		Base: charts.Base{
			Type:  charts.ChartTypes[charts.Radar],
			Title: title,
		},
	}
	radar.Options = make(map[string]interface{})
	radar.setTitle()
	return radar

}

func (r *Radar) setTitle() {
	r.Options["title"] = options.Title{
		Display: true,
		Text:    r.Title,
	}
}

func (r *Radar) Plot() func(writer http.ResponseWriter, request *http.Request) {
	path, _ := os.Getwd()
	plot := filepath.Join(path, "github.com/wuxiaoxiaoshen/gocharts/template/plot.html")
	log.Println("Server on: http://localhost:port" + "/" + r.Type)
	return func(writer http.ResponseWriter, request *http.Request) {
		tpl, _ := template.ParseFiles(plot)
		tpl.Execute(writer, r)
	}

}
