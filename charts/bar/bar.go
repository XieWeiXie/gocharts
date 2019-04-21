package bar

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wuxiaoxiaoshen/gocharts/charts"
	"github.com/wuxiaoxiaoshen/gocharts/charts/options"
)

type Bar struct {
	charts.Base
}

type OptionOfBar struct {
}

func ChartBar(title []string) *Bar {
	bar := &Bar{
		charts.Base{
			Type:  charts.ChartTypes[charts.Bar],
			Title: title,
		},
	}
	bar.Options = make(map[string]interface{})
	bar.setTitle()
	return bar

}

func (b *Bar) setTitle() {
	b.Options["title"] = options.Title{
		Display: true,
		Text:    b.Title,
	}
}

func (b *Bar) TypeName() string {
	return b.Type
}

func (b Bar) Plot() func(writer http.ResponseWriter, request *http.Request) {
	path, _ := os.Getwd()
	plot := filepath.Join(path, "github.com/wuxiaoxiaoshen/gocharts/template/plot.html")
	log.Println("Server on: http://localhost" + "/" + b.Type)
	return func(writer http.ResponseWriter, request *http.Request) {
		tpl, _ := template.ParseFiles(plot)
		tpl.Execute(writer, b)
	}

}

func (b Bar) Save(fileName string) {}
