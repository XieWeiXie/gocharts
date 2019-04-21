package options

import (
	"fmt"
	"reflect"
	"regexp"
	"unicode"
)

type Font struct {
	Size   int    `json:"fontSize"`
	Family string `json:"fontFamily"`
	Color  string `json:"fontColor"`
	Style  string `json:"fontStyle"`
	color  Color  `json:"color"`
}

func (F Font) TitleFont() interface{} {
	return F.setValue(F.structOf("title"))

}

func (F Font) structOf(prefix string) reflect.Value {
	typ := reflect.TypeOf(F)
	var newTyp []reflect.StructField
	for i := 0; i < typ.NumField(); i++ {
		tag := fmt.Sprintf("%s", typ.Field(i).Tag)
		// 不可导出自动不处理
		if unicode.IsLower(rune(typ.Field(i).Name[0])) {
			continue
		}
		one := reflect.StructField{
			Name: typ.Field(i).Name,
			Type: typ.Field(i).Type,
			Tag:  reflect.StructTag(tagHandle(tag, prefix)),
		}
		fmt.Println(tag)
		newTyp = append(newTyp, one)
	}
	return reflect.New(reflect.StructOf(newTyp)).Elem()
}

func (F Font) setValue(s reflect.Value) interface{} {
	s.FieldByName("Size").SetInt(int64(F.Size))
	s.FieldByName("Family").SetString(F.Family)
	s.FieldByName("Color").SetString(F.Color)
	s.FieldByName("Style").SetString(F.Style)
	return s.Addr().Interface()
}

func (F Font) BodyFont() interface{} {
	return F.setValue(F.structOf("body"))
}

func (F Font) FooterFont() interface{} {
	return F.setValue(F.structOf("footer"))
}

func tagHandle(value string, prefix string) string {
	reg := regexp.MustCompile(`^json:"(.*?)"`)
	if !reg.MatchString(value) {
		return "-1"
	}
	target := reg.FindAllStringSubmatch(value, -1)[0][1]
	return fmt.Sprintf(`json:"%s%c%s"`, prefix, unicode.ToUpper(rune(target[0])), target[1:])

}
