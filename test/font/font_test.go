package font

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gin-gonic/gin/json"

	"github.com/wuxiaoxiaoshen/gocharts/charts/configuration"
)

func TestFont(test *testing.T) {
	var f options.Font
	f = options.Font{
		Size: 10,
	}

	t := f.TitleFont()
	w := new(bytes.Buffer)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}

	fmt.Printf("json:  %s", w.Bytes())

}
