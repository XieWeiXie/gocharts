package configuration

type Position struct {
}

func (p Position) Top() string {
	return "top"
}
func (p Position) Left() string {
	return "left"
}

func (p Position) Right() string {
	return "right"
}
func (p Position) Bottom() string {
	return "bottom"
}

func (p Position) Average() string {
	return "average"
}

func (p Position) Nearest() string {
	return "nearest"
}

func (p Position) AvailableOption() []string {
	var result []string
	result = append(result, p.Top(), p.Left(), p.Right(), p.Bottom())
	return result
}
