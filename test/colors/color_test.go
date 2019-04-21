package colors_test

import (
	"fmt"
	"testing"

	"github.com/wuxiaoxiaoshen/gocharts/charts/options"
)

func TestColor(test *testing.T) {
	color := options.Color{}
	fmt.Println(color.ToRGB("#CC00FF"))
	fmt.Println(color.ToHex("rgb(204,00,255)"))
	fmt.Println(color.Alpha(0.2))
	fmt.Println(options.RandomColor(2))
	fmt.Println(options.RandomColorAlpha(2, 0.2))
	fmt.Println(options.RandomColorBoth(2, 0.3))
}
