package configuration

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Color struct {
	RGB string `json:"rgb"` // red, green, and blue
	Hex string `json:"hex"` // 十六进制
}

func (c Color) String() string {
	return c.RGB
}

// value: #666 --> rgb()
func (c *Color) ToRGB(value string) string {
	if strings.HasPrefix(value, "#") {
		v := strings.TrimPrefix(value, "#")
		v64, _ := strconv.ParseInt(v, 16, 32)
		red := v64 >> 16
		green := (v64 & 0x00FF00) >> 8
		blue := v64 & 0x0000FF
		c.RGB = fmt.Sprintf("rgb(%d,%d,%d)", red, green, blue)
	}
	return c.RGB
}

// value: rgb() --> #hex
func (c *Color) ToHex(value string) string {
	var result []string
	if strings.HasPrefix(value, "rgb") {
		v := strings.TrimPrefix(value, "rgb")
		list := strings.FieldsFunc(v, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})

		for _, i := range list {
			a, _ := strconv.ParseInt(i, 0, 10)
			b := strconv.FormatInt(a, 16)
			if b == "0" {
				b = "00"
			}
			result = append(result, b)
		}
	}
	c.Hex = strings.Join(result, "")
	return c.Hex
}

type Colors []Color
