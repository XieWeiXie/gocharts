package colors_test

import (
	"fmt"
	"testing"

	"github.com/wuxiaoxiaoshen/gocharts/charts/configuration"
)

func TestColor(test *testing.T) {
	color := configuration.Color{}
	fmt.Println(color.ToRGB("#CC00FF"))
	fmt.Println(color.ToHex("rgb(204,00,255)"))
}
