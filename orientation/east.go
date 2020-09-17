package orientation

type East struct {
}

func (c East) Move(x int, y int) (j, p int) {
	return x + 1, y
}

func (c East) Rotate(orientation string) Orientation {
	if orientation == "D" {
		return South{}
	}
	return North{}
}

func (c East) Position() string {
	return "Oriente"
}
