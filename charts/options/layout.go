package options

/*
布局
*/

type LayOut struct {
	Padding PaddingObject `json:"padding"`
}

type PaddingObject struct {
	Left   int `json:"left"`
	Right  int `json:"right"`
	Top    int `json:"top"`
	Bottom int `json:"bottom"`
}
