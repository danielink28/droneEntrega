package orientation

type West struct {
}

func (c West) Move(x int, y int) (j, p int) {
	return x - 1, y
}

func (c West) Rotate(orientation string) Orientation {
	if orientation == "D" {
		return North{}
	}
	return South{}
}
func (c West) Position() string {
	return "Occidente"
}
