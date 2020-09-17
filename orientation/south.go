package orientation

type South struct {
}

func (c South) Move(x int, y int) (j, p int) {
	return x, y - 1
}

func (c South) Rotate(orientation string) Orientation {
	if orientation == "D" {
		return West{}
	}
	return East{}
}

func (c South) Position() string {
	return "Sur"
}
