package node

const FACTOR = 257

type Color struct {
	R uint8
	G uint8
	B uint8
}

// NewColor return a new color, divide by 257 to have value between 0 and 255
func NewColor(r, g, b uint32) *Color {
	return &Color{
		uint8(r / FACTOR),
		uint8(g / FACTOR),
		uint8(b / FACTOR),
	}
}

// IsStartPoint return true if node it's start point
func (c Color) IsStartPoint() bool {
	return c.R == 0 && c.G == 255 && c.B == 0
}

// IsEndPoint return true if node it's end point
func (c Color) IsEndPoint() bool {
	return c.R == 255 && c.G == 0 && c.B == 0
}
