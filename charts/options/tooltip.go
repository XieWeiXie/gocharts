package options

type ToolTip struct {
	Enabled               bool     `json:"enabled"`
	Custom                func()   `json:"custom"`
	Mode                  string   `json:"mode"`
	InterSect             bool     `json:"interSect"`
	Position              Position `json:"position"`
	BackgroundColor       Color    `json:"backgroundColor"`
	TitleSpacing          int      `json:"titleSpacing"`
	TitleMarginBottom     int      `json:"titleMarginBottom"`
	TitleFont             Font
	BodySpacing           int `json:"bodySpacing"`
	BodyFont              Font
	FooterSpacing         int `json:"footerSpacing"`
	FooterMarginTop       int `json:"footerMarginTop"`
	FooterFont            Font
	XPadding              int    `json:"xPadding"`
	YPadding              int    `json:"yPadding"`
	CaretPadding          int    `json:"caretPadding"`
	CaretSize             int    `json:"caretSize"`
	CornerRadius          int    `json:"cornerRadius"`
	MultipleKeyBackground string `json:"multikeyBackground"`
	multipleKeyColor      Color
	DisplayColors         bool   `json:"displayColors"`
	BorderColor           string `json:"borderColor"`
	borderColor           Color
	BorderWidth           int `json:"borderWidth"`
}
