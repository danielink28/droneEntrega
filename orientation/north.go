package orientation

type North struct {
}

func (c North) Move(x int, y int) (j, p int) {
	return x, y + 1
}

func (c North) Rotate(orientation string) Orientation {
	if orientation == "D" {
		return East{}
	}
	return West{}
}

func (c North) Position() string {
	return "Norte"
}
