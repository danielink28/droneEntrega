package orientation

type Orientation interface {
	Rotate(orientation string) Orientation
	Position() string
	Move(x, y int) (p, j int)
}
