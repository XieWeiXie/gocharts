package configuration

import "fmt"

/*
动画
*/

type Animation struct {
	Duration   int    `json:"duration"`
	Easing     string `json:"easing"`
	OnProgress func() `json:"onProgress"`
	OnComplete func() `json:"onComplete"`
}

var easeKeys = []string{
	"Quad",
	"Cubic",
	"Quart",
	"Quint",
	"Sine",
	"Expo",
	"Circ",
	"Elastic",
	"Back",
	"Bounce",
}

// easing 选项
func easingAvailableOption() []string {
	var returnResult []string
	var prefix = []string{"In", "Out", "InOut"}
	for _, key := range easeKeys {
		for _, pre := range prefix {
			returnResult = append(returnResult, fmt.Sprintf("%s%s", pre, key))
		}
	}
	return returnResult
}
