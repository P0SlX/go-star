package node

type Color struct {
	R uint8
	G uint8
	B uint8
}

// Divise par 257 pour obtenir des valeurs entre 0 et 255
func NewColor(r, g, b uint32) *Color {
	return &Color{
		uint8(r / 257),
		uint8(g / 257),
		uint8(b / 257),
	}
}

func (c Color) IsStartPoint() bool {
	return c.R == 0 && c.G == 255 && c.B == 0
}

func (c Color) IsEndPoint() bool {
	return c.R == 255 && c.G == 0 && c.B == 0
}
