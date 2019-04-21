package bubble

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/wuxiaoxiaoshen/gocharts/charts"
	"github.com/wuxiaoxiaoshen/gocharts/charts/options"
)

type Bubble struct {
	charts.Base
}

type OptionOfBubble struct {
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
	R int `json:"r"`
}

func ChartBubble(title []string) *Bubble {
	bubble := &Bubble{
		charts.Base{
			Type:  charts.ChartTypes[charts.Bubble],
			Title: title,
		},
	}
	bubble.Options = make(map[string]interface{})
	bubble.setTitle()
	return bubble
}

func (b *Bubble) setTitle() {
	b.Options["title"] = options.Title{
		Display: true,
		Text:    b.Title,
	}
}

func (b *Bubble) AddPoints() {

}

func (b Bubble) Plot() func(writer http.ResponseWriter, request *http.Request) {
	path, _ := os.Getwd()
	plot := filepath.Join(path, "github.com/wuxiaoxiaoshen/gocharts/template/plot.html")
	log.Println("Server on: http://localhost:port" + "/" + b.Type)
	return func(writer http.ResponseWriter, request *http.Request) {
		tpl, _ := template.ParseFiles(plot)
		tpl.Execute(writer, b)
	}
}
