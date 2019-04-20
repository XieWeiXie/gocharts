package charts

import (
	"encoding/json"
)

type Charts interface {
	Plot()
	Save()
}

type Base struct {
	Type    string `json:"type"`
	Data    `json:"data"`
	Options `json:"options"`
}

var DefaultChartsBase = Base{Type: "bar", Data: Data{
	Labels: []string{"red", "da", "d"},
	DataSets: []DataSets{
		{
			Label:       "# of Votes",
			Data:        []interface{}{12, 19, 3, 5, 2, 3},
			BorderWidth: 1,
		},
	},
}}

type Data struct {
	Labels   []string   `json:"labels"`
	DataSets []DataSets `json:"datasets"`
}

type DataSets struct {
	Label           string        `json:"label"`
	Data            []interface{} `json:"data"`
	BackgroundColor []interface{} `json:"backgroundColor,omitempty"`
	BorderColor     []interface{} `json:"borderColor,omitempty"`
	BorderWidth     int           `json:"borderWidth"`
}

type Options map[string]interface{}

func (base *Base) SetLabels(labels []string) {
	base.Data.Labels = labels
}

func (base *Base) JsonMarshal() ([]byte, error) {
	return json.Marshal(base)
}

func (base Base) ChartsOpts() ([]byte, error) {
	return base.JsonMarshal()
}
