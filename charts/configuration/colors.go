package options

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Color struct {
	Name   string `json:"name"`
	ChName string `json:"ch_name"`
	Hex    string `json:"hex"` // 十六进制
	RGB    string `json:"rgb"` // red, green, and blue
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

func (c Color) Alpha(delta float64) string {
	if c.RGB == "" {
		return "-1"
	}
	if delta > 1 || delta < 0 {
		return "-1"
	}
	reg := regexp.MustCompile(`rgb\((.*?)\)`)
	if !reg.MatchString(c.RGB) {
		return "-1"
	}
	matchString := reg.FindAllStringSubmatch(c.RGB, -1)[0][1]
	return fmt.Sprintf("rgba(%s)", fmt.Sprintf("%s,%f", matchString, delta))
}

type Colors []Color

func colorCollections() Colors {
	return Colors{
		{"LightPink", "浅粉红", "#FFB6C1", "rgb(255,182,193)"},
		{"Pink", "粉红", "#FFC0CB", "rgb(255,192,203)"},
		{"Crimson", "猩红", "#DC143C", "rgb(220,20,60)"},
		{"LavenderBlush", "脸红的淡紫色", "#FFF0F5", "rgb(255,240,245)"},
		{"PaleVioletRed", "苍白的紫罗兰红色", "#DB7093", "rgb(219,112,147)"},
		{"HotPink", "热情的粉红", "#FF69B4", "rgb(255,105,180)"},
		{"DeepPink", "深粉色", "#FF1493", "rgb(255,20,147)"},
		{"MediumVioletRed", "适中的紫罗兰红色", "#C71585", "rgb(199,21,133)"},
		{"Orchid", "兰花的紫色", "#DA70D6", "rgb(218,112,214)"},
		{"Thistle", "蓟", "#D8BFD8", "rgb(216,191,216)"},
		{"plum", "李子", "#DDA0DD", "rgb(221,160,221)"},
		{"Violet", "紫罗兰", "#EE82EE", "rgb(238,130,238)"},
		{"Magenta", "洋红", "#FF00FF", "rgb(255,0,255)"},
		{"Fuchsia", "灯笼海棠(紫红色)", "#FF00FF", "rgb(255,0,255)"},
		{"DarkMagenta", "深洋红色", "#8B008B", "rgb(139,0,139)"},
		{"Purple", "紫色", "#800080", "rgb(128,0,128)"},
		{"MediumOrchid", "适中的兰花紫", "#BA55D3", "rgb(186,85,211)"},
		{"DarkVoilet", "深紫罗兰色", "#9400D3", "rgb(148,0,211)"},
		{"DarkOrchid", "深兰花紫", "#9932CC", "rgb(153,50,204)"},
		{"Indigo", "靛青", "#4B0082", "rgb(75,0,130)"},
		{"BlueViolet", "深紫罗兰的蓝色", "#8A2BE2", "rgb(138,43,226)"},
		{"MediumPurple", "适中的紫色", "#9370DB", "rgb(147,112,219)"},
		{"MediumSlateBlue", "适中的板岩暗蓝灰色", "#7B68EE", "rgb(123,104,238)"},
		{"SlateBlue", "板岩暗蓝灰色", "#6A5ACD", "rgb(106,90,205)"},
		{"DarkSlateBlue", "深岩暗蓝灰色", "#483D8B", "rgb(72,61,139)"},
		{"Lavender", "熏衣草花的淡紫色", "#E6E6FA", "rgb(230,230,250)"},
		{"GhostWhite", "幽灵的白色", "#F8F8FF", "rgb(248,248,255)"},
		{"Blue", "纯蓝", "#0000FF", "rgb(0,0,255)"},
		{"MediumBlue", "适中的蓝色", "#0000CD", "rgb(0,0,205)"},
		{"MidnightBlue", "午夜的蓝色", "#191970", "rgb(25,25,112)"},
		{"DarkBlue", "深蓝色", "#00008B", "rgb(0,0,139)"},
		{"Navy", "海军蓝", "#000080", "rgb(0,0,128)"},
		{"RoyalBlue", "皇家蓝", "#4169E1", "rgb(65,105,225)"},
		{"CornflowerBlue", "矢车菊的蓝色", "#6495ED", "rgb(100,149,237)"},
		{"LightSteelBlue", "淡钢蓝", "#B0C4DE", "rgb(176,196,222)"},
		{"LightSlateGray", "浅石板灰", "#778899", "rgb(119,136,153)"},
		{"SlateGray", "石板灰", "#708090", "rgb(112,128,144)"},
		{"DoderBlue", "道奇蓝", "#1E90FF", "rgb(30,144,255)"},
		{"AliceBlue", "爱丽丝蓝", "#F0F8FF", "rgb(240,248,255)"},
		{"SteelBlue", "钢蓝", "#4682B4", "rgb(70,130,180)"},
		{"LightSkyBlue", "淡蓝色", "#87CEFA", "rgb(135,206,250)"},
		{"SkyBlue", "天蓝色", "#87CEEB", "rgb(135,206,235)"},
		{"DeepSkyBlue", "深天蓝", "#00BFFF", "rgb(0,191,255)"},
		{"LightBLue", "淡蓝", "#ADD8E6", "rgb(173,216,230)"},
		{"PowDerBlue", "火药蓝", "#B0E0E6", "rgb(176,224,230)"},
		{"CadetBlue", "军校蓝", "#5F9EA0", "rgb(95,158,160)"},
		{"Azure", "蔚蓝色", "#F0FFFF", "rgb(240,255,255)"},
		{"LightCyan", "淡青色", "#E1FFFF", "rgb(225,255,255)"},
		{"PaleTurquoise", "苍白的绿宝石", "#AFEEEE", "rgb(175,238,238)"},
		{"Cyan", "青色", "#00FFFF", "rgb(0,255,255)"},
		{"Aqua", "水绿色", "#00FFFF", "rgb(0,255,255)"},
		{"DarkTurquoise", "深绿宝石", "#00CED1", "rgb(0,206,209)"},
		{"DarkSlateGray", "深石板灰", "#2F4F4F", "rgb(47,79,79)"},
		{"DarkCyan", "深青色", "#008B8B", "rgb(0,139,139)"},
		{"Teal", "水鸭色", "#008080", "rgb(0,128,128)"},
		{"MediumTurquoise", "适中的绿宝石", "#48D1CC", "rgb(72,209,204)"},
		{"LightSeaGreen", "浅海洋绿", "#20B2AA", "rgb(32,178,170)"},
		{"Turquoise", "绿宝石", "#40E0D0", "rgb(64,224,208)"},
		{"Auqamarin", "绿玉\\碧绿色", "#7FFFAA", "rgb(127,255,170)"},
		{"MediumAquamarine", "适中的碧绿色", "#00FA9A", "rgb(0,250,154)"},
		{"MediumSpringGreen", "适中的春天的绿色", "#00FF7F", "rgb(0,255,127)"},
		{"MintCream", "薄荷奶油", "#F5FFFA", "rgb(245,255,250)"},
		{"SpringGreen", "春天的绿色", "#3CB371", "rgb(60,179,113)"},
		{"SeaGreen", "海洋绿", "#2E8B57", "rgb(46,139,87)"},
		{"Honeydew", "蜂蜜", "#F0FFF0", "rgb(240,255,240)"},
		{"LightGreen", "淡绿色", "#90EE90", "rgb(144,238,144)"},
		{"PaleGreen", "苍白的绿色", "#98FB98", "rgb(152,251,152)"},
		{"DarkSeaGreen", "深海洋绿", "#8FBC8F", "rgb(143,188,143)"},
		{"LimeGreen", "酸橙绿", "#32CD32", "rgb(50,205,50)"},
		{"Lime", "酸橙色", "#00FF00", "rgb(0,255,0)"},
		{"ForestGreen", "森林绿", "#228B22", "rgb(34,139,34)"},
		{"Green", "纯绿", "#008000", "rgb(0,128,0)"},
		{"DarkGreen", "深绿色", "#006400", "rgb(0,100,0)"},
		{"Chartreuse", "查特酒绿", "#7FFF00", "rgb(127,255,0)"},
		{"LawnGreen", "草坪绿", "#7CFC00", "rgb(124,252,0)"},
		{"GreenYellow", "绿黄色", "#ADFF2F", "rgb(173,255,47)"},
		{"OliveDrab", "橄榄土褐色", "#556B2F", "rgb(85,107,47)"},
		{"Beige", "米色(浅褐色)", "#F5F5DC", "rgb(245,245,220)"},
		{"LightGoldenrodYellow", "浅秋麒麟黄", "#FAFAD2", "rgb(250,250,210)"},
		{"Ivory", "象牙", "#FFFFF0", "rgb(255,255,240)"},
		{"LightYellow", "浅黄色", "#FFFFE0", "rgb(255,255,224)"},
		{"Yellow", "纯黄", "#FFFF00", "rgb(255,255,0)"},
		{"Olive", "橄榄", "#808000", "rgb(128,128,0)"},
		{"DarkKhaki", "深卡其布", "#BDB76B", "rgb(189,183,107)"},
		{"LemonChiffon", "柠檬薄纱", "#FFFACD", "rgb(255,250,205)"},
		{"PaleGodenrod", "灰秋麒麟", "#EEE8AA", "rgb(238,232,170)"},
		{"Khaki", "卡其布", "#F0E68C", "rgb(240,230,140)"},
		{"Gold", "金", "#FFD700", "rgb(255,215,0)"},
		{"Cornislk", "玉米色", "#FFF8DC", "rgb(255,248,220)"},
		{"GoldEnrod", "秋麒麟", "#DAA520", "rgb(218,165,32)"},
		{"FloralWhite", "花的白色", "#FFFAF0", "rgb(255,250,240)"},
		{"OldLace", "老饰带", "#FDF5E6", "rgb(253,245,230)"},
		{"Wheat", "小麦色", "#F5DEB3", "rgb(245,222,179)"},
		{"Moccasin", "鹿皮鞋", "#FFE4B5", "rgb(255,228,181)"},
		{"Orange", "橙色", "#FFA500", "rgb(255,165,0)"},
		{"PapayaWhip", "番木瓜", "#FFEFD5", "rgb(255,239,213)"},
		{"BlanchedAlmond", "漂白的杏仁", "#FFEBCD", "rgb(255,235,205)"},
		{"NavajoWhite", "纳瓦霍白", "#FFDEAD", "rgb(255,222,173)"},
		{"AntiqueWhite", "古代的白色", "#FAEBD7", "rgb(250,235,215)"},
		{"Tan", "晒黑", "#D2B48C", "rgb(210,180,140)"},
		{"BrulyWood", "结实的树", "#DEB887", "rgb(222,184,135)"},
		{"Bisque", "(浓汤)乳脂,番茄等", "#FFE4C4", "rgb(255,228,196)"},
		{"DarkOrange", "深橙色", "#FF8C00", "rgb(255,140,0)"},
		{"Linen", "亚麻布", "#FAF0E6", "rgb(250,240,230)"},
		{"Peru", "秘鲁", "#CD853F", "rgb(205,133,63)"},
		{"PeachPuff", "桃色", "#FFDAB9", "rgb(255,218,185)"},
		{"SandyBrown", "沙棕色", "#F4A460", "rgb(244,164,96)"},
		{"Chocolate", "巧克力", "#D2691E", "rgb(210,105,30)"},
		{"SaddleBrown", "马鞍棕色", "#8B4513", "rgb(139,69,19)"},
		{"SeaShell", "海贝壳", "#FFF5EE", "rgb(255,245,238)"},
		{"Sienna", "黄土赭色", "#A0522D", "rgb(160,82,45)"},
		{"LightSalmon", "浅鲜肉(鲑鱼)色", "#FFA07A", "rgb(255,160,122)"},
		{"Coral", "珊瑚", "#FF7F50", "rgb(255,127,80)"},
		{"OrangeRed", "橙红色", "#FF4500", "rgb(255,69,0)"},
		{"DarkSalmon", "深鲜肉(鲑鱼)色", "#E9967A", "rgb(233,150,122)"},
		{"Tomato", "番茄", "#FF6347", "rgb(255,99,71)"},
		{"MistyRose", "薄雾玫瑰", "#FFE4E1", "rgb(255,228,225)"},
		{"Salmon", "鲜肉(鲑鱼)色", "#FA8072", "rgb(250,128,114)"},
		{"Snow", "雪", "#FFFAFA", "rgb(255,250,250)"},
		{"LightCoral", "淡珊瑚色", "#F08080", "rgb(240,128,128)"},
		{"RosyBrown", "玫瑰棕色", "#BC8F8F", "rgb(188,143,143)"},
		{"IndianRed", "印度红", "#CD5C5C", "rgb(205,92,92)"},
		{"Red", "纯红", "#FF0000", "rgb(255,0,0)"},
		{"Brown", "棕色", "#A52A2A", "rgb(165,42,42)"},
		{"FireBrick", "耐火砖", "#B22222", "rgb(178,34,34)"},
		{"DarkRed", "深红色", "#8B0000", "rgb(139,0,0)"},
		{"Maroon", "栗色", "#800000", "rgb(128,0,0)"},
		{"White", "纯白", "#FFFFFF", "rgb(255,255,255)"},
		{"WhiteSmoke", "白烟", "#F5F5F5", "rgb(245,245,245)"},
		{"Gainsboro", "亮灰色", "#DCDCDC", "rgb(220,220,220)"},
		{"LightGrey", "浅灰色", "#D3D3D3", "rgb(211,211,211)"},
		{"Silver", "银白色", "#C0C0C0", "rgb(192,192,192)"},
		{"DarkGray", "深灰色", "#A9A9A9", "rgb(169,169,169)"},
		{"Gray", "灰色", "#808080", "rgb(128,128,128)"},
		{"DimGray", "暗淡的灰色", "#696969", "rgb(105,105,105)"},
		{"Black", "纯黑", "#000000", "rgb(0,0,0)"},
	}
}

func RandomColor(number int) Colors {
	var results Colors
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < number; i++ {
		n := rand.Intn(len(colorCollections()))
		results = append(results, colorCollections()[n])
	}
	return results
}

func RandomColorAlpha(number int, delta float64) []string {
	results := RandomColor(number)
	var resultsOfAlpha []string
	for _, i := range results {
		resultsOfAlpha = append(resultsOfAlpha, i.Alpha(delta))
	}
	return resultsOfAlpha
}

func RandomColorBoth(number int, delta float64) ([]string, []string) {
	results := RandomColor(number)
	var resultsOfColor []string
	var resultsOfAlpha []string
	for _, i := range results {
		resultsOfColor = append(resultsOfColor, i.String())
		resultsOfAlpha = append(resultsOfAlpha, i.Alpha(delta))
	}
	return resultsOfColor, resultsOfAlpha
}
